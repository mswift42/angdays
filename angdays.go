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

type Task struct {
	Id        int64     `json:"id" datastore:"-"`
	Summary   string    `json:"summary"`
	User      string    `json:"user"`
	Content   string    `json:"content" datastore:",noindex"`
	Scheduled string    `json:"scheduled"`
	Done      bool      `json:"done"`
	Created   time.Time `json:"created"`
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
func tasklistkey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Task", "default_tasklist", 0, nil)

}
func (t *Task) key(c appengine.Context) *datastore.Key {
	if t.Id == 0 {
		t.Created = time.Now()
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
func listTasks(c appengine.Context) ([]Task, error) {
	tasks := []Task{}
	keys, err := datastore.NewQuery("Task").Ancestor(tasklistkey(c)).Order("-Done").Order("Scheduled").GetAll(c, &tasks)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(tasks); i++ {
		tasks[i].Id = keys[i].IntID()
	}
	return tasks, err
}