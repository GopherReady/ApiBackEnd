package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/GopherReady/ApiBackEnd/config"
	"github.com/GopherReady/ApiBackEnd/model"
	"github.com/GopherReady/ApiBackEnd/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// init zap logger
	logger := config.InitLogger()
	// init database
	model.DB.Init()
	defer model.DB.Close()

	// gin 有 3 种运行模式：debug、release 和 test，其中 debug 模式会打印很多 debug 信息。
	gin.SetMode(viper.GetString("runmode"))
	client := gin.New()

	var middleware []gin.HandlerFunc

	// 	Routes
	router.Load(
		// Cores
		client,
		// Middleware load
		middleware...,
	)
	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The vps has no response, or it might took too long to start up.", err)
		}
		log.Print("The vps has been deployed successfully.")
	}()

	// logger.Info("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	logger.Info("Start to listening the incoming requests on http address", zap.String("addr", viper.GetString("addr")))
	logger.Info(http.ListenAndServe(viper.GetString("addr"), client).Error())

}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/vps/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the vps, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the vps.")
}
