package worker

import (
	"orchestrator/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	//collect stats
}

func (w *Worker) RunTask() {
	//start or stop task
}

func (w *Worker) StartTask() {
	//start task
}

func (w *Worker) StopTask() {
	//stop task
}
