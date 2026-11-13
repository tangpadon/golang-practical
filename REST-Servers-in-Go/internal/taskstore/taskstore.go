package taskstore

import (
	"sync"
	"time"
)

type Task struct {
	ID   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

type TaskStore struct {
	mu     sync.Mutex
	tasks  map[int]Task
	nextID int
}

func New() *TaskStore {
	ts := &TaskStore{}
	ts.tasks = make(map[int]Task)
	ts.nextID = 0
	return ts
}

func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int {
	ts.Lock()
	defer ts.Unlock()

	task := Task{
		ID:   ts.nextID,
		Text: text,
		Due:  due,
	}
	task.Tags = make([]string, len(tags))
	copy(task.Tags, tags)

	ts.tasks[ts.nextID] = task
	ts.nextID++
	return task.ID
}

func (ts *TaskStore) GetTask(id int) (Task, error)

func (ts *TaskStore) DeleteTask(id int) error

func (ts *TaskStore) DeleteAllTasks() error

func (ts *TaskStore) GetAllTasks() []Task

func (ts *TaskStore) GetTasksByTag(tag string) []Task

func (ts *TaskStore) GetTasksByDueDate(due time.Time) []Task
