package angdays

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Task struct {
	Id        int64     `json:"id" datastore:"-"`
	User      string    `json:"user"`
	Summary   string    `json:"summary"`
	Content   string    `json:"content"`
	Scheduled time.Time `json:"scheduled"`
	Done      bool      `json:"done"`
}

type Agenda struct {
	FancyDate string `json:"fancydate"`
	Taskslice []Task `json:"taskslice"`
}

type TaskAndAgenda struct {
	Tasks       []Task   `json:"tasks"`
	Agendaslice []Agenda `json:"agendaslice"`
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/tasks", tasksHandler)
	r.HandleFunc("/tasks/{id}", taskHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "everything ok")
}
