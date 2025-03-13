package configs

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

type RedisConnection struct {
	Client			*redis.Client
	CacheStorage	*cache.Cache
}

func NewRedisConnection(settings *Settings) *RedisConnection {
	redisConnection := &RedisConnection{}

	redisUrl := fmt.Sprintf(
		"redis://%v:%v@%v:%v/%v?protocol=%v",
		settings.RedisUser,
		settings.RedisPassword,
		settings.RedisHost,
		settings.RedisPort,
		settings.RedisDB,
		settings.RedisProtocol,
	)
	opts, err := redis.ParseURL(redisUrl)

	if err != nil {
		log.Fatalln(err)
	}

	rdb := redis.NewClient(opts)

	myCache := cache.New(&cache.Options{
        Redis:      rdb,
        LocalCache: cache.NewTinyLFU(1000, time.Minute * 5),
    })

	redisConnection.Client = rdb
	redisConnection.CacheStorage = myCache

	log.Println("========== *** Redis connected. Cache storage initialized. *** ==========")

	return redisConnection
}
