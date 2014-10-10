package angdays

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"code.google.com/p/appengine-go/appengine/aetest"

	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {
	assert := assert.New(t)
	time1 := "01/02/2003"
	time2 := "02/03/2004"
	t1, err := parseTime(time1)
	if err != nil {
		t.Error("Expected nil:, got", err)
	}
	t2, _ := parseTime(time2)
	assert.Equal(t1.Month(), time.February)
	assert.Equal(t2.Month(), time.March)
	assert.Equal(t1.Day(), 1)
	assert.Equal(t2.Day(), 2)
}

func TestFormatDate(t *testing.T) {
	assert := assert.New(t)
	var times = []string{"01/02/2003",
		"02/04/2005", "31/12/2014",
		"24/12/2001", "01/01/2000",
		"01/03/2014", "30/06/2008"}
	for _, i := range times {
		t1, err := parseTime(i)
		if err != nil {
			t.Error("Expected nil, got: ", err)
		}
		assert.Equal(formatDate(t1), i)
	}
}

var fancyDates = []struct {
	normal string
	fancy  string
}{
	{"10/05/2014", "Saturday, 10 May 2014"},
	{"03/09/2014", "Wednesday, 03 Sep 2014"},
	{"04/08/2014", "Monday, 04 Aug 2014"},
	{"25/04/2010", "Sunday, 25 Apr 2010"},
	{"15/09/2011", "Thursday, 15 Sep 2011"},
	{"30/12/2011", "Friday, 30 Dec 2011"},
}

func TestFormatDateFancy(t *testing.T) {
	assert := assert.New(t)
	for _, i := range fancyDates {
		t1, err := parseTime(i.normal)
		if err != nil {
			t.Error("Expected nil, got: ", err)
		}
		assert.Equal(i.fancy, formatDateFancy(t1))
	}
}
func TestHandler(t *testing.T) {
	c, cerr := aetest.NewContext(nil)
	if cerr != nil {
		t.Fatal(cerr)
	}
	resp := httptest.NewRecorder()
	uri := "/tasks"
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}
	http.DefaultServeMux.ServeHTTP(c, req)
	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else {
		if strings.Contains(string(p), "Error") {
			t.Errorf("header response shoudn't return error: %s", p)
		}
	}

}
