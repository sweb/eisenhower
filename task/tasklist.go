package task

import (
	"encoding/json"
	"io/ioutil"
)

// A slice of tasks
type TaskList struct {
	Tasks []*Task
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
func LoadTasks() (*TaskList, error) {
	filename := "data/taskstore.json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	tl := &TaskList{}
	err = decodeTasksFromJson(body, tl)
	return tl, nil
}
