package main

import (
	"log"
	"net/http"
	"time"

	"github.com/allegro/bigcache"
	"github.com/gorilla/mux"
	"mesutpiskin.com/gostock/api"
)

const port string = "5000"

// Main function
func main() {
	// Init router
	router := mux.NewRouter()
	// Init memory cache
	cache, initErr := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if initErr != nil {
		log.Fatal(initErr)
	}
	// Route handles & endpoints

	router.HandleFunc("/funds/{name}", func(w http.ResponseWriter, r *http.Request) {
		api.GetFund(w, r, cache)
	}).Methods("GET")
	router.HandleFunc("/health", api.Healthy).Methods("GET")

	// Start server
	print("Server starting... \n http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
