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
	collector.OnHTML("#MainContent_PanelInfo > div.main-indicators", func(content *colly.HTMLElement) {
		mainIndicators := content
		name  := mainIndicators.ChildText("#MainContent_FormViewMainIndicators_LabelFund")
		lastPrice := mainIndicators.ChildText("ul.top-list > li:nth-child(1) > span")
		dailyReturn := mainIndicators.ChildText("ul.top-list > li:nth-child(2) > span")
		pcs := mainIndicators.ChildText("ul.top-list > li:nth-child(3) > span")
		totalValue := mainIndicators.ChildText("ul.top-list > li:nth-child(4) > span")
		category := mainIndicators.ChildText("ul.top-list > li:nth-child(5) > span")
		investors := mainIndicators.ChildText("ul:nth-child(3) > li:nth-child(2) > span")
		marketShare := mainIndicators.ChildText("ul:nth-child(3) > li:nth-child(3) > span")

		fundPrice = model.Fund{
			Price: lastPrice, 
			DailyReturn: dailyReturn, 
			Pcs: pcs, 
			TotalValue: totalValue, 
			Category: category, 
			Code: code, 
			Name: name,
			Investors: investors,
			MarketShare: marketShare,
			FromCache: false,
			DateTime: time.Now()}
	})
	collector.Visit(TefasUrl + code)

	return fundPrice
}
