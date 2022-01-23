package router

import (
	"net/http"

	"github.com/GopherReady/ApiBackEnd/handle/health"
	"github.com/GopherReady/ApiBackEnd/middleware"
	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 	middlewares
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache, middleware.Options, middleware.Secure)
	g.Use(middleware.ZapLogger())
	g.Use(mw...)
	// 404 handeler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "API router is not found")
	})

	// The health check handlers
	VPSHealth := g.Group("/vps")
	{
		VPSHealth.GET("/health", health.HealthCheck)
		VPSHealth.GET("/disk", health.DiskCheck)
		VPSHealth.GET("/cpu", health.CPUCheck)
		VPSHealth.GET("/ram", health.RAMCheck)
	}

	return g
}
