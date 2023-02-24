package tick

import (
	"fmt"
	"time"
	"wally/task"
)

// 示例任务: tick 1/s
func BuildTask() task.Task {
	return task.Task{
		ID:   jobIdTick,
		Spec: "0/1 * * * * *",
		Cmd:  tick,
	}
}

func tick() {
	fmt.Println(time.Now(), "tick")
}
