package controller

import (
	"github.com/gin-gonic/gin"
	"zx/database"
	"zx/model"
)

var db = database.Init()

func Index(c *gin.Context) {
	user := model.User{
		Name: c.Query("name"),
		Sex:  c.Query("sex"),
	}
	db.Create(&user)

	c.JSON(200, gin.H{
		"msg":  true,
		"data": &user,
	})

}
