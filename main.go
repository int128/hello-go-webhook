package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "OK")
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	log.Fatal(http.ListenAndServe(":3000", router))
}
