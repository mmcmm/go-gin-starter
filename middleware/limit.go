package middleware

import (
	"github.com/gin-gonic/gin"
)

// LimitMax simultaneous connections
func LimitMax(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire() // before request
		c.Next()
		release() // after request
	}
}
