package helper

import (
	"time"

	"mesutpiskin.com/gostock/model"

	"github.com/gocolly/colly"
)

const TefasUrl = "https://www.tefas.gov.tr/FonAnaliz.aspx?FonKod="

//ScrapeFundByCode get fund price model by fund code
func ScrapeFundByCode(code string) model.Fund {
	collector := colly.NewCollector()
	var fundPrice model.Fund
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

		fundPrice = model.Fund{
			Price:              lastPrice,
			DailyReturn:        dailyReturn,
			Pcs:                pcs,
			TotalValue:         totalValue,
			Category:           category,
			Code:               code,
			Name:               name,
			Investors:          investors,
			MarketShare:        marketShare,
			Last1MounthsPrice:  last1MounthsPrice,
			Last3MounthsPrice:  last3MounthsPrice,
			Last6MounthsPrice:  last6MounthsPrice,
			Last12MounthsPrice: last12MounthsPrice,
			FromCache:          false,
			DateTime:           time.Now().Format("02-01-2006")}
	})
	collector.Visit(TefasUrl + code)

	return fundPrice
}

//ScrapeFundProfileByCode get fund prfile model by fund code
func ScrapeFundProfileByCode(fundCode string) model.FundProfile {
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
			BuyingValueDate:              buyingbaluedate,
			SaleValueDate:                salevaluedate,
			MinPurchaseTransactionAmount: minpurchasetransactionamount,
			MinSalesTransactionAmount:    minsalestransactionamount,
			MaxPurchaseTransactionAmount: maxpurchasetransactionamount,
			MaxSalesTransactionAmount:    maxsalestransactionamount,
			Status:                       status,
			EntryCommission:              entrycommission,
			ExitCommission:               exitcommission,
			KapUrl:                       kapURL,
			FromCache:                    false,
			DateTime:                     time.Now().Format("02-01-2006")}
	})
	collector.Visit(TefasUrl + fundCode)
	return fundProfile
}
