package hello

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//Time struct comment blah blah
type Time struct {
	Name string
	Time time.Time
}

//Times array because Time objects have to go somewhere
type Times []Time

//GetTime does.. well... it gets the time
func GetTime(w http.ResponseWriter, r *http.Request) {

	t := Times{
		Time{Name: "Kodee"},
		Time{Name: "Kdizzle"},
	}

	json.NewEncoder(w).Encode(t)

	fmt.Print("Time for fun")

}
