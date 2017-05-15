package api

import (
	"encoding/json"
	"fmt"
	"github.com/dimfeld/httptreemux"
	"github.com/kauehmoreno/microservico-golang/db"
	"github.com/kauehmoreno/microservico-golang/entities"
	"log"
	"net/http"
)

type MainHandler struct{}

func (main *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Página Principal")
}

type MateriaHandler struct {
	Repository *db.MateriaRepository
}

func (materia *MateriaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())

	article, err := materia.Repository.Get(params["id"])
	encoder := json.NewEncoder(w)
	if err == nil {
		erro := encoder.Encode(article)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
		}
	} else {
		log.Println("Failed to findById ", err)
		fmt.Println("not found")
	}
}

type MateriasHandler struct {
	Repository *db.MateriaRepository
}

func (materias *MateriasHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	articles, err := materias.Repository.GetAll()

	if err != nil {
		log.Println("Failed to fetch materia:, ", err)
		fmt.Println("Ocorreu um problema")
	}
	encode := json.NewEncoder(w)
	erro := encode.Encode(articles)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
}

type PostMateriaHandler struct {
	Repository *db.MateriaRepository
}

func (materia *PostMateriaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//pegar no contexto
	article := &entities.Materia{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(article)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	error := materia.Repository.Insert(article)

	if error == db.ErrDuplicatedMateria {
		log.Printf("%s is already Created \n", article.Titulo)
	} else if err != nil {
		log.Printf("Fail to create team:", err)
	}
	fmt.Fprintln(w, "OK")
}
