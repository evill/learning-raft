package inbound

import openapi "github.com/evill/learning-raft/internal/transports/http/inbound/native"

type Router struct {
	handler *Handler
}

func NewRouter(handler *Handler) *Router {
	return &Router{handler: handler}
}

func (r *Router) Routes() openapi.Routes {
	return openapi.Routes{
		"GetCurrentNode": openapi.Route{
			Method:      "GET",
			Pattern:     "/api/internal/nodes/current",
			HandlerFunc: r.handler.GetCurrentNode,
		},
	}
}
