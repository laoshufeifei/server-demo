package main

import (
	"fmt"
	"os"
	"server-demo/config"
	"server-demo/routers"
	"server-demo/utils"
	"time"

	"server-demo/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @basePath /
func main() {
	utils.IncreaseResourcesLimit()

	ginEngine := gin.New()

	file, _ := os.OpenFile("access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	c := gin.LoggerConfig{
		Output:    file,
		SkipPaths: []string{},
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("[%s] - %s \"%s %s %s %d %s %s\"\n",
				params.TimeStamp.Format(time.RFC3339),
				params.ClientIP,
				params.Method,
				params.Path,
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
				params.ErrorMessage,
				// param.Request.UserAgent(),
			)
		},
	}
	ginEngine.Use(gin.LoggerWithConfig(c))

	routers.Register(ginEngine)
	// 0.0.0.0:9090
	conf := config.GetGlobalConfig()
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	docs.SwaggerInfo.BasePath = conf.Prefix
	swaggerPrefix := conf.Prefix + "/swagger"
	swaggerURL := ginSwagger.URL(swaggerPrefix + "/doc.json")
	ginEngine.GET(swaggerPrefix+"/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL),
	)

	ginEngine.Run(addr)
}
