package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mtdx/case-api/common"
	"github.com/mtdx/case-api/middleware"
)

func main() {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	db := common.Connect("main")
	defer db.Close()

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	router.Run()
}
