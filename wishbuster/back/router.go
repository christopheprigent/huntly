package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/token", s.GetToken).Methods("GET")         //ok
	s.Router.HandleFunc("/list/create", s.CreateList).Methods("GET") //ok
	s.Router.HandleFunc("/list", s.GetList).Methods("GET")           //ok
	s.Router.HandleFunc("/lists", s.GetLists).Methods("GET")         //ok
	// s.Router.HandleFunc("/list/publish", s.PublishList).Methods("GET")
	// s.Router.HandleFunc("/list/unpublish", s.UnpublishList).Methods("GET")
	s.Router.HandleFunc("/list/delete", s.DeleteList).Methods("GET") //ok
	s.Router.HandleFunc("/list/product/add", s.AddProduct).Methods("GET")
	s.Router.HandleFunc("/list/product/delete", s.DeleteProduct).Methods("GET")
	s.Router.HandleFunc("/product", s.GetProduct).Methods("GET")
}
