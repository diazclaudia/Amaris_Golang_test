package api

import (
	"encoding/json"
	"net/http"

	read "backend/domain/read"
	update "backend/domain/write"
	"github.com/go-chi/chi"
)

type handler struct {
	pointServiceRead   read.Service
	pointServiceUpdate update.Service
}

func NewHandler(pointServiceRead read.Service, pointServiceUpdate update.Service) PointsHandler {
	return &handler{
		pointServiceRead:   pointServiceRead,
		pointServiceUpdate: pointServiceUpdate,
	}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	p, err := h.pointServiceRead.Find(id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&p)
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	points := chi.URLParam(r, "points")
	id := chi.URLParam(r, "id")

	p, err := h.pointServiceUpdate.Update(id, points)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&p)
}
