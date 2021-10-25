package pages

import (
	"BxhUiTest/internal"
)

type LoginPage struct {
	internal.Page
	Login      internal.Element
	Password   internal.Element
	Repassword internal.Element
	Key        internal.Element
	Confirm    internal.Element
}

func NewLoginPage() LoginPage {
	page := internal.NewPage()
	return LoginPage{
		Page:       page,
		Login:      internal.NewElement("xpath", "//*[@id=\"root\"]/section/header/div[2]/div[2]", page.Driver),
		Password:   internal.NewElement("xpath", "//*[@id=\"password\"]", page.Driver),
		Repassword: internal.NewElement("xpath", "//*[@id=\"confirm\"]", page.Driver),
		Key:        internal.NewElement("xpath", "//*[@id=\"upload\"]", page.Driver),
		Confirm:    internal.NewElement("xpath", "/html/body/div[2]/div/div[2]/div/div[2]/div/form/div[4]/div/div/div/button", page.Driver),
	}
}
