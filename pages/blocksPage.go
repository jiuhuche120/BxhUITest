package pages

import "BxhUiTest/internal"

type BlocksPage struct {
	internal.Page
	Confirm       internal.Element
	SearchInput   internal.Element
	SearchConfirm internal.Element
	BlocksHeight  internal.Elements
	BlockHash     internal.Element
	EmptyMsg      internal.Element
}

func NewBlocksPage() BlocksPage {
	page := internal.NewPage()
	return BlocksPage{
		Page:          page,
		Confirm:       internal.NewElement("xpath", "/html/body/div[3]/div/div[2]/div/div[2]/div/div/div[2]/button", page.Driver),
		SearchInput:   internal.NewElement("xpath", "//*[@id=\"query\"]", page.Driver),
		SearchConfirm: internal.NewElement("xpath", "//*[@id=\"root\"]/section/section/main/div/div[2]/div/div/form/div/div[2]/div/div/div/button", page.Driver),
		BlocksHeight:  internal.NewElements("xpath", "//*[@id=\"root\"]/section/section/main/div/div[3]/div/div[2]/div/div/div/div/div/table/tbody/tr/td[1]/div", page.Driver),
		BlockHash:     internal.NewElement("xpath", "//*[@id=\"root\"]/section/section/main/div[2]/div/div/div[2]/div[2]", page.Driver),
		EmptyMsg:      internal.NewElement("xpath", "//*[@id=\"root\"]/section/section/main/div[3]/div/div[2]", page.Driver),
	}
}
