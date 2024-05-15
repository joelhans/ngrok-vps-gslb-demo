// A (very) simple API for demonstrating ngrok's global server load balancing and API
// gateway functionality.

package main

import (
	"context"
	"os"
	"log"
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type API struct {
	ID			string `json:"id"`
	DC 	 		string `json:"dc"`
}

var api []API

func getAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(api)
}

func GenerateUUID() string {
    id := uuid.New()
    return id.String()
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	// Create mock data about the API deployment.
	uniqueID := GenerateUUID()
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	api = append(api, API{ID: uniqueID, DC: name})

	router:=mux.NewRouter()
	router.HandleFunc("/api", getAPI).Methods("GET")

	return http.ListenAndServe(":5000", router)
}
