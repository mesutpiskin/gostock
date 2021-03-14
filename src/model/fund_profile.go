package model

import "time"

//FundProfile model
type FundProfile struct {
	Code                         string    `json:"code"`
	Isin                         string    `json:"isin"`
	StartingTime                 string    `json:"startingtime"`
	EndTime                      string    `json:"endtime"`
	BuyingValueDate              float64   `json:"buyingbaluedate"`
	SaleValueDate                float64   `json:"salevaluedate"`
	MinPurchaseTransactionAmount float64   `json:"minpurchasetransactionamount"`
	MinSalesTransactionAmount    float64   `json:"minsalestransactionamount"`
	MaxPurchaseTransactionAmount float64   `json:"maxpurchasetransactionamount"`
	MaxSalesTransactionAmount    float64   `json:"maxsalestransactionamount"`
	Status                       string    `json:"status"`
	EntryCommission              float64   `json:"entrycommission"`
	ExitCommission               float64   `json:"exitcommission"`
	KapUrl                       string    `json:"kapurl"`
	FromCache                    bool      `json:"fromcache"`
	DateTime                     time.Time `json:"datetime"`
}
