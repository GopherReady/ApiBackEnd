package initialize

import (
	"net/http"

	"github.com/GopherReady/ApiBackEnd/api/health"
	"github.com/GopherReady/ApiBackEnd/api/user"
	"github.com/GopherReady/ApiBackEnd/global"
	"github.com/GopherReady/ApiBackEnd/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RouterInitialize() {
	// gin 有 3 种运行模式：debug、release 和 test，其中 debug 模式会打印很多 debug 信息。
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	// 	middlewares
	g.Use(gin.Recovery())
	g.Use(middleware.ZapLogger())
	g.Use(middleware.Cors())
	g.Use(middleware.RequestId())
	g.Use(middleware.NoCache, middleware.Options, middleware.Secure)
	// g.Use(mw...)
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

	u := g.Group("/v1/user")
	{
		u.POST("", user.Create)       // 创建用户
		u.DELETE("/:id", user.Delete) // 删除用户
		u.PUT("/:id", user.Update)    // 更新用户
		u.GET("", user.List)          // 用户列表
		u.GET("/:username", user.Get) // 获取指定用户的详细信息
	}

	global.Logger.Infof("Start to listening the incoming requests on http address %s", viper.GetString("addr"))
	global.Logger.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())

}
