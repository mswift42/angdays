package angdays

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"appengine/aetest"
	"appengine/datastore"
	"appengine/user"
)

func TestTaskStruct(t *testing.T) {
	t1 := Task{Summary: "task1", User: "test@example.com", Done: "Todo"}
	if t1.User != "test@example.com" {
		t.Error("Expected <test@example.com>, got: ", t1.User)
	}
	if t1.Done != "Todo" {
		t.Error("Expected <Todo>, got: ", t1.Done)
	}
	if t1.Summary != "task1" {
		t.Error("Expected <task1>, got: ", t1.Summary)
	}

}

func TestAgendaStruct(t *testing.T) {
	task1 := Task{Summary: "task1", User: "test@example.com",
		Done: "Todo", Scheduled: parseTime("02/02/2014")}
	task2 := Task{Summary: "task2", Done: "Done",
		Scheduled: parseTime("04/05/2006")}
	tasks := []Task{task1, task2}
	t1 := Agenda{FancyDate: "Saturday, 05 Jul 2014", Taskslice: tasks}
	if t1.FancyDate != "Saturday, 05 Jul 2014" {
		t.Error("Expected <Saturday, 05 Jul 2014, got: ", t1.FancyDate)
	}
	tsklist := t1.Taskslice
	if tsklist[0].Summary != "task1" {
		t.Error("Exptedted <task1>, got: ", tsklist[0].Summary)
	}
}

func TestParseTime(t *testing.T) {
	assert := assert.New(t)
	time1 := "01/02/2003"
	time2 := "02/03/2004"
	t1 := parseTime(time1)
	t2 := parseTime(time2)
	assert.Equal(t1.Month(), 2, "t1.Month == February")
	assert.Equal(t2.Month(), time.March, "t2.Month === March")
}

func TestFormatDate(t *testing.T) {
	assert := assert.New(t)
	s1 := "01/02/2003"
	s2 := "03/04/2005"
	time1 := parseTime(s1)
	time2 := parseTime(s2)
	formatted := formatDate(time1)
	formatted2 := formatDate(time2)
	assert.Equal(formatted, "01/02/2003", "formattedDate = 01/02/2003")
	assert.Equal(formatted2, "03/04/2005", "formattedDate = 03/04/2005")
}

func TestTasks(t *testing.T) {
	t1 := Task{Summary: "task1", Content: "Some content", Done: "Done", Scheduled: parseTime("01/01/2012")}
	t2 := Task{Summary: "task2", Done: "Todo"}
	c, err := aetest.NewContext(nil)
	u := user.Current(c)
	if u != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	key := datastore.NewKey(c, "Task", "", 1, nil)
	if _, err := datastore.Put(c, key, &t1); err != nil {
		t.Fatal(err)
	}
	nkey := datastore.NewKey(c, "Task", "", 2, nil)
	if _, err := datastore.Put(c, nkey, &t2); err != nil {
		t.Fatal(err)
	}
	gt1 := Task{}
	gt2 := Task{}
	if err := datastore.Get(c, key, &gt1); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(gt1, t1) {
		t.Error("Expected gt1 == t1, got: ", gt1, t1)
	}

	if err := datastore.Get(c, nkey, &gt2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(gt2, t2) {
		t.Error("Expected gt2 == t2, got: ", gt2, t2)
	}

	defer c.Close()
}
