package handler

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxcraig112/go-crud/gcp"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Handler struct {
	Ctx     context.Context
	Clients *gcp.Clients
	Mws     []func(http.Handler) http.Handler
}

func NewHandler(ctx context.Context) *Handler {
	return &Handler{
		Ctx: ctx,
		Mws: make([]func(http.Handler) http.Handler, 0),
	}
}

func (h *Handler) WithClients(clients *gcp.Clients) *Handler {
	h.Clients = clients
	return h
}

func (h *Handler) WithMiddleware(mws func(http.Handler) http.Handler) *Handler {
	h.Mws = append(h.Mws, mws)
	return h
}

// Register sets up the routes with the provided router, applying all middleware defined.
func (h *Handler) Register(r *mux.Router, routes []Route, mws ...func(http.Handler) http.Handler) {
	for _, route := range routes {
		var wrapped http.Handler
		// apply handler specific middleware
		for _, mw := range h.Mws {
			if wrapped == nil {
				wrapped = mw(route.HandlerFunc)
			} else {
				wrapped = mw(wrapped)
			}
		}

		// apply route specific middleware
		for _, mw := range mws {
			if wrapped == nil {
				wrapped = mw(route.HandlerFunc)
			} else {
				wrapped = mw(wrapped)
			}
		}

		r.Handle(route.Pattern, wrapped).Methods(route.Method)
	}
}
