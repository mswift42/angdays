package angdays

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
)

// Task - struct for datastore table.
// Contains a summary and the contents of a task, the scheduled
// time for the task and whether it is done or not.
type Task struct {
	Summary   string `json:"summary"`
	Content   string `json:"content" datastore:",noindex"`
	Scheduled string `json:"scheduled"`
	Done      string `json:"done"`
	Id        int64  `json:"id" datastore:"-"`
}

// parseTime - convert a time string with layout
// dd/mm/yyyy to time.Time type.
func parseTime(s string) (time.Time, error) {
	layout := "02/01/2006"
	t, err := time.Parse(layout, s)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// formatDate - convert a time.Time type
// to a string with layout dd/mm/yyyy
func formatDate(t time.Time) string {
	layout := "02/01/2006"
	return t.Format(layout)
}

func formatDateFancy(t time.Time) string {
	layout := "Monday, 02 Jan 2006"
	return t.Format(layout)
}

func defaultTaskList(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Task", "default_tasklist", 0, nil)

}
func init() {
	http.HandleFunc("/tasks", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	val, err := handleTasks(c, r)
	if err == nil {
		err = json.NewEncoder(w).Encode(val)
	}
	if err != nil {
		c.Errorf("task errof: %#v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

		return getAllTasks(c)
		// case "DELETE":
		// 	return nil, deleteDoneTasks(c)
	}
	return nil, fmt.Errorf("method not implemented")
}

func (t *Task) key(c appengine.Context) *datastore.Key {
	if t.Id == 0 {
		return datastore.NewIncompleteKey(c, "Task", defaultTaskList(c))
	}
	return datastore.NewKey(c, "Task", "", t.Id, defaultTaskList(c))
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

func getAllTasks(c appengine.Context) ([]Task, error) {
	tasks := []Task{}
	ks, err := datastore.NewQuery("Task").Ancestor(defaultTaskList(c)).Order("Done").Order("-Scheduled").GetAll(c, &tasks)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(tasks); i++ {
		tasks[i].Id = ks[i].IntID()
	}
	return tasks, nil
}
