package BxhUiTest

import (
	"github.com/tebeka/selenium/log"
)

var Config RunConfig

type RunConfig struct {
	DriverType  string
	Port        int
	BxhAddress  string
	LukaAddress string
	Rerun       uint64
	MaxFail     uint64
	Driver      string
	ReportPath  string
	LogLevel    log.Level
}

func init() {
	Config = RunConfig{
		DriverType:  "chrome",
		Port:        9222,
		BxhAddress:  "172.16.30.83:60011",
		LukaAddress: "http://172.16.30.87:8888",
		Rerun:       1,
		MaxFail:     5,
		//Driver:      "/Users/jiuhuche120/code/golang/BxhUITest/driver/chromedriver",
		Driver:     "C:\\Users\\大仓鼠\\go\\src\\BxhUITest\\driver\\chromedriver.exe",
		ReportPath: "C:\\Users\\大仓鼠\\go\\src\\BxhUITest\\reports",
		LogLevel:   log.Debug,
	}
}
