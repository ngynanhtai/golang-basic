package initializers

import (
	"github.com/gin-gonic/gin"
	"go-demo/common"
	"go-demo/models"
	"net/http"
)

func SyncDatabase() func(*gin.Context) {
	return func(c *gin.Context) {
		if err := DB.AutoMigrate(&models.User{}); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrDB(models.ErrCannotSyncUser))
		}
	}
}
