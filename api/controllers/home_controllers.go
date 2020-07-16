package controllers

import (
	"net/http"

	"github.com/AlexSwiss/cyza-task/api/responses"
)

// Home function welcomes us to the api
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
