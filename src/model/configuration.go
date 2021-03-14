package model

// ConfigurationModel modeel
type ConfigurationModel struct {
	Mongo        MongoDBModel `json:"mongodb"`
	ScreaperUrls ScreaperUrls `json:"screaperurls"`
}

// MongoDBModel configuration model
type MongoDBModel struct {
	HostURL  string `json:"host"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

//Screaper urls
type ScreaperUrls struct {
	Tefas string `json:"tefas"`
}
