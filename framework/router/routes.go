package router

import (
	hello "github.com/kvmac/server/framework/handlers/hello"
	time "github.com/kvmac/server/framework/handlers/time"
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
