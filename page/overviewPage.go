package page

import "BxhUiTest/internel"

type OverViewPage struct {
	internel.Page
	ChainNum    internel.Element
	TxNum       internel.Element
	BlockHeight internel.Element
}

func NewOverviewPage() OverViewPage {
	page := internel.NewPage()
	return OverViewPage{
		Page:        page,
		ChainNum:    internel.NewElement("xpath", "//*[@id=\"root\"]/section/section/main/div/div[3]/div/div/div[2]/div[1]/div/div[2]/div", page.Driver),
		TxNum:       internel.NewElement("xpath", "//*[@id=\"root\"]/section/section/main/div/div[3]/div/div/div[2]/div[2]/div/div[2]/div", page.Driver),
		BlockHeight: internel.NewElement("xpath", "//*[@id=\"root\"]/section/section/main/div/div[3]/div/div/div[2]/div[3]/div/div[2]/div", page.Driver),
	}
}
