package internal

import (
	"github.com/tebeka/selenium"
)

var _ selenium.WebElement = (*Element)(nil)

type Element struct {
	by     string
	value  string
	driver selenium.WebDriver
}

func (e Element) Click() error {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return err
	}
	return element.Click()
}

func (e Element) SendKeys(keys string) error {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return err
	}
	return element.SendKeys(keys)
}

func (e Element) Submit() error {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return err
	}
	return element.Submit()
}

func (e Element) Clear() error {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return err
	}
	return element.Clear()
}

func (e Element) MoveTo(xOffset, yOffset int) error {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return err
	}
	return element.MoveTo(xOffset, yOffset)
}

func (e Element) FindElement(by, value string) (selenium.WebElement, error) {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return nil, err
	}
	return element.FindElement(by, value)
}

func (e Element) FindElements(by, value string) ([]selenium.WebElement, error) {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return nil, err
	}
	return element.FindElements(by, value)
}

func (e Element) TagName() (string, error) {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return "", err
	}
	return element.TagName()
}

func (e Element) Text() (string, error) {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return "", err
	}
	return element.Text()
}

func (e Element) IsSelected() (bool, error) {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return false, err
	}
	return element.IsSelected()
}

func (e Element) IsEnabled() (bool, error) {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return false, err
	}
	return element.IsEnabled()
}

func (e Element) IsDisplayed() (bool, error) {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return false, err
	}
	return element.IsDisplayed()
}

func (e Element) GetAttribute(name string) (string, error) {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return "", err
	}
	return element.GetAttribute(name)
}

func (e Element) Location() (*selenium.Point, error) {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return nil, err
	}
	return element.Location()
}

func (e Element) LocationInView() (*selenium.Point, error) {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return nil, err
	}
	return element.LocationInView()
}

func (e Element) Size() (*selenium.Size, error) {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return nil, err
	}
	return element.Size()
}

func (e Element) CSSProperty(name string) (string, error) {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return "", err
	}
	return element.CSSProperty(name)
}

func (e Element) Screenshot(scroll bool) ([]byte, error) {
	element, err := e.driver.FindElement(e.by, e.value)
	if err != nil {
		return []byte{}, err
	}
	return element.Screenshot(scroll)
}
