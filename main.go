package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/GopherReady/ApiBackEnd/global"
	"github.com/GopherReady/ApiBackEnd/initialize"
	"github.com/jinzhu/gorm"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	// init config
	if err := initialize.InitViper(*cfg); err != nil {
		panic(err)
	}

	// init zap logger
	initialize.InitLogger()

	// init database
	initialize.InitGorm()
	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {
			global.Logger.Error("Gorm database closed failed", err)
		}
	}(initialize.DB)

	initialize.RouterInitialize()

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			global.Logger.Fatal("The vps has no response, or it might took too long to start up.", err)
		}
		global.Logger.Info("The vps has been deployed successfully.")
	}()

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
		global.Logger.Info("Waiting for the vps, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the vps")
}
