package main

import (
	"github.com/mtdx/case-api/handlers"
)

func initRoutes() {
	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/", handlers.Index)
	}
}
