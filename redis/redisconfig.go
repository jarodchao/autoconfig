package redis

import (
	"github.com/go-redis/redis"
	"fmt"
	"time"
	"autoconfig/type"
)

var redisClient *redis.Client

func InitRedis(r autoconfig.Redis) {

	if r.Address != "" {
		url := r.ToUrl()

		options := redis.Options{}

		options.Addr = url
		options.DB = r.Db
		options.Password = r.Address

		options.PoolSize = r.Pool.MaxIdle
		options.PoolTimeout = time.Duration(r.Pool.PoolTimeout) * time.Microsecond

		redisClient = redis.NewClient(&options)
	}else {

		options := redis.FailoverOptions{}

		options.MasterName = r.Sentinel.Master
		options.SentinelAddrs = r.Sentinel.Nodes
		options.DB = r.Db
		options.Password = r.Password

		options.PoolSize = r.Pool.MaxIdle
		options.PoolTimeout = time.Duration(r.Pool.PoolTimeout) * time.Microsecond

		redisClient = redis.NewFailoverClient(&options)
	}

	result, err := redisClient.Ping().Result()

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Connect Redis Server Result is : %s\n", result)

}
