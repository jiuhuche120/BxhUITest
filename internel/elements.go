package internel

import (
	"github.com/tebeka/selenium"
	"time"
)

var _ ElementObject = (*Elements)(nil)

type Elements struct {
	typ    string
	value  string
	driver selenium.WebDriver
	index  int
}

func (e Elements) Click() error {
	elements, err := e.findElements()
	if err != nil {
		return err
	}
	if len(elements) <= e.index {
		panic("index out side")
	}
	err = elements[e.index].Click()
	if err != nil {
		return err
	}
	return nil
}

func (e Elements) SendKeys(keys string) error {
	elements, err := e.findElements()
	if err != nil {
		return err
	}
	if len(elements) <= e.index {
		panic("index out side")
	}
	err = elements[e.index].SendKeys(keys)
	if err != nil {
		return err
	}
	return nil
}

func (e Elements) Text() (string, error) {
	elements, err := e.findElements()
	if err != nil {
		return "", err
	}
	if len(elements) <= e.index {
		panic("index out side")
	}
	text, err := elements[e.index].Text()
	if err != nil {
		return "", err
	}
	return text, nil
}

func (e Elements) Get(i int) Elements {
	e.index = i
	return e
}

func (e Elements) findElements() ([]selenium.WebElement, error) {
	var elements []selenium.WebElement
	var err error
	err = e.driver.SetImplicitWaitTimeout(time.Second * 10)
	if err != nil {
		return nil, err
	}
	elements, err = e.driver.FindElements(e.typ, e.value)
	if err != nil {
		return nil, err
	}
	return elements, err
}
