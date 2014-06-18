package task

import (
	"encoding/json"
	"fmt"
	"time"
)

// Type task is used to store information about tasks, like a description or
// the current status.
type Task struct {
	Title       string
	Description string
	Important   bool
	Status      string
	Created     time.Time
	LastChange  time.Time
	DueTime     time.Time
}

// Constructor for a new task. Created and last change are set to now, the first
// status is Collect
func NewTask(title, description string, important bool, dueTime time.Time) *Task {
	return &Task{Title: title, Description: description, Important: important,
		Status: "Collect", Created: time.Now(), LastChange: time.Now(), DueTime: dueTime}
}

// Checks if the task is already urgent by comparing the due date with the
// current date plus two day (hard coded at the moment)
func (t *Task) Urgent() bool {
	whenTodayGetsUrgent := time.Now().Add(24 * time.Hour)
	return whenTodayGetsUrgent.After(t.DueTime)
}

// Evaluates the quadrant of the task, concerning the eisenhower matrix.
func (t *Task) Quadrant() int {
	switch {
	case t.Urgent() && t.Important:
		return 1
	case !t.Urgent() && t.Important:
		return 2
	case t.Urgent() && !t.Important:
		return 3
	}
	return 4
}

// Encodes to json...may not be necessary.
func (t *Task) encodeToJson() []byte {
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println("error:", err)
	}
	return b
}
