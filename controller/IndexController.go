package controller

import (
	"github.com/gin-gonic/gin"
	"zx/database"
	"zx/model"
	"zx/untils"
)

var db = database.Init()

func Index(c *gin.Context) {
	user := model.User{
		Name: c.Query("name"),
		Sex:  c.Query("sex"),
	}
	create := db.Create(&user)
	if create != nil {
		return
	}

	c.JSON(200, gin.H{
		"msg":  true,
		"data": &user,
	})

}

func Test(c *gin.Context) error {
	return untils.ParameterError("出现了未知的错误")
}
