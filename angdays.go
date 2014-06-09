package angdays

import (
	"fmt"
	"net/http"
)

type Task struct {
	Id        int64
	User      string
	Content   string
	Scheduled string
	Done      bool
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
