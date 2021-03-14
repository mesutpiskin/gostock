package helper

import (
	"strconv"
	"strings"
	"time"

	"mesutpiskin.com/gostock/model"

	"github.com/gocolly/colly"
)

//ScrapeFundByCode get fund price model by fund code
func ScrapeFundByCode(code string, url string) model.BasicFund {
	collector := colly.NewCollector()
	var fundPrice model.BasicFund
	collector.OnHTML("#MainContent_PanelInfo", func(content *colly.HTMLElement) {
		mainIndicators := content
		name := mainIndicators.ChildText("#MainContent_FormViewMainIndicators_LabelFund")
		lastPrice := mainIndicators.ChildText("div.main-indicators>ul.top-list > li:nth-child(1) > span")
		dailyReturn := mainIndicators.ChildText("div.main-indicators>ul.top-list > li:nth-child(2) > span")
		pcs := mainIndicators.ChildText("div.main-indicators>ul.top-list > li:nth-child(3) > span")
		totalValue := mainIndicators.ChildText("div.main-indicators>ul.top-list > li:nth-child(4) > span")
		category := mainIndicators.ChildText("div.main-indicators>ul.top-list > li:nth-child(5) > span")
		investors := mainIndicators.ChildText("div.main-indicators>ul:nth-child(3) > li:nth-child(2) > span")
		marketShare := mainIndicators.ChildText("div.main-indicators>ul:nth-child(3) > li:nth-child(3) > span")
		last1MounthsPrice := mainIndicators.ChildText("div.price-indicators>ul > li:nth-child(1) > span")
		last3MounthsPrice := mainIndicators.ChildText("div.price-indicators>ul > li:nth-child(2) > span")
		last6MounthsPrice := mainIndicators.ChildText("div.price-indicators>ul > li:nth-child(3) > span")
		last12MounthsPrice := mainIndicators.ChildText("div.price-indicators>ul > li:nth-child(4) > span")

		fundPrice = model.BasicFund{
			Price:              textToFloat(lastPrice),
			DailyReturn:        textToFloat(dailyReturn),
			Pcs:                textToFloat(pcs),
			TotalValue:         textToFloat(totalValue),
			Category:           category,
			Code:               code,
			Name:               name,
			Investors:          textToFloat(investors),
			MarketShare:        textToFloat(marketShare),
			Last1MounthsPrice:  textToFloat(last1MounthsPrice),
			Last3MounthsPrice:  textToFloat(last3MounthsPrice),
			Last6MounthsPrice:  textToFloat(last6MounthsPrice),
			Last12MounthsPrice: textToFloat(last12MounthsPrice),
			DateTime:           time.Now()}
	})
	collector.Visit(url + code)

	return fundPrice
}

//ScrapeFundProfileByCode get fund prfile model by fund code
func ScrapeFundProfileByCode(fundCode string, url string) model.FundProfile {
	collector := colly.NewCollector()
	var fundProfile model.FundProfile
	collector.OnHTML("#MainContent_DetailsViewFund", func(content *colly.HTMLElement) {
		mainIndicators := content
		code := mainIndicators.ChildText("tr:nth-child(1) > td.fund-profile-item")
		isin := mainIndicators.ChildText("tr:nth-child(2) > td.fund-profile-item")
		startingtime := mainIndicators.ChildText("tr:nth-child(3) > td.fund-profile-item")
		endtime := mainIndicators.ChildText("tr:nth-child(4) > td.fund-profile-item")
		buyingbaluedate := mainIndicators.ChildText("tr:nth-child(5) > td.fund-profile-item")
		salevaluedate := mainIndicators.ChildText("tr:nth-child(6) > td.fund-profile-item")
		minpurchasetransactionamount := mainIndicators.ChildText("tr:nth-child(7) > td.fund-profile-item")
		minsalestransactionamount := mainIndicators.ChildText("tr:nth-child(8) > td.fund-profile-item")
		maxpurchasetransactionamount := mainIndicators.ChildText("tr:nth-child(9) > td.fund-profile-item")
		maxsalestransactionamount := mainIndicators.ChildText("tr:nth-child(10) > td.fund-profile-item")
		status := mainIndicators.ChildText("tr:nth-child(11) > td.fund-profile-item")
		entrycommission := mainIndicators.ChildText("tr:nth-child(12) > td.fund-profile-item")
		exitcommission := mainIndicators.ChildText("tr:nth-child(13) > td.fund-profile-item")
		kapURL := mainIndicators.ChildText("tr:nth-child(14) > td.fund-profile-header > a")

		fundProfile = model.FundProfile{
			Code:                         code,
			Isin:                         isin,
			StartingTime:                 startingtime,
			EndTime:                      endtime,
			BuyingValueDate:              textToFloat(buyingbaluedate),
			SaleValueDate:                textToFloat(salevaluedate),
			MinPurchaseTransactionAmount: textToFloat(minpurchasetransactionamount),
			MinSalesTransactionAmount:    textToFloat(minsalestransactionamount),
			MaxPurchaseTransactionAmount: textToFloat(maxpurchasetransactionamount),
			MaxSalesTransactionAmount:    textToFloat(maxsalestransactionamount),
			Status:                       status,
			EntryCommission:              textToFloat(entrycommission),
			ExitCommission:               textToFloat(exitcommission),
			KapUrl:                       kapURL,
			FromCache:                    false,
			DateTime:                     time.Now()}
	})
	collector.Visit(url + fundCode)
	return fundProfile
}

func textToFloat(text string) float64 {
	if text == "" {
		return 0
	}
	number := strings.Replace(text, "%", "", -1)
	number = strings.Replace(number, ".", "", -1)
	number = strings.Replace(number, ",", ".", -1)
	if number == "" {
		return 0
	}
	price, err := strconv.ParseFloat(number, 64)

	if err != nil {
		panic(err.Error())
	}

	return price
}
