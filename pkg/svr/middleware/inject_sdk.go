package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golanggo/hzs-wecom"
)

func InjectSdk(ww wework.IWeWork) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ww", ww)
		c.Next()
	}
}
