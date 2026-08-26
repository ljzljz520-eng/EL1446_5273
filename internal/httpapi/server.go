package httpapi

import (
	"encoding/json"
	"example.com/inspection14/internal/catalog"
	"example.com/inspection14/internal/domain"
	"example.com/inspection14/internal/query"
	"example.com/inspection14/internal/workflow14"
	"net/http"
)

type Server struct {
	Catalog  *catalog.Service
	Workflow *workflow14.Service
	Query    *query.Service
}

func New(c *catalog.Service, w *workflow14.Service, q *query.Service) *Server {
	return &Server{Catalog: c, Workflow: w, Query: q}
}
func (s *Server) Routes() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK); w.Write([]byte("ok")) })
	m.HandleFunc("/vehicles", s.vehicles)
	m.HandleFunc("/inspections/transition", s.transition)
	return m
}
func (s *Server) vehicles(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("q")
	v, e := s.Query.FindVehicle(term)
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(v)
}
func (s *Server) transition(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	to := r.URL.Query().Get("to")
	e := s.Workflow.Transition(id, domainStatus(to), "api", "request")
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func domainStatus(s string) domain.InspectionStatus {
	switch s {
	case "submitted":
		return domain.StatusSubmitted
	case "approved":
		return domain.StatusApproved
	case "archived":
		return domain.StatusArchived
	case "rejected":
		return domain.StatusRejected
	}
	return domain.InspectionStatus(s)
}
