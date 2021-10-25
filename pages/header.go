package pages

import (
	"BxhUiTest/internal"
	"github.com/tebeka/selenium"
)

type Header struct {
	Login       internal.Element
	Password    internal.Element
	Repassword  internal.Element
	Key         internal.Element
	Confirm     internal.Element
	ChainManage internal.Element
}

func NewHeader(driver selenium.WebDriver) Header {
	return Header{
		Login:       internal.NewElement("xpath", "//*[@id=\"root\"]/section/header/div[2]/div[2]", driver),
		Password:    internal.NewElement("xpath", "//*[@id=\"password\"]", driver),
		Repassword:  internal.NewElement("xpath", "//*[@id=\"confirm\"]", driver),
		Key:         internal.NewElement("xpath", "//*[@id=\"upload\"]", driver),
		Confirm:     internal.NewElement("xpath", "/html/body/div[2]/div/div[2]/div/div[2]/div/form/div[4]/div/div/div/button", driver),
		ChainManage: internal.NewElement("xpath", "//*[@id=\"root\"]/section/section/aside/div[1]/ul/li[2]/span/div/span[2]", driver),
	}
}
