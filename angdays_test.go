package angdays

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
	task1 := Task{Summary: "task1", User: "test@example.com", Done: "Todo"}
	task2 := Task{Summary: "task2", Done: "Done", Scheduled: "05/07/2014"}
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
