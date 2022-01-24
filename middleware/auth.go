package middleware

import (
	"github.com/GopherReady/ApiBackEnd/api/response"
	"github.com/GopherReady/ApiBackEnd/pkg/errno"
	"github.com/GopherReady/ApiBackEnd/pkg/token"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			response.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
