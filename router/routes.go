package router

import (
	handler "github.com/kvmac/server/handlers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var AppRoutes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
	},
	Route{
		"Hello",
		"GET",
		"/hello",
		handler.Hello,
	},
	Route{
		"GetTime",
		"GET",
		"/time",
		handler.GetTime,
	},
}
