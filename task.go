package main

import "time"

type Task struct {
	Title       string
	Description string
	Important   bool
	Status      string
	Created     time.Time
	LastChange  time.Time
	DueTime     time.Time
}

func NewTask(title, description string, important bool, dueTime time.Time) *Task {
	return &Task{Title: title, Description: description, Important: important,
		Status: "Collect", Created: time.Now(), LastChange: time.Now(), DueTime: dueTime}
}

func (t *Task) Urgent() bool {
	whenTodayGetsUrgent := time.Now().Add(24 * time.Hour)
	return whenTodayGetsUrgent.After(t.DueTime)
}

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
