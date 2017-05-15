package router

import (
	"github.com/dimfeld/httptreemux"
	"github.com/kauehmoreno/microservico-golang/api"
	"github.com/kauehmoreno/microservico-golang/db"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func Routers() http.Handler {

	session, err := mgo.Dial("localhost:27017/MateriaGolang")

	if err != nil {
		log.Fatal(err)
	}

	repository := db.NewMateriaRepository(session)

	routerDefinition := httptreemux.NewContextMux()
	routerDefinition.Handler(http.MethodGet, "/", &api.MainHandler{})
	routerDefinition.Handler(http.MethodGet, "/materia/:id", &api.MateriaHandler{Repository: repository})
	routerDefinition.Handler(http.MethodGet, "/materias/:limit", &api.MateriasHandler{Repository: repository})

	// posts
	routerDefinition.Handler(http.MethodPost, "/api/v1/:operation/", &api.PostMateriaHandler{Repository: repository})
	//routerDefinition.Handler(http.MethodPost, "/api/v1/file_upload/", &api.FileUploadHandler{Repository: repository})

	return routerDefinition
}
