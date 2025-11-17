package handler

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxcraig112/go-crud/gcp"
)

type Route struct {
	method      string
	pattern     string
	handlerFunc http.HandlerFunc
}

type Handler struct {
	Ctx     context.Context
	Clients *gcp.Clients
	Mws     []func(http.Handler) http.Handler
}

// Register sets up the routes with the provided router, applying all middleware defined.
func (h *Handler) Register(r *mux.Router, routes []Route, mws ...func(http.Handler) http.Handler) {
	for _, route := range routes {
		var wrapped http.Handler
		// apply handler specific middleware
		for _, mw := range h.Mws {
			if wrapped == nil {
				wrapped = mw(route.handlerFunc)
			} else {
				wrapped = mw(wrapped)
			}
		}

		// apply route specific middleware
		for _, mw := range mws {
			if wrapped == nil {
				wrapped = mw(route.handlerFunc)
			} else {
				wrapped = mw(wrapped)
			}
		}

		r.Handle(route.pattern, wrapped).Methods(route.method)
	}
}
