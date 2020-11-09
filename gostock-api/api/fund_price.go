package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"mesutpiskin.com/gostock/model"

	"github.com/allegro/bigcache"
	"github.com/gorilla/mux"
	"mesutpiskin.com/gostock/helper"
	"mesutpiskin.com/gostock/helper/nosql"
)

// Init contacts var as a slice Fund struct
var funds []model.Fund

//GetFund Get single contact
func GetFund(w http.ResponseWriter, r *http.Request, cache *bigcache.BigCache) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	fundCode := strings.ToUpper(params["name"])

	//Getting fund data from cache by fund name
	cacheData, getErr := cache.Get(fundCode)
	if getErr == nil {
		var jsonFund model.Fund
		if err := helper.DecodeFromBase64(&jsonFund, string(cacheData)); err != nil {
			panic(err)
		}
		jsonFund.FromCache = true
		json.NewEncoder(w).Encode(jsonFund)
		return
	}

	//Getting from website
	fundData := helper.ScrapeFundByCode(fundCode)
	fundProfiledata := helper.ScrapeFundProfileByCode(fundCode)
	fundData.FundProfile = fundProfiledata

	enc, err := helper.EncodeToBase64(fundData)
	if err != nil {
		panic(err)
	}
	setErr := cache.Set(fundCode, []byte(enc))
	if setErr != nil {
		panic(setErr)
	}

	json.NewEncoder(w).Encode(fundData)
}

//GetFundProfile get fund profile data
func GetFundProfile(w http.ResponseWriter, r *http.Request, cache *bigcache.BigCache) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	fundCode := strings.ToUpper(params["name"])
	var key string
	key = "fund:profile-" + fundCode
	//Getting fund data from cache by fund name
	cacheData, getErr := cache.Get(key)
	if getErr == nil {
		var jsonFund model.FundProfile
		if err := helper.DecodeFromBase64(&jsonFund, string(cacheData)); err != nil {
			panic(err)
		}
		jsonFund.FromCache = true
		json.NewEncoder(w).Encode(jsonFund)
		return
	}

	//Getting from website
	fundData := helper.ScrapeFundProfileByCode(fundCode)

	enc, err := helper.EncodeToBase64(fundData)
	if err != nil {
		panic(err)
	}
	setErr := cache.Set(key, []byte(enc))
	if setErr != nil {
		panic(setErr)
	}

	json.NewEncoder(w).Encode(fundData)
}

// GetFunds get multiple fund
func GetFunds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var fundcodes []model.FundCode

	err := json.NewDecoder(r.Body).Decode(&fundcodes)
	if err != nil {
		panic(err)
	}
	var fundModels []model.Fund

	for _, element := range fundcodes {
		fundData := helper.ScrapeFundByCode(strings.ToUpper(element.Code))
		fundProfiledata := helper.ScrapeFundProfileByCode(strings.ToUpper(element.Code))
		fundData.FundProfile = fundProfiledata
		fundModels = append(fundModels, fundData)

		go saveFundToDB(fundData)
	}

	json.NewEncoder(w).Encode(fundModels)
}

func saveFundToDB(fundModel model.Fund) {
	collection := nosql.MongoClient.Database("db0").Collection("funds")
	insertResult, err := collection.InsertOne(context.TODO(), fundModel)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)
}
