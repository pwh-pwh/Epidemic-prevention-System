package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pwh-pwh/Epidemic-prevention-System/settings"
)

var (
	client *redis.Client
)

/*type SliceCmd = redis.SliceCmd
type StringStringMapCmd = redis.StringStringMapCmd*/

// Init 初始化连接
func InitRedis(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func GetRedisClient() *redis.Client {
	return client
}

func Close() {
	_ = client.Close()
}
