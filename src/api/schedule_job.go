package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"mesutpiskin.com/gostock/helper"
	"mesutpiskin.com/gostock/helper/nosql"
	"mesutpiskin.com/gostock/model"
)

//ScrapeAndPersistTodayFundPriceByCode getting all fund data and save to mongo
func ScrapeAndPersistTodayFundPriceByCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var fundcodes []model.FundCode

	err := json.NewDecoder(r.Body).Decode(&fundcodes)
	if err != nil {
		panic(err)
	}
	var configurationModel model.ConfigurationModel
	configurationModel = helper.GetConfiguration()
	for _, element := range fundcodes {
		basicFund := helper.ScrapeFundByCode(strings.ToUpper(element.Code), configurationModel.ScreaperUrls.Tefas)
		saveFundPriceToDB(basicFund)
	}

	json.NewEncoder(w).Encode("Task completed.")
}

//ScrapeAndPersistTodayFundProfileByCode getting fund profile
func ScrapeAndPersistTodayFundProfileByCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var fundcodes []model.FundCode

	err := json.NewDecoder(r.Body).Decode(&fundcodes)
	if err != nil {
		panic(err)
	}
	var configurationModel model.ConfigurationModel
	configurationModel = helper.GetConfiguration()
	for _, element := range fundcodes {
		fundProfiledata := helper.ScrapeFundProfileByCode(strings.ToUpper(element.Code), configurationModel.ScreaperUrls.Tefas)
		go saveFundProfileToDB(fundProfiledata)
	}

	json.NewEncoder(w).Encode("Task completed.")
}

func saveFundPriceToDB(fundPriceModel model.BasicFund) {
	collection := nosql.MongoClient.Database("db0").Collection("price")
	insertResult, err := collection.InsertOne(context.TODO(), fundPriceModel)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)
}

func saveFundProfileToDB(fundProfileModel model.FundProfile) {
	collection := nosql.MongoClient.Database("db0").Collection("profile")
	insertResult, err := collection.InsertOne(context.TODO(), fundProfileModel)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)
}
