package initialize

import (
	"gin-vue-admin/global"
	"github.com/go-redis/redis"
)

func Redis() {
	redisCfg := global.PantaConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.PantaLog.Error(err)
	} else {
		global.PantaLog.Info("redis connect ping response:", pong)
		global.PantaRedis = client
	}
}
