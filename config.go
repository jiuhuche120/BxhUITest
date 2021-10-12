package BxhUiTest

import (
	"github.com/tebeka/selenium/log"
)

var Config RunConfig

type RunConfig struct {
	DriverType string
	Port       int
	BxhAddress string
	Rerun      uint64
	MaxFail    uint64
	Driver     string
	ReportPath string
	LogLevel   log.Level
}

func init() {
	Config = RunConfig{
		DriverType: "chrome",
		Port:       9222,
		BxhAddress: "http://127.0.0.1:60011",
		Rerun:      1,
		MaxFail:    5,
		Driver:     "C:\\Users\\大仓鼠\\go\\src\\github.com\\jiuhuche120\\BxhUiTest\\driver\\chromedriver.exe",
		ReportPath: "",
		LogLevel:   log.Info,
	}
}
