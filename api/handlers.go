package api

import (
	"fmt"
	"net/http"
)

type MainHandler struct{}

func (main *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Página Principal")
}
