package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Europe/Berlin")
	ti, _ := time.ParseInLocation("02.Jan.2006", "03.Feb.2013", loc)
	t := NewTask("Test", "Testdesc", true, ti)
	fmt.Printf("Hello, world.\n%v", t)

	tl, err := loadTasks()
	if err != nil {
		fmt.Println("error:", err)
	}
	err = tl.saveTasks()

	if err != nil {
		fmt.Println("error:", err)
	}
}
