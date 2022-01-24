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
		global.Logger.Infof("Client query wrong router link %s", c.Request.URL)
	})

	// The health check handlers
	VPSHealth := g.Group("/vps")
	{
		VPSHealth.GET("/health", health.HealthCheck)
		VPSHealth.GET("/disk", health.DiskCheck)
		VPSHealth.GET("/cpu", health.CPUCheck)
		VPSHealth.GET("/ram", health.RAMCheck)
	}
	// api for authentication functionalities
	u := g.Group("/v1/user")
	u.POST("/login", user.Login)     // 用户注册获得签名认证
	u.POST("/register", user.Create) // 创建用户
	u.Use(middleware.JwtAuth())
	{
		u.DELETE("/:id", user.Delete) // 删除用户
		u.PUT("/:id", user.Update)    // 更新用户
		u.GET("", user.List)          // 用户列表
		u.GET("/:username", user.Get) // 获取指定用户的详细信息
	}

	// Start ListenAndServeTLS HTTPS requests
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			global.Logger.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
			global.Logger.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}
	global.Logger.Infof("Start to listening the incoming requests on http address %s", viper.GetString("addr"))
	global.Logger.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}
