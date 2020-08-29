package test

import (
	"fmt"
	"net/http"
	"server-demo/config"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

// HandlePingMysql handle /test/url
func HandlePingMysql(ctx *gin.Context) {
	msg := "failed"

	conf := config.GetGlobalConfig()
	url := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.MysqlUser, conf.MysqlPasswd, conf.MysqlHost, conf.MysqlPort, conf.MysqlDB)
	db, err := gorm.Open("mysql", url)
	defer db.Close()

	if err == nil {
		err = db.DB().Ping()
		if err == nil {
			msg = "PONG"
		} else {
			fmt.Println("mysql ping error:", err)
		}
	} else {
		fmt.Println("mysql open have error:", err)
	}

	tabs := []string{}
	err = db.Raw("show tables").Pluck("Tables_in_mysql", &tabs).Error
	if err != nil {
		fmt.Println("mysql show tables had error:", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":    msg,
		"tables": tabs,
	})
}
