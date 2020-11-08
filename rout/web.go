package rout

import (
	"github.com/gin-gonic/gin"
	"zx/controller"
	"zx/database"
	"zx/migrate"
)

var db = database.Init()

func Rout(r *gin.Engine) {
	//数据库迁移
	migrate.Model()
	//路由
	r.GET("/index", controller.Index)
	defer db.Close()
}
