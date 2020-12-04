package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sepehrhakimi90/go-tic-tac-toe/web/controller"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*.html")
	r.GET("/", controller.Index)
	r.POST("/newGame", controller.NewGame)
	r.POST("/move", controller.Move)
	r.Static("/public", "./public")

	r.Run(":8080")
}
