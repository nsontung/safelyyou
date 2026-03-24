package middleware

import (
	"fmt"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

// RequestLogger is a middleware that logs the incoming HTTP requests, including headers and body
// This is useful for debugging and monitoring purposes
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// DumpRequest clones the request, including headers and body
		requestDump, _ := httputil.DumpRequest(c.Request, true)
		fmt.Println(string(requestDump))
		c.Next()
	}
}
