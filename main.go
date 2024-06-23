package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-demo/common"
	"go-demo/initializers"
	"go-demo/middleware"
	ginItem "go-demo/modules/item/transport/gin"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.Use(middleware.Recovery())
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginItem.CreateItem(initializers.DB))
			items.GET("", ginItem.ListItem(initializers.DB))
			items.GET("/:id", ginItem.GetItem(initializers.DB))
			items.PATCH("/:id", ginItem.UpdateItem(initializers.DB))
			items.DELETE("/:id", ginItem.DeleteItem(initializers.DB))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		go func() {
			defer common.Recovery()
			fmt.Println([]int{}[0])
		}()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run(":3000")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
