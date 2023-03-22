package controller

import (
	dboption "loverWait/DBoption"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TitlePage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func Wonder(c *gin.Context) {
	// res := dboption.SelectByplace("黄龙\r")
	c.HTML(200, "SightsPlace.html", nil)
}

func Choose(c *gin.Context) {
	name := c.PostForm("ID")
	id, _ := strconv.ParseInt(name, 10, 10)
	tes := dboption.SelectByID(int(id))
	c.HTML(200, "selectRes.html", gin.H{
		"地点": tes.Place,
		"去过": tes.Did,
		"和谁": tes.Who,
	})
}

func Comments(c *gin.Context) {
	c.HTML(200, "comment.html", nil)

}

func CommentsPost(c *gin.Context) {
	say := c.PostForm("comment")
	dboption.InsertComment(say)
	c.JSON(200, "留言成功！！！！！！")

}
