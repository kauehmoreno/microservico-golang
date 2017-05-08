package main

import (
	"github.com/dimfeld/httptreemux"
	"github.com/kauehmoreno/microservico-golang/api"
	"gopkg.in/mgo.v2"
	"net/http"
)

func routers() {

	session, err := mgo.Dial("localhost:27017/Soccer2017")

	if err != nil {
		log.Fatal(err)
	}

	repository := db.NewSoccerRepository(session)

	router := httptreemux.NewContextMux()
	router.Handler(http.MethodGet, "/", &api.MainHandler)
	router.Handler(http.MethodGet, "/materia/:id", &api.MateriaHandler{})
	router.Handler(http.MethodGet, "/materias/:order_by/:limit", &api.MateriasHandler{})

	// posts
	router.Handler(http.MethodPost, "/api/v1/:operation/", &api.PostMateriaHandler{Repository: repository})
	router.Handler(http.MethodPost, "/api/v1/file_upload/", &api.FileUploadHandler{Repository: repository})

	return router
}
