package manager

import (
	"orchestrator/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending        queue.Queue
	TaskDb         map[string][]task.Task
	EventDb        map[string][]task.TaskEvent
	Workers        []string
	WorkersTaskMap map[string][]uuid.UUID
	TaskWorkerMap  map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	//select appropriate worker
}

func (m *Manager) UpdateTasks() {
	//update tasks
}

func (m *Manager) SendWork() {
	//send work to workers
}
