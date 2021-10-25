package pages

import "BxhUiTest/internal"

type TxsViewPage struct {
	internal.Page
	Confirm       internal.Element
	GoTxs         internal.Element
	SearchInput   internal.Element
	SearchConfirm internal.Element
}

func NewTxsPage() TxsViewPage {
	page := internal.NewPage()
	return TxsViewPage{
		Page:          page,
		Confirm:       internal.NewElement("xpath", "/html/body/div[3]/div/div[2]/div/div[2]/div/div/div[2]/button", page.Driver),
		GoTxs:         internal.NewElement("xpath", "//*[@id=\"root\"]/section/section/main/div/div[3]/div/div/div[2]/div[2]/div", page.Driver),
		SearchInput:   internal.NewElement("xpath", "//*[@id=\"query\"]", page.Driver),
		SearchConfirm: internal.NewElement("xpath", "//*[@id=\"root\"]/section/section/main/div/div[2]/div/div/form/div/div[2]/div/div/div/button", page.Driver),
	}
}
