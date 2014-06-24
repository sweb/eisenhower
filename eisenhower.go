package main

import (
	"flag"
	"github.com/sweb/eisenhower/task"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
)

var (
	addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)

var tl = &task.TaskList{}
var templates = template.Must(template.ParseFiles("resources/tmpl/edit_task.html",
	"resources/tmpl/view_task.html", "resources/tmpl/task_list.html"))

var validPath = regexp.MustCompile("^/tasks/(edit|save|view)/([0-9]+)$")

func main() {

	var err error = nil
	tl, err = task.InitTaskList()
	if err != nil {
		log.Println("error:", err)
	}
	flag.Parse()
	http.HandleFunc("/tasks", handler)
	http.HandleFunc("/tasks/view/", makeHandler(viewHandler))
	http.HandleFunc("/tasks/edit/", makeHandler(editHandler))
	http.HandleFunc("/tasks/save/", makeHandler(saveHandler))
	http.HandleFunc("/tasks/add/", addTaskHandler)

	if *addr {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
		if err != nil {
			log.Fatal(err)
		}
		s := &http.Server{}
		s.Serve(l)
		return
	}
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "task_list.html", tl)
}

func viewHandler(w http.ResponseWriter, r *http.Request, taskId string) {
	p, err := tl.TaskById(taskId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		renderTemplate(w, "view_task", p)
	}
}

func editHandler(w http.ResponseWriter, r *http.Request, taskId string) {
	editTask, err := tl.TaskById(taskId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "edit_task", editTask)
}

func saveHandler(w http.ResponseWriter, r *http.Request, taskId string) {
	tId := taskId
	if taskId == "0" {
		tId = tl.AddTask(task.NewMinimalTask())
	}
	t, err := tl.TaskById(tId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Update(r.FormValue("title"), r.FormValue("description"),
		r.FormValue("important"), r.FormValue("dueTime"))
	err = tl.SaveTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/tasks/view/"+tId, http.StatusFound)
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	t := task.NewMinimalTask()
	renderTemplate(w, "edit_task", t)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, t interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
