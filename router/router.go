package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	// create new router
	r := mux.NewRouter()

	for _, route := range AppRoutes {
		// var handler http.Handler

		// handler = route.HandlerFunc
		// handler = Logger(handler, route.Name)

		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	r.Path("/").Handler(http.FileServer(http.Dir("ng2-app/")))

	return r

}
