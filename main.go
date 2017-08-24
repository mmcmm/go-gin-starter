package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mtdx/case-api/middleware"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.Run(":9000")
}
