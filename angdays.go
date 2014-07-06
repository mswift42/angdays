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
func listTasks(c appengine.Context) (TaskAndAgenda, error) {
	tasks := []Task{}
	keys, err := datastore.NewQuery("Task").Ancestor(tasklistkey(c)).Order("-Done").Order("Scheduled").GetAll(c, &tasks)
	if err != nil {
		return TaskAndAgenda{}, err
	}
	for i := 0; i < len(tasks); i++ {
		tasks[i].Id = keys[i].IntID()
	}
	ag := agendaOverview(tasks, time.Now())
	taa := TaskAndAgenda{Tasks: tasks, Agendaslice: ag}
	return taa, err
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
		c.Errorf("task error: %#v", err)
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
		return listTasks(c)
		// case "DELETE":
		// 	return nil, deleteDoneTodos(c)
	}
	return nil, fmt.Errorf("method not implemented")
}
