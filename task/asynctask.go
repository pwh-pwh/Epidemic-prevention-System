package task

var pool *Pool = Default()

func AddTask(task func()) {
	pool.NewTask(task)
}
