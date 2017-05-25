package router

import (
	hello "github.com/kvmac/server/handlers/hello"
	index "github.com/kvmac/server/handlers/index"
	time "github.com/kvmac/server/handlers/time"
	"net/http"
)

//Route is a structure for the api router call
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes because Route objects have to go somewhere
type Routes []Route

var appRoutes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		index.Index,
	},
	Route{
		"Hello",
		"GET",
		"/hello",
		hello.Hello,
	},
	Route{
		"GetTime",
		"GET",
		"/time",
		time.GetTime,
	},
}