package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"mesutpiskin.com/gostock/model"

	"github.com/allegro/bigcache"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
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
	var configurationModel model.ConfigurationModel
	configurationModel = helper.GetConfiguration()
	fundPrice := helper.ScrapeFundByCode(fundCode, configurationModel.ScreaperUrls.Tefas)
	fundProfiledata := helper.ScrapeFundProfileByCode(fundCode, configurationModel.ScreaperUrls.Tefas)
	fundResult := model.Fund{BasicFund: fundPrice, FundProfile: fundProfiledata}

	enc, err := helper.EncodeToBase64(fundResult)
	if err != nil {
		panic(err)
	}
	setErr := cache.Set(fundCode, []byte(enc))
	if setErr != nil {
		panic(setErr)
	}

	json.NewEncoder(w).Encode(fundResult)
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
	var configurationModel model.ConfigurationModel
	configurationModel = helper.GetConfiguration()
	//Getting from website
	fundData := helper.ScrapeFundProfileByCode(fundCode, configurationModel.ScreaperUrls.Tefas)

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
	var configurationModel model.ConfigurationModel
	configurationModel = helper.GetConfiguration()
	for _, element := range fundcodes {
		fundPrice := helper.ScrapeFundByCode(strings.ToUpper(element.Code), configurationModel.ScreaperUrls.Tefas)
		fundProfiledata := helper.ScrapeFundProfileByCode(strings.ToUpper(element.Code), configurationModel.ScreaperUrls.Tefas)
		fundResult := model.Fund{BasicFund: fundPrice, FundProfile: fundProfiledata}
		fundModels = append(fundModels, fundResult)
	}

	json.NewEncoder(w).Encode(fundModels)
}

// GetAllFundsFromDb from database
func GetAllFundsFromDb(w http.ResponseWriter, r *http.Request, cache *bigcache.BigCache) {
	w.Header().Set("Content-Type", "application/json")
	var funds []bson.M
	var key string = "ALL_FUNDS"
	cacheData, cacheErr := cache.Get(key)
	if cacheErr == nil {
		if err := helper.DecodeFromBase64(&funds, string(cacheData)); err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(funds)
		return
	}
	collection := nosql.MongoClient.Database("db0").Collection("price")
	cursor, err := collection.Find(context.TODO(), bson.M{"datetime": time.Now().Format("02-01-2006")})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &funds); err != nil {
		log.Fatal(err)
	}

	enc, err := helper.EncodeToBase64(funds)
	if err != nil {
		panic(err)
	}
	setErr := cache.Set(key, []byte(enc))
	if setErr != nil {
		panic(setErr)
	}

	json.NewEncoder(w).Encode(funds)
}

// GetAllFundsReportFromDb from database, historical data
func GetAllFundsReportFromDb(w http.ResponseWriter, r *http.Request, cache *bigcache.BigCache) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	code := strings.ToUpper(params["name"])
	fundcodes := strings.SplitAfter(code, ";")
	var result []model.Report

	for _, fundCode := range fundcodes {
		var funds []bson.M
		var bsonFund model.Report
		var key string = "FUND_REPORT_" + fundCode

		cacheData, cacheErr := cache.Get(key)
		if cacheErr == nil {
			if err := helper.DecodeFromBase64(&funds, string(cacheData)); err != nil {
				panic(err)
			}
			bsonFund = model.Report{HistoricalData: funds, Code: fundCode, FromCache: true}
			result = append(result, bsonFund)
			continue
		}
		// Get data from db
		collection := nosql.MongoClient.Database("db0").Collection("funds")
		/*	orQuery := []bson.M{}
			uidFindQuery := bson.M{"code": "AFT"}
			nameFindQuery := bson.M{"code": "NNF"}
			orQuery = append(orQuery, uidFindQuery, nameFindQuery)
			cursor, err := collection.Find(context.TODO(), bson.M{"$or": orQuery})*/

		cursor, err := collection.Find(context.TODO(), bson.M{"code": fundCode})
		if err != nil {
			log.Fatal(err)
		}
		// Get all data
		if err = cursor.All(context.TODO(), &funds); err != nil {
			log.Fatal(err)
		}
		cursor.Close(context.TODO())
		// Encode for caching
		enc, err := helper.EncodeToBase64(funds)
		if err != nil {
			panic(err)
		}
		setErr := cache.Set(key, []byte(enc))
		if setErr != nil {
			panic(setErr)
		}
		bsonFund = model.Report{HistoricalData: funds, Code: fundCode, FromCache: false}
		result = append(result, bsonFund)
	}

	json.NewEncoder(w).Encode(result)
}

//GetFundProfileFromDb profile getter from mongo
func GetFundProfileFromDb(w http.ResponseWriter, r *http.Request, cache *bigcache.BigCache) {
	params := mux.Vars(r) // Gets params
	fundCode := strings.ToUpper(params["code"])
	var fundProfile bson.M
	var key string = "FUND_PROFILE_" + fundCode

	cacheData, cacheErr := cache.Get(key)
	if cacheErr == nil {
		if err := helper.DecodeFromBase64(&fundProfile, string(cacheData)); err != nil {
			panic(err)
		}
	}
	// Get data from db
	collection := nosql.MongoClient.Database("db0").Collection("profile")
	if err := collection.FindOne(context.TODO(), bson.M{"code": fundCode}).Decode(&fundProfile); err != nil {
		panic(err)
	}

	// Encode for caching
	enc, err := helper.EncodeToBase64(fundProfile)
	if err != nil {
		panic(err)
	}
	setErr := cache.Set(key, []byte(enc))
	if setErr != nil {
		panic(setErr)
	}

	json.NewEncoder(w).Encode(fundProfile)
}

//GetFundFromDb get fund price from mongo
func GetFundFromDb(w http.ResponseWriter, r *http.Request, cache *bigcache.BigCache) {
	params := mux.Vars(r) // Gets params
	fundCode := strings.ToUpper(params["code"])
	var fundProfile bson.M
	var key string = "FUND_PROFILE_" + fundCode

	cacheData, cacheErr := cache.Get(key)
	if cacheErr == nil {
		if err := helper.DecodeFromBase64(&fundProfile, string(cacheData)); err != nil {
			panic(err)
		}
	}
	// Get data from db
	collection := nosql.MongoClient.Database("db0").Collection("price")
	if err := collection.FindOne(context.TODO(), bson.M{"code": fundCode, "datetime": time.Now().Format("02-01-2006")}).Decode(&fundProfile); err != nil {
		panic(err)
	}

	// Encode for caching
	enc, err := helper.EncodeToBase64(fundProfile)
	if err != nil {
		panic(err)
	}
	setErr := cache.Set(key, []byte(enc))
	if setErr != nil {
		panic(setErr)
	}

	json.NewEncoder(w).Encode(fundProfile)
}
