package model

// ConfigurationModel modeel
type ConfigurationModel struct {
	Mongo MongoDBModel `json:"mongodb"`
}

// MongoDBModel configuration model
type MongoDBModel struct {
	HostURL  string `json:"host"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
