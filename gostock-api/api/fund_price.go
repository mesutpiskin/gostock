package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"mesutpiskin.com/gostock/model"

	"github.com/allegro/bigcache"
	"github.com/gorilla/mux"
	"mesutpiskin.com/gostock/helper"
)

// Init contacts var as a slice Fund struct
var funds []model.Fund

// GetFund Get single contact
func GetFund (w http.ResponseWriter, r *http.Request, cache *bigcache.BigCache) {
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

// GetFunds get multiple fund
func GetFunds (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var fundcodes []model.FundCode
	err := json.NewDecoder(r.Body).Decode(&fundcodes)
	if err != nil {
		panic(err)
	}
	var fundModels []model.Fund
	for _, element := range fundcodes{
        data := helper.ScrapeFundByCode(strings.ToUpper(element.Code))  	    
		fundModels = append(fundModels, data)
	}   
	
	json.NewEncoder(w).Encode(fundModels)
}

