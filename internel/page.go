package internel

import (
	"BxhUiTest"
	"fmt"
	"github.com/tebeka/selenium"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
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

func (p Page) Get(url string) error {
	err := p.Driver.Get(url)
	if err != nil {
		return err
	}
	return nil
}

func (p Page) Title() (string, error) {
	title, err := p.Driver.Title()
	if err != nil {
		return "", err
	}
	return title, nil
}

func (p Page) GetUrl() (string, error) {
	url, err := p.Driver.CurrentURL()
	if err != nil {
		return "", err
	}
	return url, nil
}

func (p Page) ScreenShots(path, filename string) error {
	if path == "" {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		path = dir
	}
	if filename == "" {
		filename = time.Now().Format("20060102150102") + ".png"
	}
	path = filepath.Join(path, filename)
	data, err := p.Driver.Screenshot()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, data, 0777)
	if err != nil {
		return err
	}
	return nil
}

func (p Page) Close() {
	_ = p.Driver.Close()
}
