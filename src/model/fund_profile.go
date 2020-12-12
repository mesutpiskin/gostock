package model

//FundProfile model
type FundProfile struct {
	Code                         string `json:"code"`
	Isin                         string `json:"isin"`
	StartingTime                 string `json:"startingtime"`
	EndTime                      string `json:"endtime"`
	BuyingValueDate              string `json:"buyingbaluedate"`
	SaleValueDate                string `json:"salevaluedate"`
	MinPurchaseTransactionAmount string `json:"minpurchasetransactionamount"`
	MinSalesTransactionAmount    string `json:"minsalestransactionamount"`
	MaxPurchaseTransactionAmount string `json:"maxpurchasetransactionamount"`
	MaxSalesTransactionAmount    string `json:"maxsalestransactionamount"`
	Status                       string `json:"status"`
	EntryCommission              string `json:"entrycommission"`
	ExitCommission               string `json:"exitcommission"`
	KapUrl                       string `json:"kapurl"`
	FromCache                    bool   `json:"fromcache"`
	DateTime                     string `json:"datetime"`
}
