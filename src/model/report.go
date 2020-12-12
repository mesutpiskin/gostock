package model

import "go.mongodb.org/mongo-driver/bson"

// Report struct (Model)
type Report struct {
	HistoricalData []bson.M `json:"historicaldata"`
	Code           string   `json:"code"`
	FromCache      bool     `json:"fromcache"`
}
