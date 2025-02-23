package inbound

import (
	"context"
	"errors"
	"fmt"
	"github.com/evill/learning-raft/internal/core/interfaces"
	"log"
	"net/http"
	"runtime/debug"
)

const DefaultPort = 8080

type CustomHandler struct {
	service     *Service
	handlerFunc func(http.ResponseWriter, *http.Request, *Service)
}

func (ch *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovered from panic: %v\n%s", err, debug.Stack())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}()
	ch.handlerFunc(w, r, ch.service)
}

func NewCustomHandler(service *Service, handlerFunc func(http.ResponseWriter, *http.Request, *Service)) *CustomHandler {
	return &CustomHandler{
		service:     service,
		handlerFunc: handlerFunc,
	}
}

type Service struct {
	nc     interfaces.INodeOperations
	server *http.Server
}

func (service *Service) Start(port uint16) {
	if port == 0 {
		port = DefaultPort
	}

	address := fmt.Sprintf(":%d", port)

	log.Printf("Starting node HTTP API on port %s", address)
	log.Printf("HTTP API Internal URI %s", service.nc.GetNodeInfo().Address.Internal.Uri)
	// TODO: Move to constructor
	service.server = &http.Server{
		Addr:    address,
		Handler: newServerMux(service),
	}

	err := service.server.ListenAndServe()
	if err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			log.Println("HTTP server stopped")
			return
		default:
			log.Panicln("Error starting HTTP server:", err)
		}
	}
}

func (service *Service) Stop() {
	log.Println("Stopping node HTTP API...")
	if err := service.server.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("HTTP server Shutdown: %v", err)
	}
	log.Println("Node HTTP API stopped")
}

func NewApiService(nc interfaces.INodeOperations /**, pass config with port */) *Service {
	return &Service{
		nc: nc,
	}
}

func newServerMux(service *Service) *http.ServeMux {
	mux := http.NewServeMux()
	registerNodeEndpoints(mux, service)
	return mux
}
