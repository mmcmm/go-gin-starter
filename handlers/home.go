package handlers

import (
	"github.com/gin-gonic/gin"
)

// Index... page
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
