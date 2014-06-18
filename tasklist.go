package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TaskList struct {
	Tasks []*Task
}

func (tl *TaskList) encodeTasksToJson() []byte {
	b, err := json.Marshal(tl)
	if err != nil {
		fmt.Println("error:", err)
	}
	return b
}

func (tl *TaskList) saveTasks() error {
	filename := "data/taskstore.json"
	return ioutil.WriteFile(filename, tl.encodeTasksToJson(), 0600)

}

func decodeTasksFromJson(b []byte, tl *TaskList) error {
	return json.Unmarshal(b, tl)
}

func loadTasks() (*TaskList, error) {
	filename := "data/taskstore.json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	tl := &TaskList{}
	err = decodeTasksFromJson(body, tl)
	return tl, nil
}
