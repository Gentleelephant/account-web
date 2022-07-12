package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CrossDomain(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS,Patch")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,X-Requested-With,X-CSRF-Token,X-token,Token")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}
