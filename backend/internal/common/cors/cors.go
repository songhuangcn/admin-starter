package cors

import (
	"github.com/gin-gonic/gin"
)

// gin-contrib/cors 插件有问题，因此自己写
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept-Language")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		}
	}
}
