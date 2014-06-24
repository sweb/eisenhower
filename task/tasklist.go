package task

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
)

// A slice of tasks
type TaskList struct {
	Tasks     []*Task
	idCounter int
}

// Encodes the complete task list to json
func (tl *TaskList) encodeTasksToJson() ([]byte, error) {
	return json.Marshal(tl)
}

// Saves the task list to the hard drive
func (tl *TaskList) SaveTasks() error {
	filename := "data/taskstore.json"
	encodedTaskList, err := tl.encodeTasksToJson()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, encodedTaskList, 0600)
}

// Decodes a byte slice containing json encoded data and stores it in a
// TaskList.
func decodeTasksFromJson(b []byte, tl *TaskList) error {
	return json.Unmarshal(b, tl)
}

// Loads a TaskList from the hard drive by decoding it.
func loadTasks() (*TaskList, error) {
	filename := "data/taskstore.json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	tl := &TaskList{}
	err = decodeTasksFromJson(body, tl)
	return tl, err
}

// Initializes the task list for example at server start.
func InitTaskList() (*TaskList, error) {
	tl, err := loadTasks()
	if err != nil {
		return &TaskList{}, err
	}
	tl.idCounter = maxCurrentTaskId(tl.Tasks)
	return tl, nil
}

// Returns the currently highest taskId to determine the next possible taskId.
// This function should only be necessary on server start, afterwards the
// counter should be synchronized.
func maxCurrentTaskId(tasks []*Task) (m int) {
	for _, task := range tasks {
		if task.Id > m {
			m = task.Id
		}
	}
	return
}

// Checks if tasks are available.
func (tl *TaskList) HasTasks() bool {
	return tl.idCounter > 0
}

func (tl *TaskList) TaskById(id string) (*Task, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	for _, task := range tl.Tasks {
		if task.Id == intId {
			return task, nil
		}
	}
	return nil, errors.New("Task not found...")
}

func (tl *TaskList) DeleteByTaskId(id int) error {
	taskId := -1
	for i := 0; i < len(tl.Tasks); i++ {
		if id == tl.Tasks[i].Id {
			taskId = i
			break
		}
	}
	if taskId == -1 {
		return errors.New("Task not found...")
	}
	copy(tl.Tasks[taskId:], tl.Tasks[taskId+1:])
	tl.Tasks[len(tl.Tasks)-1] = nil
	tl.Tasks = tl.Tasks[:len(tl.Tasks)-1]
	return nil
}

func (tl *TaskList) AddTask(task *Task) string {
	tl.idCounter++
	task.Id = tl.idCounter
	tl.Tasks = append(tl.Tasks, task)
	return strconv.Itoa(tl.idCounter)
}
