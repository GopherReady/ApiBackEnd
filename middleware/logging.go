package middleware

import (
	"time"

	"github.com/GopherReady/ApiBackEnd/global"
	"github.com/gin-gonic/gin"
)

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		// *zap.Logger配置
		// logger.Info(path,
		// 	zap.Int("status", c.Writer.Status()),
		// 	zap.String("method", c.Request.Method),
		// 	zap.String("path", path),
		// 	zap.String("query", query),
		// 	zap.String("ip", c.ClientIP()),
		// 	zap.String("user-agent", c.Request.UserAgent()),
		// 	zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		// 	zap.Duration("cost", cost),
		// )
		global.Logger.Infof("url %s statusCode: %d method: %s query: %s ip: %s user-agent: %s"+
			" errors: %s cost %s", path, c.Writer.Status(), c.Request.Method, query, c.ClientIP(), c.Request.UserAgent(),
			c.Errors.ByType(gin.ErrorTypePrivate).String(), cost)
	}

}
