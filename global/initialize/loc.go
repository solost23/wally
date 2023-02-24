package initialize

import (
	"time"
	"wally/global"
)

func InitLoc() {
	global.Loc, _ = time.LoadLocation(global.ServerConfig.TimeLocation)
}
