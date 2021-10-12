package internel

import (
	"github.com/tebeka/selenium"
	"time"
)

var _ ElementObject = (*Element)(nil)

type Element struct {
	typ     string
	value   string
	driver  selenium.WebDriver
}

func NewElement(typ, value string, driver selenium.WebDriver) Element {
	return Element{
		typ:     typ,
		value:   value,
		driver:  driver,
	}
}

func (e Element) Click() error {
	element, err := e.findElement()
	if err != nil {
		return err
	}
	err = element.Click()
	if err != nil {
		return err
	}
	return err
}

func (e Element) SendKeys(keys string) error {
	element, err := e.findElement()
	if err != nil {
		return err
	}
	err = element.SendKeys(keys)
	if err != nil {
		return err
	}
	return nil
}

func (e Element) Text() (string, error) {
	element, err := e.findElement()
	if err != nil {
		return "", err
	}
	text, err := element.Text()
	if err != nil {
		return "", err
	}
	return text, nil
}

func (e Element) findElement() (selenium.WebElement, error) {
	var element selenium.WebElement
	var err error
	err = e.driver.SetImplicitWaitTimeout(time.Second * 10)
	if err != nil {
		return nil, err
	}
	element, err = e.driver.FindElement(e.typ, e.value)
	if err != nil {
		return nil, err
	}
	return element, err
}
