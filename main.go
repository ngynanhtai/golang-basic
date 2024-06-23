package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-demo/common"
	"go-demo/middleware"
	ginItem "go-demo/modules/item/transport/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return
	}

	r := gin.Default()
	r.Use(middleware.Recovery())
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginItem.CreateItem(db))
			items.GET("", ginItem.ListItem(db))
			items.GET("/:id", ginItem.GetItem(db))
			items.PATCH("/:id", ginItem.UpdateItem(db))
			items.DELETE("/:id", ginItem.DeleteItem(db))
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

	err = r.Run(":3000")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
