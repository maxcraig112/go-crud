package handler

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxcraig112/go-crud/gcp"
)

type Handler struct {
	Ctx     context.Context
	Router  *mux.Router
	Clients *gcp.Clients
	Mws     []func(http.Handler) http.Handler
}

func NewHandler(ctx context.Context, r *mux.Router) *Handler {
	return &Handler{
		Ctx:    ctx,
		Router: r,
		Mws:    make([]func(http.Handler) http.Handler, 0),
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
func (h *Handler) Register(method, pattern string, handlerFunc http.HandlerFunc) {
	var wrapped http.Handler
	// apply handler specific middleware
	for _, mw := range h.Mws {
		if wrapped == nil {
			wrapped = mw(handlerFunc)
		} else {
			wrapped = mw(wrapped)
		}
	}

	h.Router.Handle(pattern, wrapped).Methods(method)

}
