package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-demo/common"
	"go-demo/controllers"
	"go-demo/initializers"
	"go-demo/middleware"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.Use(middleware.Recovery())
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items", middleware.RequireAuth())
		{
			items.POST("", controllers.CreateItem(initializers.DB))
			items.GET("", controllers.ListItem(initializers.DB))
			items.GET("/:id", controllers.GetItem(initializers.DB))
			items.PATCH("/:id", controllers.UpdateItem(initializers.DB))
			items.DELETE("/:id", controllers.DeleteItem(initializers.DB))
		}
		v1.POST("/sign-up", controllers.SignUp(initializers.DB))
		v1.POST("/login", controllers.Login(initializers.DB))
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
