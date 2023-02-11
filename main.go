package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"Reward/common"
	"Reward/config"
	"Reward/log"
	"Reward/model"
)

func main() {
	log.Init()
	defer log.SyncLogger()

	if err := config.Init(common.StringEmpry); err != nil {
		panic(err)
	}

	if err := model.InitDatabase(); err != nil {
		panic(err)
	}

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.",
				zap.String("reason", err.Error()))
		}
		log.Info("The router has been deployed successfully.")
	}()

	g := gin.New()
}

func pingServer() error {
	times := viper.GetInt("max_ping_count")
	for i := 0; i < times; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}
