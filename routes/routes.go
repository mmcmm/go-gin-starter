package routes

import (
	"github.com/gin-contrib/gzip"
	ratelimit "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	"github.com/mtdx/case-api/handlers"
	"github.com/mtdx/case-api/middleware"
	"github.com/mtdx/case-api/steamauth"
)

var router *gin.Engine

func initRoutes() {
	authMiddleware := middleware.Jwt()

	// limit simultaneous connections
	// router.Use(middleware.LimitMax(200))
	// Gzip
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	// POST max size, 2MB
	router.Use(ratelimit.RateLimiter(2 * 1024 * 1024))
	// Database
	router.Use(middleware.DbConnPool())

	router.GET("/login", steamauth.LoginHandler) // TODO: implement

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/", handlers.Index)
		restricted := apiv1.Group("/").Use(authMiddleware.MiddlewareFunc())
		{
			restricted.GET("/refresh-token", authMiddleware.RefreshHandler)
			restricted.GET("/authenticated", handlers.Authenticated)
		}

	}
}

// SetupRouter ...
func SetupRouter() *gin.Engine {
	router = gin.Default()
	initRoutes()

	return router
}
