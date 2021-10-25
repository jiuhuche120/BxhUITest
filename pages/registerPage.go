package pages

import "BxhUiTest/internal"

type RegisterPage struct {
	internal.Page
	RegisterChainBtn internal.Element
	ChainID          internal.Element
	ChainName        internal.Element
	Managers         internal.Elements
	AddManageBtn     internal.Element
	ChainType        internal.Element
	Rule             internal.Element
	ChainDesc        internal.Element
	Reason           internal.Element
}

func NewRegisterPage() RegisterPage {
	page := internal.NewPage()
	return RegisterPage{
		Page:             page,
		RegisterChainBtn: internal.NewElement("xpath", "//*[@id=\"root\"]/section/section/main/div/div/div[2]/div/div/div/div/div[3]/button", page.Driver),
		ChainID:          internal.NewElement("xpath", "//*[@id=\"chainId\"]", page.Driver),
		ChainName:        internal.NewElement("xpath", "//*[@id=\"chainName\"]", page.Driver),
		Managers:         internal.NewElements("xpath", "//*[starts-with(@id,'adminAddrs_')]", page.Driver),
		AddManageBtn:     internal.NewElement("xpath", "//*[contains(text(),\"添加\")]", page.Driver),
		ChainType:        internal.NewElement("xpath", "//*[@id=\"type\"]", page.Driver),
		Rule:             internal.NewElement("xpath", "//*[@id=\"rule\"]", page.Driver),
		ChainDesc:        internal.NewElement("xpath", "//*[@id=\"desc\"]", page.Driver),
		Reason:           internal.NewElement("xpath", "//*[@id=\"reason\"]", page.Driver),
	}
}
