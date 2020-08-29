package routers

import (
	"server-demo/config"
	"server-demo/handlers/feedback"
	"server-demo/handlers/myauth"
	"server-demo/handlers/table"
	"server-demo/handlers/test"
	"server-demo/handlers/upload"
	"server-demo/handlers/user"
	"server-demo/handlers/websocket"

	"github.com/gin-gonic/gin"
)

// Register router
func Register(engine *gin.Engine) {
	conf := config.GetGlobalConfig()

	testGroup := engine.Group(conf.Prefix + "/test")
	{
		testGroup.GET("/ping", test.HandlePing()...)

		testGroup.GET("/ip", test.HandleIP)
		testGroup.GET("/url", test.HandleURL)
		testGroup.GET("/header", test.HandlerHeader)

		testGroup.GET("/time", test.HandleSleep)
		testGroup.GET("/sleep/:second", test.HandleSleep)

		testGroup.GET("/cookie", test.HandleCookie)

		testGroup.GET("/cache-control", test.HandleCacheControl)
		testGroup.GET("/match", test.HandleMatch)

		testGroup.GET("/mysql/ping", test.HandlePingMysql)
		testGroup.GET("/redis/ping", test.HandlePingRedis)

		testGroup.POST("/post", test.HandlePostJSON)
	}

	apiGroup := engine.Group(conf.Prefix + "/user")
	{
		apiGroup.POST("/login", user.HandleLogin)
		apiGroup.POST("/logout", user.HandleLogout)
		apiGroup.GET("/info", user.HandleInfo)
	}

	tableGroup := engine.Group(conf.Prefix + "/table")
	{
		tableGroup.GET("/list", table.HandleTableList)
	}

	feedbackGroup := engine.Group(conf.Prefix + "/feedback")
	{
		feedbackGroup.GET("/list", feedback.HandleFeedbackList)
	}

	// web-socket
	engine.GET("/echo", websocket.HandleWSEcho)

	// download static file
	engine.Static("/data", "./data")

	// zip on the fly
	engine.GET("/zip1", test.ZipOnFly1)

	// zip on the fly
	engine.GET("/zip2", test.ZipOnFly2)

	uploadGroup := engine.Group(conf.Prefix + "/upload")
	{
		// upload single file
		uploadGroup.POST("/single", upload.HandleSingalFile)

		// upload multiple files
		uploadGroup.POST("/multiple", upload.HandleMultipleFile)
	}

	// Group using gin.BasicAuth() middleware
	authorized := engine.Group("/auth", myauth.BasicAuth())
	{
		authorized.GET("/admin", myauth.HandleAuthAdmin)

		// static file with basic auth
		authorized.Static("/data", "./data")
	}
}
