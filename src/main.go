package main

import (
	"fmt"
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
	go nosql.ConnectToMongoDB()
	// Init router
	router := mux.NewRouter()
	// Init memory cache
	cache, initErr := bigcache.NewBigCache(bigcache.DefaultConfig(60 * time.Minute))
	if initErr != nil {
		log.Fatal(initErr)
	}
	// Route handles & endpoints

	router.HandleFunc("/fund/{name}", func(w http.ResponseWriter, r *http.Request) {
		api.GetFund(w, r, cache)
	}).Methods("GET")
	router.HandleFunc("/profile/{name}", func(w http.ResponseWriter, r *http.Request) {
		api.GetFundProfile(w, r, cache)
	}).Methods("GET")
	router.HandleFunc("/funds", func(w http.ResponseWriter, r *http.Request) {
		api.GetFunds(w, r)
	}).Methods("POST")
	router.HandleFunc("/allfund", func(w http.ResponseWriter, r *http.Request) {
		api.GetAllFundsFromDb(w, r, cache)
	}).Methods("GET")
	router.HandleFunc("/allfundreport/{name}", func(w http.ResponseWriter, r *http.Request) {
		api.GetAllFundsReportFromDb(w, r, cache)
	}).Methods("GET")
	router.HandleFunc("/health", api.Healthy).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})
	handler := c.Handler(router)

	// Start server
	fmt.Println("Server starting... \n http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}
