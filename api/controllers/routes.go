package controllers

import "github.com/teten777/golang_fullstack/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home()).Methods("GET"))
}
