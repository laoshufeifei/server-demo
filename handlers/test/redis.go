package test

import (
	"fmt"
	"net/http"
	"server-demo/config"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

// HandlePingRedis handle /test/url
func HandlePingRedis(ctx *gin.Context) {
	msg := "failed"

	conf := config.GetGlobalConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisURL,
		Password: conf.RedisPasswd,  // no password set
		DB:       conf.RedisDBIndex, // use default DB
	})

	msg, err := client.Ping().Result()
	if err != nil {
		fmt.Println("ping redis had error", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
