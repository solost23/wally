package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/solost23/go_interface/gen_go/oss"
	"github.com/solost23/go_interface/gen_go/push"
	"gorm.io/gorm"
	"time"
	"wally/configs"
)

var (
	DB            *gorm.DB
	Loc           *time.Location
	ServerConfig  = &configs.ServerConfig{}
	RedisMapPool  = make(map[int]*redis.Client)
	PushSrvClient push.PushClient
	OSSSrvClient  oss.OssClient
)
