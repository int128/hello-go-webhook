package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.POST("/", receiveWebhook)
	log.Fatal(http.ListenAndServe(":3000", router))
}
