package internel

type PageObject interface {
	Get(url string) error
	Title() (string, error)
	GetUrl() (string, error)
	ScreenShots(path, filename string) error
	Close()
}
