package test_dir

import (
	"BxhUiTest/pages"
	"time"
)

type Model2 struct {
	*Snake
}

func (suite Model2) NewPage() pages.LoginPage {
	page := pages.NewLoginPage()
	err := page.Get("http://172.16.30.87/overview")
	suite.Nil(err)
	time.Sleep(time.Second)
	err = page.Login.Click()
	suite.Nil(err)
	return page
}
func (suite Model2) Test0201_LoginRelayAdminIsSuccess() {
	page := suite.NewPage()
	defer page.Close()
	err := page.Password.SendKeys("a123456")
	suite.Nil(err)
	err = page.Repassword.SendKeys("a123456")
	suite.Nil(err)
	err = page.Key.SendKeys("C:\\Users\\大仓鼠\\go\\src\\BxhUITest\\config\\node1.json")
	suite.Nil(err)
	err = page.Confirm.Click()
	time.Sleep(time.Second)
	res, err := page.ExecuteScript("return localStorage.password;", nil)
	suite.Nil(err)
	suite.Equal(res, suite.MD5("a123456"), page)
}

func (suite Model2) Test0202_LoginRelayWithEmptyPwdIsFail() {
	page := suite.NewPage()
	defer page.Close()
	err := page.Repassword.SendKeys("a123456")
	suite.Nil(err)
	err = page.Key.SendKeys("C:\\Users\\大仓鼠\\go\\src\\BxhUITest\\config\\node1.json")
	suite.Nil(err)
	err = page.Confirm.Click()
	time.Sleep(time.Second)
	res, err := page.ExecuteScript("return localStorage.password;", nil)
	suite.Nil(err)
	suite.Nil(res, page)
}

func (suite Model2) Test0203_LoginRelayWithEmptyRePwdIsFail() {
	page := suite.NewPage()
	defer page.Close()
	err := page.Password.SendKeys("a123456")
	suite.Nil(err)
	err = page.Key.SendKeys("C:\\Users\\大仓鼠\\go\\src\\BxhUITest\\config\\node1.json")
	suite.Nil(err)
	err = page.Confirm.Click()
	time.Sleep(time.Second)
	res, err := page.ExecuteScript("return localStorage.password;", nil)
	suite.Nil(err)
	suite.Nil(res, page)
}

func (suite Model2) Test0204LoginRelayWithEmptyKeyIsFail() {
	page := suite.NewPage()
	defer page.Close()
	err := page.Password.SendKeys("a123456")
	suite.Nil(err)
	err = page.Repassword.SendKeys("a123456")
	suite.Nil(err)
	err = page.Confirm.Click()
	time.Sleep(time.Second)
	res, err := page.ExecuteScript("return localStorage.password;", nil)
	suite.Nil(err)
	suite.Nil(res, page)
}

func (suite Model2) Test0205_LoginRelayWithDiffPwdIsFail() {
	page := suite.NewPage()
	defer page.Close()
	err := page.Password.SendKeys("a123456")
	suite.Nil(err)
	err = page.Repassword.SendKeys("b123456")
	suite.Nil(err)
	err = page.Key.SendKeys("C:\\Users\\大仓鼠\\go\\src\\BxhUITest\\config\\node1.json")
	suite.Nil(err)
	err = page.Confirm.Click()
	time.Sleep(time.Second)
	res, err := page.ExecuteScript("return localStorage.password;", nil)
	suite.Nil(err)
	suite.Nil(res, page)
}
