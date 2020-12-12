package model

// Fund struct (Model)
type Fund struct {
	Name               string      `json:"name"`
	Code               string      `json:"code"`
	Price              string      `json:"price"`
	DailyReturn        string      `json:"dailyreturn"`
	Pcs                string      `json:"pcs"`
	TotalValue         string      `json:"totalvalue"`
	Category           string      `json:"category"`
	Investors          string      `json:"investors"`
	MarketShare        string      `json:"marketshare"`
	Last1MounthsPrice  string      `json:"last1mounthsprice"`
	Last3MounthsPrice  string      `json:"last3mounthsprice"`
	Last6MounthsPrice  string      `json:"last6mounthsprice"`
	Last12MounthsPrice string      `json:"last12mounthsprice"`
	FromCache          bool        `json:"fromcache"`
	DateTime           string      `json:"datetime"`
	FundProfile        FundProfile `json:"fundprofile"`
}

// FundCode struct (Model)
type FundCode struct {
	Code string `json:"code"`
}
