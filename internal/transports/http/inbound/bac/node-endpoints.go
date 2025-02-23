package inbound

import (
	"encoding/json"
	"log"
	"net/http"
)

func registerNodeEndpoints(mux *http.ServeMux, service *Service) {
	mux.Handle("/internal/nodes/current", NewCustomHandler(service, getNodeInfo))
}

func getNodeInfo(w http.ResponseWriter, r *http.Request, service *Service) {
	if r.Method != http.MethodGet {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	nodeInfo := (*service).nc.GetNodeInfo()
	jsonData, err := json.Marshal(nodeInfo)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, wErr := w.Write(jsonData)

	if wErr != nil {
		log.Println(wErr)
	}
}
