package internel

type ElementObject interface {
	Click() error
	SendKeys(keys string) error
	Text() (string, error)
}
