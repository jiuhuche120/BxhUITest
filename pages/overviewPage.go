package pages

import "BxhUiTest/internal"

type OverViewPage struct {
	internal.Page
	Confirm         internal.Element
	Login           internal.Element
	AppchainCount   internal.Element
	InterchainCount internal.Element
	BlockHeight     internal.Element
	CurrentDayTx    internal.Element
	ServiceCount    internal.Element
	LastedBlocks    internal.Elements
	LastedTxs       internal.Elements
}

func NewOverviewPage() OverViewPage {
	page := internal.NewPage()
	return OverViewPage{
		Page:            page,
		Confirm:         internal.NewElement("xpath", "/html/body/div[2]/div/div[2]/div/div[2]/div/div/div[2]/button", page.Driver),
		Login:           internal.NewElement("xpath", "//*[@id=\"root\"]/section/header/div[2]/div[2]/div", page.Driver),
		AppchainCount:   internal.NewElement("xpath", "//*[@id=\"root\"]/section/section/main/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div[2]/div", page.Driver),
		InterchainCount: internal.NewElement("xpath", "//*[@id=\"root\"]/section/section/main/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div[2]/div", page.Driver),
		BlockHeight:     internal.NewElement("xpath", "//*[@id=\"root\"]/section/section/main/div/div/div[2]/div[2]/div/div/div[2]/div[3]/div/div[2]/div", page.Driver),
		CurrentDayTx:    internal.NewElement("xpath", "//*[@id=\"root\"]/section/section/main/div/div/div[2]/div[2]/div/div/div[2]/div[4]/div/div[2]/div", page.Driver),
		ServiceCount:    internal.NewElement("xpath", "//*[@id=\"root\"]/section/section/main/div/div/div[2]/div[2]/div/div/div[2]/div[5]/div/div[2]/div", page.Driver),
		LastedBlocks:    internal.NewElements("xpath", "//*[@id=\"root\"]/section/section/main/div/div/div[2]/div[4]/div[1]/div[2]/div/div[1]/div/div[1]/div[1]/div", page.Driver),
		LastedTxs:       internal.NewElements("xpath", "//*[@id=\"root\"]/section/section/main/div/div/div[2]/div[4]/div[2]/div[2]/div/div[1]/div/div[1]/div[1]/div/span", page.Driver),
	}
}
