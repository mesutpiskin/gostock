package helper

import (
	"encoding/json"
	"os"

	"mesutpiskin.com/gostock/model"
)

//GetConfiguration from file
func GetConfiguration() model.ConfigurationModel {
	file, err := os.Open("./config/config.development.json")
	if err != nil {
		panic(err)
	}

	//Decoding configuration
	var configuration model.ConfigurationModel
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}

	return configuration
}
