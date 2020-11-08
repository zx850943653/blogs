package main

import (
	"github.com/gin-gonic/gin"
	"zx/rout"
)

func main() {
	r := gin.Default()
	rout.Rout(r)
	r.Run(":9000")
}
