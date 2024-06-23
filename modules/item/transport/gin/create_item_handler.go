package ginItem

import (
	"github.com/gin-gonic/gin"
	"go-demo/common"
	"go-demo/modules/item/biz"
	"go-demo/modules/item/model"
	"go-demo/modules/item/storage"
	"gorm.io/gorm"
	"net/http"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreateItemBiz(store)

		if err := business.CreateItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
