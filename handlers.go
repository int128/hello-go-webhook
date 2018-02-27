package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "OK")
}

func receiveWebhook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("content-type") != "application/json" || r.Body == nil {
		http.Error(w, "Body must be application/json", 400)
	}

	var webhookEvent WebhookEvent
	if err := json.NewDecoder(r.Body).Decode(&webhookEvent); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprintf(w, "id = %d", webhookEvent.ID)
}
