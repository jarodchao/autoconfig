package autoconfig

import (
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"autoconfig/type"

	"autoconfig/mysql"
	"autoconfig/redis"
)

const configPath string = "./conf/application.yml"

func init() {
	data, err := ioutil.ReadFile(configPath)

	if err != nil {
		return
	}

	//p("Application Config: %s", string(data))

	var appConfig autoconfig.AppConfig

	err1 := yaml.Unmarshal(data, &appConfig)

	if err1 != nil {
		log.Fatalf("error: %v", err)
	}

	if appConfig.Database.Url != "" {
		mysql.InitDB(appConfig.Database)
	}

	if appConfig.Redis.Address != "" || appConfig.Redis.Sentinel.Master != "" {
		redis.InitRedis(appConfig.Redis)
	}
}
