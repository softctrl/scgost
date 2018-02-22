//
// Author:
//  Carlos Timoshenko
//  carlostimoshenkorodrigueslopes@gmail.com
//
//  https://github.com/softctrl
//
// This project is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation; either
// version 3 of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.
//
package scgost

import (
	f "fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/softctrl/scgotils/schttp"
	l "github.com/softctrl/scgotils/sclogger"
)

const (
	_TAG_    = "route.go"
	_STR_FMT = "%s"
	_API_FMT = "/api/%s/%s"
	_LOG_FMT = "\tAdding Route [%s]-[%s] - %s"
)

//
// Route object.
//
type Route struct {
	Name        string
	Version     string
	Methods     schttp.Methods
	Pattern     string
	HandlerFunc http.HandlerFunc
	Root        bool
}

//
// Route Factory.
//
type RouteFactory interface {
	Get() *mux.Router
}

//
// Route slices.
//
type Routes []Route

//
// Format the API versioned URL, only if it is not a root route.
//
func (__obj *Route) ApiVersionUrl() string {

	if __obj.Root {
		return f.Sprintf(_STR_FMT, __obj.Pattern)
	} else {
		return f.Sprintf(_API_FMT, __obj.Version, __obj.Pattern)
	}

}

//
// Configure a routhe that will identified by a fixed path
//
func _ConfigureRouteMethod(__router *mux.Router, __route Route) *mux.Router {

	var handler http.Handler

	_url := __route.ApiVersionUrl()

	handler = __route.HandlerFunc
	handler = l.HttpHandlerLogger(handler, __route.Name)

	for _, __Method := range __route.Methods {

		l.I(_TAG_, _LOG_FMT, __Method, _url, __route.Name)

		__router.Methods(string(__Method)).
			Path(_url).
			Name(__route.Name).
			Handler(handler)

	}

	return __router

}

//
// Configure a routhe that will be handled by url prefix
// TODO Under development
//
func _ConfigureRoutePrefix(__router *mux.Router, __route Route) *mux.Route {

	return nil

}

//
// Add all routes into <__routes>.
//
func LoadRoutes(__router *mux.Router, __routes Routes) {

	for _, _route := range __routes {

		_ConfigureRouteMethod(__router, _route)

	}

}
