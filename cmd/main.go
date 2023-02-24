package main

import (
	"flag"
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"syscall"
	"wally/global"
	"wally/global/initialize"
	"wally/task/register"
)

var (
	WebConfigPath = "configs/config.yml"
	version       = "__BUILD_VERSION_"
	execDir       string
	provider      string
	st, v, V      bool
)

func main() {
	flag.StringVar(&execDir, "d", ".", "项目目录")
	flag.StringVar(&provider, "p", "consul", "项目配置提供者")
	flag.BoolVar(&v, "v", false, "查看版本号")
	flag.BoolVar(&V, "V", false, "查看版本号")
	flag.BoolVar(&st, "s", false, "项目状态")
	flag.Parse()

	if v || V {
		fmt.Println(version)
		return
	}
	initialize.Initialize(WebConfigPath)
	// 初始化对象
	t := cron.New(cron.WithSeconds(), cron.WithLocation(global.Loc))
	// 注册任务
	err := register.RegisterTasks(t)
	if err != nil {
		fmt.Println("Register task failed!, err: ", err.Error())
		return
	}
	// 开启定时任务
	t.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-c
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

}
