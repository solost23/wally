package tick

import (
	"fmt"
	"wally/task"
)

// 示例任务: tick/s
func BuildTask() task.Task {
	return task.Task{
		ID:   jobIdTick,
		Spec: "@every 1s",
		Cmd:  tick,
	}
}

func tick() {
	fmt.Println("tick")
}
