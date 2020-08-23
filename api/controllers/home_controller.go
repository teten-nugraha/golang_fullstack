package controllers

import (
	"net/http"

	"github.com/teten777/golang_fullstack/api/responses"
)

func (server *Server) Home(w http.ReponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this API")
}
