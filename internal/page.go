package internal

import (
	"BxhUiTest"
	"fmt"
	"github.com/tebeka/selenium"
)

var _ PageObject = (*Page)(nil)

type Page struct {
	Driver selenium.WebDriver
}

func NewPage() Page {
	_, err := selenium.NewChromeDriverService(BxhUiTest.Config.Driver, BxhUiTest.Config.Port, []selenium.ServiceOption{}...)
	if err != nil {
		panic(err)
	}
	var caps selenium.Capabilities
	if BxhUiTest.Config.DriverType == "chrome" {
		caps = map[string]interface{}{"browserName": "chrome"}
	} else {
		caps = map[string]interface{}{"browserName": "firefox"}
	}
	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", BxhUiTest.Config.Port))
	if err != nil {
		panic(err)
	}
	return Page{
		Driver: driver,
	}
}

func (p Page) Back() error {
	err := p.Driver.Back()
	if err != nil {
		return err
	}
	return nil
}

func (p Page) Close() {
	_ = p.Driver.Close()
}

func (p Page) Forward() error {
	err := p.Driver.Forward()
	if err != nil {
		return err
	}
	return nil
}

func (p Page) Get(url string) error {
	err := p.Driver.Get(url)
	if err != nil {
		return err
	}
	return nil
}

func (p Page) GetUrl() (string, error) {
	url, err := p.Driver.CurrentURL()
	if err != nil {
		return "", err
	}
	return url, nil
}

func (p Page) Title() (string, error) {
	title, err := p.Driver.Title()
	if err != nil {
		return "", err
	}
	return title, nil
}

func (p Page) Screenshot() ([]byte, error) {
	pic, err := p.Driver.Screenshot()
	if err != nil {
		return nil, err
	}
	return pic, nil
}

func (p Page) ExecuteScript(script string, args []interface{}) (interface{}, error) {
	res, err := p.Driver.ExecuteScript(script, args)
	if err != nil {
		return nil, err
	}
	return res, nil
}
