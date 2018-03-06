package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pschlump/jsonp"
	log "github.com/sirupsen/logrus"
)

type List struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Visibility bool      `json:"visibility"`
	Token      string    `json:"token"`
	Products   []Product `json:"products"`
	Views      int64     `json:"views"`
}

type Member struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type Product struct {
	ID          string  `json:"id"`
	URL         string  `json:"url"`
	Ref         string  `json:"ref"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"`
	HuntlyPrice float64 `json:"huntly_price"`
	Picture     string  `json:"picture"`
	ListID      string  `json:"list_id"`
	Count       int     `json:"count"`
}

//GetToken get user token
func (s *Server) GetToken(w http.ResponseWriter, r *http.Request) {
	var member Member
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		log.Error(err.Error())
		return
	}

	m, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		log.Error(err.Error())
		return
	}
	member.ID = m.Get("id")

	if err := r.Body.Close(); err != nil {
		log.Error(err.Error())
		return
	}

	u, err := RepoCreateUser(s.DB, member)
	if err != nil {
		log.Error(err.Error())
		return
	}
	res, _ := json.Marshal(u)

	io.WriteString(w, jsonp.JsonP(string(res), w, r))
}

//CreateList create a list
func (s *Server) CreateList(w http.ResponseWriter, r *http.Request) {
	var list List
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		log.Error(err.Error())
		return
	}

	m, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		log.Error(err.Error())
		return
	}
	list.Name = m.Get("name")
	list.Token = m.Get("token")
	if m.Get("visibility") == "true" {
		list.Visibility = true
	}

	if err := r.Body.Close(); err != nil {
		log.Error(err.Error())
		return
	}

	list.ID, err = RepoCreateList(s.DB, list)
	if err != nil {
		log.Error(err.Error())
	}
	res, _ := json.Marshal(list)
	io.WriteString(w, jsonp.JsonP(string(res), w, r))
}

//GetList get a list
func (s *Server) GetList(w http.ResponseWriter, r *http.Request) {
	var list List
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		log.Error(err.Error())
		return
	}

	m, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		log.Error(err.Error())
		return
	}
	list.Token = m.Get("token")
	list.ID = m.Get("id")

	if err := r.Body.Close(); err != nil {
		log.Error(err.Error())
		return
	}

	list, err = RepoGetList(s.DB, list)
	if err != nil {
		log.Error(err.Error())
		return
	}
	res, _ := json.Marshal(list)
	io.WriteString(w, jsonp.JsonP(string(res), w, r))
}

//GetLists get a list of product lists associated with a token
func (s *Server) GetLists(w http.ResponseWriter, r *http.Request) {
	var member Member
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		log.Error(err.Error())
		return
	}

	m, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		log.Error(err.Error())
		return
	}
	member.Token = m.Get("token")

	if err := r.Body.Close(); err != nil {
		log.Error(err.Error())
		return
	}

	lists, err := RepoGetLists(s.DB, member)
	if err != nil {
		log.Error(err.Error())
		return
	}
	res, _ := json.Marshal(lists)
	io.WriteString(w, jsonp.JsonP(string(res), w, r))
}

//PublishList publish a list
//TODO
// func (s *Server) PublishList(w http.ResponseWriter, r *http.Request) {
// 	// var list List
// 	// uri, err := url.ParseRequestURI(r.RequestURI)
// 	// if err != nil {
// 	// 	log.Error(err.Error())
// 	// 	return
// 	// }

// 	// m, err := url.ParseQuery(uri.RawQuery)
// 	// if err != nil {
// 	// 	log.Error(err.Error())
// 	// 	return
// 	// }
// 	// list.Token = m.Get("token")
// 	// list.ID = m.Get("id")

// 	// if err := r.Body.Close(); err != nil {
// 	// 	log.Error(err.Error())
// 	// 	return
// 	// }

// 	// list, err = RepoPublishList(s.DB, list)
// 	// if err != nil {
// 	// 	log.Error(err.Error())
// 	// 	return
// 	// }
// 	// res, _ := json.Marshal(list)
// 	// io.WriteString(w, jsonp.JsonP(string(res), w, r))
// }

//UnpublishList unpublish a list
//TODO
// func (s *Server) UnpublishList(w http.ResponseWriter, r *http.Request) {
// 	// var list List
// 	// uri, err := url.ParseRequestURI(r.RequestURI)
// 	// if err != nil {
// 	// 	log.Error(err.Error())
// 	// 	return
// 	// }

// 	// m, err := url.ParseQuery(uri.RawQuery)
// 	// if err != nil {
// 	// 	log.Error(err.Error())
// 	// 	return
// 	// }
// 	// list.Token = m.Get("token")
// 	// list.ID = m.Get("id")

// 	// if err := r.Body.Close(); err != nil {
// 	// 	log.Error(err.Error())
// 	// 	return
// 	// }

// 	// list, err = RepoUnpublishList(s.DB, list)
// 	// if err != nil {
// 	// 	log.Error(err.Error())
// 	// 	return
// 	// }
// 	// res, _ := json.Marshal(list)
// 	// io.WriteString(w, jsonp.JsonP(string(res), w, r))
// }

//DeleteList delete a list
func (s *Server) DeleteList(w http.ResponseWriter, r *http.Request) {
	var list List
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		log.Error(err.Error())
		return
	}

	m, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		log.Error(err.Error())
		return
	}
	list.Token = m.Get("token")
	list.ID = m.Get("id")

	if err := r.Body.Close(); err != nil {
		log.Error(err.Error())
		return
	}

	err = RepoDeleteList(s.DB, list)
	if err != nil {
		log.Error(err.Error())
	}
	res, _ := json.Marshal(err)
	io.WriteString(w, jsonp.JsonP(string(res), w, r))
}

//UpdateList update a list
func (s *Server) UpdateList(w http.ResponseWriter, r *http.Request) {
	var list List
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		log.Error(err.Error())
		return
	}

	m, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		log.Error(err.Error())
		return
	}
	list.Token = m.Get("token")
	list.ID = m.Get("id")
	list.Name = m.Get("name")

	if err := r.Body.Close(); err != nil {
		log.Error(err.Error())
		return
	}

	err = RepoUpdateList(s.DB, list)
	if err != nil {
		log.Error(err.Error())
	}
	res, _ := json.Marshal(err)
	io.WriteString(w, jsonp.JsonP(string(res), w, r))
}

//AddProduct add product to list
func (s *Server) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		log.Error(err.Error())
		return
	}

	m, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		log.Error(err.Error())
		return
	}
	product.ListID = m.Get("list_id")
	product.Brand = m.Get("brand")
	product.Price, _ = strconv.ParseFloat(m.Get("price"), 64)
	product.HuntlyPrice, _ = strconv.ParseFloat(m.Get("huntly_price"), 64)
	product.URL = m.Get("url")
	product.Ref = m.Get("ref")
	product.Picture = m.Get("picture")

	if err := r.Body.Close(); err != nil {
		log.Error(err.Error())
		return
	}

	product.ListID, err = RepoAddProduct(s.DB, product)
	if err != nil {
		log.Error(err.Error())
		return
	}
	res, _ := json.Marshal(product)
	io.WriteString(w, jsonp.JsonP(string(res), w, r))
}

//DeleteProduct delete product from a list
func (s *Server) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		log.Error(err.Error())
		return
	}

	m, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		log.Error(err.Error())
		return
	}
	product.ListID = m.Get("list_id")
	product.ID = m.Get("id")

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	err = RepoDeleteProduct(s.DB, product)
	if err != nil {
		log.Error(err.Error())
	}
	res, _ := json.Marshal(err)
	io.WriteString(w, jsonp.JsonP(string(res), w, r))
}

//GetProduct get product
func (s *Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		log.Error(err.Error())
		return
	}

	m, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		log.Error(err.Error())
		return
	}
	product.ID = m.Get("id")

	if err := r.Body.Close(); err != nil {
		log.Error(err.Error())
		return
	}

	product, err = RepoGetProduct(s.DB, product)
	if err != nil {
		log.Error(err.Error())
		return
	}
	res, _ := json.Marshal(product)
	io.WriteString(w, jsonp.JsonP(string(res), w, r))
}
