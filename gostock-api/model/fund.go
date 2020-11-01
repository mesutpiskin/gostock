package model

// Fund struct (Model)
type Fund struct {
	Name 		string `json:"name"`
	Code        string `json:"code"`
	Price       string `json:"price"`
	DailyReturn string `json:"dailyreturn"`
	Pcs         string `json:"pcs"`
	TotalValue  string `json:"totalvalue"`
	Category    string `json:"category"`
	Investors   string `json:"investors"`
	MarketShare string `json:"marketshare"`
	FromCache   string `json:"fromcache"`
}
