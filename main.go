package main

import (
	"github.com/dimfeld/httptreemux"
	"github.com/kauehmoreno/microservico-golang/router"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func main() {
	router = router.routers()
	address := "127.0.0.0:8081"
	log.Printf("Running api restfull on : http://%s\n", address)
	log.Fatal(http.ListenAndServe(address, router))

}
