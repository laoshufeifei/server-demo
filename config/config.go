package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// ServerConfigs is the config for the server
type ServerConfigs struct {
	Prefix string `json:"prefix"`

	Host string `json:"host"`
	Port int    `json:"port"`

	MysqlHost   string `json:"mysql_host"`
	MysqlPort   int    `json:"mysql_port"`
	MysqlUser   string `json:"mysql_user"`
	MysqlDB     string `json:"mysql_db"`
	MysqlPasswd string `json:"mysql_passwd"`

	RedisURL     string `json:"redis_url"`
	RedisPasswd  string `json:"redis_passwd"`
	RedisDBIndex int    `json:"redis_db_index"`
}

var config ServerConfigs

func init() {
	log.Println("config.init")
	cfgPath := "config.json"
	_, err := os.Stat(cfgPath)
	if err == nil {
		contens, _ := ioutil.ReadFile(cfgPath)
		json.Unmarshal([]byte(contens), &config)

		config.Prefix = strings.TrimRight(config.Prefix, "/")
	} else {
		fmt.Printf("%s does not existing, will use default config\n", cfgPath)
		config.Prefix = ""
		config.Host = "0.0.0.0"
		config.Port = 9090
		config.MysqlHost = "127.0.0.1"
		config.MysqlPort = 3306
		config.MysqlUser = "root"
		config.MysqlDB = "test"
	}
}

// GetGlobalConfig return the global config
func GetGlobalConfig() *ServerConfigs {
	return &config
}
