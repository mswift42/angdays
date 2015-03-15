package angdays

import (
	"appengine"
	"appengine/datastore"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Task struct {
	Id        int64     `json:"id" datastore:"-"`
	Summary   string    `json:"summary"`
	User      string    `json:"user"`
	Content   string    `json:"content" datastore:",noindex"`
	Scheduled time.Time `json:"scheduled"`
	Done      string    `json:"done"`
	Category  string    `json:"category"`
}

func keyForID(c appengine.Context, id int64) *datastore.Key {
	return datastore.NewKey(c, "Task", "", id, tasklistkey(c))
}

// parseTime - convert a time string with layout
// dd/mm/yyyy to time.Time type.
func parseTime(s string) time.Time {
	layout := "02/01/2006"
	t, _ := time.Parse(layout, s)
	return t
}

// formatDate - convert a time.Time type
// to a string with layout dd/mm/yyyy
func formatDate(t time.Time) string {
	layout := "02/01/2006"
	return t.Format(layout)
}

func tasklistkey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Task", "default_tasklist", 0, nil)

}
func (t *Task) key(c appengine.Context) *datastore.Key {
	if t.Id == 0 {
		return datastore.NewIncompleteKey(c, "Task", tasklistkey(c))
	}
	return datastore.NewKey(c, "Task", "", t.Id, tasklistkey(c))
}

func (t *Task) save(c appengine.Context) (*Task, error) {
	k, err := datastore.Put(c, t.key(c), t)
	if err != nil {
		return nil, err
	}
	t.Id = k.IntID()
	return t, nil
}
func decodeTask(r io.ReadCloser) (*Task, error) {
	defer r.Close()
	var task Task
	err := json.NewDecoder(r).Decode(&task)
	return &task, err
}
func listTasks(c appengine.Context) (*[]Task, error) {
	tasks := []Task{}
	keys, err := datastore.NewQuery("Task").Ancestor(tasklistkey(c)).Order("-Done").Order("Scheduled").GetAll(c, &tasks)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(tasks); i++ {
		tasks[i].Id = keys[i].IntID()
	}
	return &tasks, err
}

func listTask(c appengine.Context, id int64) (*Task, error) {
	tasks := []Task{}
	_, err := datastore.NewQuery("Task").Ancestor(tasklistkey(c)).Filter("Id =", id).GetAll(c, &tasks)
	if err != nil {
		return nil, err
	}
	return &tasks[0], err
}
func (t *Task) delete(c appengine.Context) error {
	return datastore.Delete(c, t.key(c))
}

func init() {
	// router := httprouter.New()
	// router.GET("/tasks", handler)
	// router.POST("/tasks", handler)
	// router.DELETE("/tasks", deleteTaskHandler)
	// router.PATCH("/tasks/:taskid", updateTaskHandler)
	// http.HandleFunc("/tasks", handler)
	// router.HandleFunc("/tasks/user/", tasksHandler)
	// router.HandleFunc("/tasks/{id}", deleteTaskHandler).Methods("DELETE")
	// router.HandleFunc("/tasks", handler)
	// r := mux.NewRouter().StrictSlash(true)
	// tasks := r.Path("/tasks").Subrouter()
	// tasks.Methods("GET").HandlerFunc(handler)
	// tasks.Methods("POST").HandlerFunc(handler)
	// task := r.PathPrefix("/tasks/{id}").Subrouter()
	// task.Methods("DELETE").HandlerFunc(deleteTaskHandler)
	// http.Handle("/tasks", r)
	r := mux.NewRouter().PathPrefix("/api/").Subrouter()

	r.HandleFunc("/tasks", handler).Methods("GET")
	r.HandleFunc("/tasks", handler).Methods("POST")
	r.HandleFunc("/tasks/{id}", deleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/tasks/{id}", updateTaskHandler).Methods("POST")
	http.Handle("/api/", r)

}
func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	idstring := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idstring, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	task := Task{}
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	oldtask := Task{}
	key := keyForID(c, id)
	err = datastore.Get(c, key, &oldtask)
	if err == datastore.ErrNoSuchEntity {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if _, err := datastore.Put(c, key, &task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	val, err := handleTasks(c, r)
	if err == nil {
		err = json.NewEncoder(w).Encode(val)
	}
	if err != nil {
		c.Errorf("task error: %#v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	stringid := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(stringid, 0, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := datastore.Delete(c, datastore.NewKey(c, "Task", "", id, tasklistkey(c))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	val, err := handleTasks(c, r)
	if err == nil {
		err = json.NewEncoder(w).Encode(val)
	}
	if err != nil {
		c.Errorf("task error: %#v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func listTaskHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	task, err := decodeTask(r.Body)
	if err == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if _, err := task.save(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func postTaskHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	task, err := decodeTask(r.Body)
	if err != nil {
		c.Errorf("task error: %#v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if _, err := task.save(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func handleTask(c appengine.Context, r *http.Request) (interface{}, error) {
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		return nil, err
	}
	lt, err := listTask(c, id)
	if err != nil {
		return nil, err
	}
	return lt, nil
}
func handleTasks(c appengine.Context, r *http.Request) (interface{}, error) {
	switch r.Method {
	case "POST":
		task, err := decodeTask(r.Body)
		if err != nil {
			return nil, err
		}
		return task.save(c)
	case "GET":
		return listTasks(c)
	}
	return nil, fmt.Errorf("method not implemented")
}
