package salary

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func runServer() {
	router := mux.NewRouter()

	router.HandleFunc("/knapsack", createKnapsack)
	router.HandleFunc("/knapsack/{id}", fetchKnapsack)
	router.HandleFunc("/shutdown", shutdown)

	log.Printf("Running http server")
	err := http.ListenAndServe(":6543", router)
	if err != nil {
		log.Printf("http server failed with %v", err)
	}
	log.Printf("http server closed down")
}

