package main

import (
  "fmt"
  "time"
)

type Task struct {
  Title string
  Description string
  Important bool
  Urgent bool
  Status string
  Created time.Time
  LastChange time.Time
}

func (t *Task) quadrant() int32 {
  switch {
  case t.Urgent && t.Important:
    return 1
  case !t.Urgent && t.Important:
    return 2
  case t.Urgent && !t.Important:
    return 3
  }
  return 4
}


func main() {
  t := &Task{Title: "Test", Description: "Testdesc", Important: true, 
    Urgent:false, Status: "Collect", Created: time.Now(), LastChange: time.Now()}
	fmt.Printf("Hello, world.\n" + t.Title)
}

