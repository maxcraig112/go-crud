package handler

import "net/http"

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRoute(method, pattern string, handlerFunc http.HandlerFunc) Route {
	return Route{
		Method:      method,
		Pattern:     pattern,
		HandlerFunc: handlerFunc,
	}
}
