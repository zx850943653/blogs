package main

import (
	"github.com/gin-gonic/gin"
	"zx/rout"
	"zx/untils"
)

func main() {
	r := gin.Default()
	r.Use(untils.LogerMiddleware())
	rout.Rout(r)
	r.Run(":9000")
}
