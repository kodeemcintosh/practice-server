package main

import (
	// "fmt"
	// "time"
	// "encoding/json"
	// "github.com/gorilla/mux"
	router "github.com/kvmac/server/framework/router"
	"net/http"
)

func main() {
	r := router.Router()
	// r.HandleFunc("/", index.Index)

	// r := mux.NewRouter()
	// r.Handle("/", http.FileServer(http.Dir("./ng2-app/")))
	// // r.PathPrefix("/").Handler(http.FileServer(http.Dir("./ng2-app/")))
	// r.HandleFunc("/", Index)
	// // r.Handle("/", Index)
	// // api := r.PathPrefix("/api/v1").Subrouter()

	// // api.HandleFunc("/hello", Hello)
	// // api.HandleFunc("/time", GetTime)
	// r.HandleFunc("/hello", Hello)
	// r.HandleFunc("/time", GetTime)

	http.ListenAndServe(":8080", r)

}
