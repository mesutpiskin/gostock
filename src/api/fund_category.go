package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/allegro/bigcache"
	"go.mongodb.org/mongo-driver/bson"
	"mesutpiskin.com/gostock/helper/nosql"
)

func GetFundsCategoryValues(w http.ResponseWriter, r *http.Request, cache *bigcache.BigCache) {
	w.Header().Set("Content-Type", "application/json")
	collection := nosql.MongoClient.Database("db0").Collection("price")
	var result []bson.M
	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id":        "$category",
				"totalprice": bson.M{"$sum": "$investors"},
			},
		},
		{
			"$sort": bson.M{"totalprice": -1},
		},
	}
	showInfoCursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		panic(err)
	}
	if err = showInfoCursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(result)
}
