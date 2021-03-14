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
	router.HandleFunc("/fund/tefas/{name}", func(w http.ResponseWriter, r *http.Request) {
		api.GetFund(w, r, cache)
	}).Methods("GET")

	router.HandleFunc("/fund/tefas/profile/{name}", func(w http.ResponseWriter, r *http.Request) {
		api.GetFundProfile(w, r, cache)
	}).Methods("GET")

	router.HandleFunc("/fund/tefas/multiple", func(w http.ResponseWriter, r *http.Request) {
		api.GetFunds(w, r)
	}).Methods("POST")

	router.HandleFunc("/fund/allfund", func(w http.ResponseWriter, r *http.Request) {
		api.GetAllFundsFromDb(w, r, cache)
	}).Methods("GET")

	router.HandleFunc("/fund/{code}", func(w http.ResponseWriter, r *http.Request) {
		api.GetFundFromDb(w, r, cache)
	}).Methods("GET")

	router.HandleFunc("/fund/profile/{code}", func(w http.ResponseWriter, r *http.Request) {
		api.GetFundProfileFromDb(w, r, cache)
	}).Methods("GET")

	router.HandleFunc("/fund/history/{name}", func(w http.ResponseWriter, r *http.Request) {
		api.GetAllFundsReportFromDb(w, r, cache)
	}).Methods("GET")

	// Categories
	router.HandleFunc("/fund/categories/last6mounthvalues", func(w http.ResponseWriter, r *http.Request) {
		api.GetFundsCategoryValues(w, r, cache)
	}).Methods("GET")

	//scheduled end points
	router.HandleFunc("/job/tefas/price", func(w http.ResponseWriter, r *http.Request) {
		api.ScrapeAndPersistTodayFundPriceByCode(w, r)
	}).Methods("POST")

	router.HandleFunc("/job/tefas/profile", func(w http.ResponseWriter, r *http.Request) {
		api.ScrapeAndPersistTodayFundProfileByCode(w, r)
	}).Methods("POST")

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
