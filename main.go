package main

import (
	"github.com/kauehmoreno/microservico-golang/router"
	"log"
	"net/http"
)

func main() {
	rt := router.Routers()
	address := "127.0.0.1:8081"
	log.Printf("Running api restfull on : http://%s\n", address)
	log.Fatal(http.ListenAndServe(address, rt))

}
