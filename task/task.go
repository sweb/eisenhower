package task

import (
	"errors"
	"time"
)

// Type task is used to store information about tasks, like a description or
// the current status.
type Task struct {
	Id          int
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

func NewMinimalTask() *Task {
	return &Task{Status: "Collect", Created: time.Now(), LastChange: time.Now()}
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

func (t *Task) CreatedString() string {
	return timeString(t.Created)
}

func (t *Task) LastChangeString() string {
	return timeString(t.LastChange)
}

// Returns the due time formatted as string without time, just as date.
func (t *Task) DueTimeString() string {
	return dateString(t.DueTime)
}

func timeString(ti time.Time) string {
	return ti.Format("02.01.2006 15:04")
}

func dateString(ti time.Time) string {
	return ti.Format("02.01.2006")
}

// Returns boolean values in a more readable fashion.
func (t *Task) ImportantString() string {
	if t.Important {
		return "Yes"
	}
	return "No"
}

// Sets the due date by using a given string formatted date. If the date is
// malformed an error is returned. The date format is "02.01.2006"
func (t *Task) setDueTime(newTime string) error {
	parsedTime, err := time.ParseInLocation("02.01.2006", newTime, time.Local)
	if err != nil {
		return err
	}
	t.DueTime = parsedTime
	return nil
}

func (t *Task) setImportantFlag(important string) error {
	if important == "Yes" {
		t.Important = true
		return nil
	}
	if important == "No" {
		t.Important = false
		return nil
	}
	return errors.New("No acceptable boolean flag...")
}

func (t *Task) Update(title, description, important, dueTime string) {
	t.Title = title
	t.Description = description
	t.setImportantFlag(important)
	t.setDueTime(dueTime)
	t.LastChange = time.Now()
}
