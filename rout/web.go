package rout

import (
	"github.com/gin-gonic/gin"
	"zx/controller"
	"zx/database"
	"zx/migrate"
	"zx/untils"
)

var db = database.Init()

func Rout(r *gin.Engine) {
	//数据库迁移
	migrate.Model()

	//路由
	r.GET("/index", controller.Index)
	r.GET("/error", untils.Wrapper(controller.Test))
	defer db.Close()
}
