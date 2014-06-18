package main

import (
	"fmt"
	"github.com/sweb/eisenhower/task"
	"net/http"
)

var tl = &task.TaskList{}

func main() {

	var err error = nil
	tl, err = task.LoadTasks()
	if err != nil {
		fmt.Println("error:", err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	/*err = tl.SaveTasks()

	if err != nil {
		fmt.Println("error:", err)
	}*/
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My tasks:\n")
	for _, value := range tl.Tasks {
		fmt.Fprintf(w, "%v\n", value)
	}

}
