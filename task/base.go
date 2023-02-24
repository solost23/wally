package task

type Cmd func()

type Task struct {
	ID   string
	Spec string
	Cmd  Cmd
}
