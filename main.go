package main

import (
	"loverWait/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// gin.SetMode(gin.ReleaseMode)
	// r.LoadHTMLFiles("Statics/index.html")
	// r.LoadHTMLFiles("Statics/sightsPlace.html")

	// dboption.CreateTableComment()

	r.Static("/arr", "./statics/table")
	r.LoadHTMLGlob("Statics/view/*.html")
	r.GET("", controller.TitlePage)
	r.GET("/wonder", controller.Wonder)
	r.GET("/comment", controller.Comments)
	r.POST("/comment", controller.CommentsPost)
	r.POST("/select", controller.Choose)
	r.Run(":8088")

}
