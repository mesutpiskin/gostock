package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/allegro/bigcache"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"mesutpiskin.com/gostock/api"
	"mesutpiskin.com/gostock/helper/nosql"
)

const port string = "5000"

// Main function
func main() {

	client := nosql.ConnectToMongoDB()
	client.Ping(context.TODO(), nil)

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
	router.HandleFunc("/fund", func(w http.ResponseWriter, r *http.Request) {
		api.GetFunds(w, r)
	}).Methods("POST")
	router.HandleFunc("/health", api.Healthy).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})
	handler := c.Handler(router)

	// Start server
	print("Server starting... \n http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}
