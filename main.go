package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mtdx/case-api/routes"
)

func main() {
	
	db := db.Init()
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)
	router := routes.SetupRouter()
	router.Run()
}
