package model

import "time"

// Fund struct (Model)
type Fund struct {
	BasicFund   BasicFund   `json:"basicfund"`
	FromCache   bool        `json:"fromcache"`
	FundProfile FundProfile `json:"fundprofile"`
}

// FundCode struct (Model)
type FundCode struct {
	Code string `json:"code"`
}

// BasicFund model excluded fund profile
type BasicFund struct {
	Name               string    `json:"name"`
	Code               string    `json:"code"`
	Price              float64   `json:"price"`
	DailyReturn        float64   `json:"dailyreturn"`
	Pcs                float64   `json:"pcs"`
	TotalValue         float64   `json:"totalvalue"`
	Category           string    `json:"category"`
	Investors          float64   `json:"investors"`
	MarketShare        float64   `json:"marketshare"`
	Last1MounthsPrice  float64   `json:"last1mounthsprice"`
	Last3MounthsPrice  float64   `json:"last3mounthsprice"`
	Last6MounthsPrice  float64   `json:"last6mounthsprice"`
	Last12MounthsPrice float64   `json:"last12mounthsprice"`
	DateTime           time.Time `json:"datetime"`
}
