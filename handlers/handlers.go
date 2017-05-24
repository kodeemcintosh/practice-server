package handlers

import (
	"fmt"
	"net/http"
	"time"
)

type Welcome struct {
	Greet string
}

type Time struct {
	Name string
	Time time.Time
}

type Times []Time

func Index(w http.ResponseWriter, r *http.Request) {
	// http.FileServer(http.Dir("./ng2-app/")).ServeHTTP(w, r)
	http.ServeFile(w, r, "ng2-app/index.html")

}

func Hello(w http.ResponseWriter, r *http.Request) {

	// hi := Welcome{Greet: "Oh shit, waddup"}
	// json.NewEncoder(w).Encode(hi)

	fmt.Print("Oh shit, waddup")
}

func GetTime(w http.ResponseWriter, r *http.Request) {

	// t := Times{
	// 	Time{Name: "Kodee"},
	// 	Time{Name: "Kwik"},
	// }

	// json.NewEncoder(w).Encode(t)
	fmt.Print("Time for fuk")

}
