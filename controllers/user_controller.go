package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-demo/biz"
	"go-demo/common"
	"go-demo/models/user"
	"go-demo/storage"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

func SignUp(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data user.UserCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewUserBiz(store)

		if err := business.CreateUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}

func Login(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data user.UserCreation
		var entity *user.User

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewUserBiz(store)

		entity, err := business.GetUserByEmail(c.Request.Context(), data.Email)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(entity.Password), []byte(data.Password)); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(common.ErrLoginFailed))
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": entity.Id,
			"exp": time.Now().Add(time.Hour).Unix(),
		})
		tokenJwt, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenJwt, 3600, "", "", false, true)
		c.JSON(http.StatusOK, gin.H{
			"login": true,
		})
	}
}
