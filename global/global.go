package global

import (
	"github.com/GopherReady/ApiBackEnd/config"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	Config config.Config
	DB     *gorm.DB
	Rdb    *redis.Client
	Logger *zap.SugaredLogger
)
