package global

import (
	"gin-vue-admin/config"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	PantaDb     *gorm.DB
	PantaRedis  *redis.Client
	PantaConfig config.Server
	PantaVp     *viper.Viper
	PantaLog    *oplogging.Logger
)
