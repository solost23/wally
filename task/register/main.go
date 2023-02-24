package register

import (
	"github.com/robfig/cron/v3"
	"wally/mission/tick"
	"wally/task"
)

// 初始化任务列表
var tasks = []task.Task{
	// 注册任务 tick
	tick.BuildTask(),
}

// 注册任务
func RegisterTasks(c *cron.Cron) (err error) {
	for _, task := range tasks {
		_, err = c.AddFunc(task.Spec, task.Cmd)
		if err != nil {
			return err
		}
	}
	return nil
}
