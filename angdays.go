package angdays

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"appengine"
	"appengine/datastore"
)

type Task struct {
	Id        int64     `json:"id" datastore:"-"`
	Summary   string    `json:"summary"`
	User      string    `json:"user"`
	Content   string    `json:"content" datastore:",noindex"`
	Scheduled time.Time `json:"scheduled"`
	Done      string    `json:"done"`
}

// Agenda - struct for Overview of upcoming tasks.
// Contains a date in format <Weekday, day, monthabbr year>
// and a slice of tasks for the date.
type Agenda struct {
	FancyDate string `json:"fancydate"`
	Taskslice []Task `json:"taskslice"`
}

type TaskAndAgenda struct {
	Tasks       []Task   `json:"tasks"`
	Agendaslice []Agenda `json:"agendaslice"`
}

func keyForID(c appengine.Context, id int64) *datastore.Key {
	return datastore.NewKey(c, "Task", "", id, tasklistkey(c))
}

// agendasize constant, describes size of agendaoverview in days
const (
	agendasize int64 = 10
)

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

func formatDateFancy(t time.Time) string {
	layout := "Monday, 02 Jan 2006"
	return t.Format(layout)
}

// weekDates - takes a day
// and returns a slice of dates of range startday - 10 days from startday.
func weekDates(day time.Time) []time.Time {
	week := make([]time.Time, agendasize)
	for i := int64(0); i < agendasize; i++ {
		week[i] = addDay(day, i-3)
	}

	return week
}

// addDay - add to a given starting day, a number of days
// and return the resulting date.
func addDay(startday time.Time, day int64) time.Time {
	length := 24 * day
	return startday.Add(time.Duration(length) * time.Hour)
}

// agendaOverview - takes a taskslice and a day
// and builds an overview of all coming dates in range of today -
// agendasize. For every day it builds a struct agenda with a
// formatted date and a slice of tasks, due at that date and with
// status "Todo". Finally the slice of 'Agendastructs is returned.
func agendaOverview(ts []Task, d time.Time) []Agenda {
	week := weekDates(d)
	a := make([]Agenda, agendasize)
	for i, j := range week {
		a[i].FancyDate = formatDateFancy(j)
	}
	for i := range a {
		ag := make([]Task, 0)
		for _, k := range ts {
			if formatDate(week[i]) == formatDate(k.Scheduled) && k.Done == "Todo" {
				ag = append(ag, k)
			}
		}
		a[i].Taskslice = ag
	}
	return a
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
	router := mux.NewRouter()
	router.HandleFunc("/tasks", handler)
	router.HandleFunc("/tasks/user/", tasksHandler)
	router.HandleFunc("/tasks/{id}/", listTaskHandler)
	router.HandleFunc("/tasks/{id}/", deleteTaskHandler).Methods("DELETE")
	http.Handle("/tasks", router)

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
	id, err := strconv.ParseInt(stringid, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := datastore.Delete(c, datastore.NewKey(c, "Task", "", id, tasklistkey(c))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	case "DELETE":
		task, err := decodeTask(r.Body)
		if err != nil {
			return nil, err
		}
		return nil, task.delete(c)
	}
	return nil, fmt.Errorf("method not implemented")
}
