package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SergioVenicio/neo4jMusic/repositories"
	"github.com/go-chi/chi"
)

type AlbumHandler struct {
	Repository *repositories.AlbumRepository
}

func (h *AlbumHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	albums, err := h.Repository.FindAll(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(albums)
}


func (h *AlbumHandler) FindByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := chi.URLParam(r, "name")
	ctx := context.Background()
	albums, err := h.Repository.FindByName(ctx, name)
	if err != nil || len(albums) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(albums)
}

func NewAlbumHandler(repository *repositories.AlbumRepository) *AlbumHandler {
	return &AlbumHandler{
		Repository: repository,
	}
}
