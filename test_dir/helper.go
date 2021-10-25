package test_dir

import (
	"BxhUiTest"
	"BxhUiTest/internal"
	"BxhUiTest/pages"
	"BxhUiTest/remote"
	"BxhUiTest/remote/luka"
	"BxhUiTest/remote/rpcx"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/meshplus/bitxhub-kit/crypto"
	"github.com/meshplus/bitxhub-kit/crypto/asym"
	"github.com/stretchr/testify/suite"
	"github.com/tebeka/selenium"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Snake struct {
	suite.Suite
}

func (suite Snake) NewClient(page interface{}) remote.Remote {
	return remote.Remote{
		Client: rpcx.NewBxhClient(),
		Page:   page,
		Luka:   luka.NewLuck(),
		Name:   suite.T().Name(),
	}
}

func (suite Snake) SetupSuite() {
	client := rpcx.NewBxhClient()
	nonce, err := client.GetPendingNonceByAccount("0xc7F999b83Af6DF9e67d0a37Ee7e900bF38b3D013")
	suite.Require().Nil(err)
	rpcx.Nonce1 = nonce
	nonce, err = client.GetPendingNonceByAccount("0xc7F999b83Af6DF9e67d0a37Ee7e900bF38b3D013")
	suite.Require().Nil(err)
	rpcx.Nonce2 = nonce
	nonce, err = client.GetPendingNonceByAccount("0xc7F999b83Af6DF9e67d0a37Ee7e900bF38b3D013")
	suite.Require().Nil(err)
	rpcx.Nonce3 = nonce
	nonce, err = client.GetPendingNonceByAccount("0xc7F999b83Af6DF9e67d0a37Ee7e900bF38b3D013")
	suite.Require().Nil(err)
	rpcx.Nonce4 = nonce
}

func (suite Snake) MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func (suite Snake) Login(pk crypto.PrivateKey, page internal.PageObject) error {
	_, err := page.ExecuteScript("localStorage.setItem('password','"+suite.MD5("a123456")+"')", nil)
	if err != nil {
		return err
	}
	bytes, err := pk.Bytes()
	if err != nil {
		return err
	}
	privKey := common.Bytes2Hex(bytes)
	address, err := pk.PublicKey().Address()
	if err != nil {
		return err
	}
	_, err = page.ExecuteScript("localStorage.setItem('priv','"+privKey+"')", nil)
	if err != nil {
		return err
	}
	json := "{\"privateKey\":\"" + privKey + "\",\"address\":\"" + address.String()[2:] + "\"}"
	_, err = page.ExecuteScript("sessionStorage.setItem('SDK_STORAGE','"+json+"')", nil)
	if err != nil {
		return err
	}
	return nil
}

func (suite Snake) Login02(pk crypto.PrivateKey, driver selenium.WebDriver) error {
	//TODO:use local storage login
	header := pages.NewHeader(driver)
	err := header.Login.Click()
	if err != nil {
		return err
	}
	err = header.Password.SendKeys("a123456")
	if err != nil {
		return err
	}
	err = header.Repassword.SendKeys("a123456")
	if err != nil {
		return err
	}
	_, err = os.Stat("../config/key.json")
	if err == nil {
		err := os.Remove("../config/key.json")
		if err != nil {
			return err
		}
	}
	err = asym.StorePrivateKey(pk, "../config/key.json", "bitxhub")
	if err != nil {
		return err
	}
	err = header.Key.SendKeys("C:\\Users\\大仓鼠\\go\\src\\BxhUITest\\config\\key.json")
	if err != nil {
		return err
	}
	err = header.Confirm.Click()
	if err != nil {
		return err
	}
	err = header.ChainManage.Click()
	suite.Nil(err)
	return nil
}

// Equal use args[0] as internal.PageObject, args[1:] as msgAndArgs
func (suite Snake) Equal(expected interface{}, actual interface{}, args ...interface{}) {
	var page internal.PageObject
	var msgAndArgs []interface{}
	if args != nil {
		page = args[0].(internal.PageObject)
		msgAndArgs = args[1:]
	}
	if !suite.Assert().Equal(expected, actual, msgAndArgs) && page != nil {
		err := suite.doSome(page)
		if err != nil {
			suite.Assert().Error(err)
		}
	}
}

// NotEqual use args[0] as internal.Page, args[1:] as msgAndArgs
func (suite Snake) NotEqual(expected interface{}, actual interface{}, args ...interface{}) {
	var page internal.PageObject
	var msgAndArgs []interface{}
	if args != nil {
		page = args[0].(internal.PageObject)
		msgAndArgs = args[1:]
	}
	if !suite.Assert().NotEqual(expected, actual, msgAndArgs) && page != nil {
		err := suite.doSome(page)
		if err != nil {
			suite.Assert().Error(err)
		}
	}
}

// Nil use args[0] as internal.Page, args[1:] as msgAndArgs
func (suite Snake) Nil(object interface{}, args ...interface{}) {
	var page internal.PageObject
	var msgAndArgs []interface{}
	if args != nil {
		page = args[0].(internal.PageObject)
		msgAndArgs = args[1:]
	}
	if !suite.Assert().Nil(object, msgAndArgs) && page != nil {
		err := suite.doSome(page)
		if err != nil {
			suite.Assert().Error(err)
		}
	}
}

// NotNil use args[0] as internal.Page, args[1:] as msgAndArgs
func (suite Snake) NotNil(object interface{}, args ...interface{}) {
	var page internal.PageObject
	var msgAndArgs []interface{}
	if args != nil {
		page = args[0].(internal.PageObject)
		msgAndArgs = args[1:]
	}
	if !suite.Assert().NotNil(object, msgAndArgs) && page != nil {
		err := suite.doSome(page)
		if err != nil {
			suite.Assert().Error(err)
		}
	}
}

// GreaterOrEqual use args[0] as internal.Page, args[1:] as msgAndArgs
func (suite Snake) GreaterOrEqual(e1 interface{}, e2 interface{}, args ...interface{}) {
	var page internal.PageObject
	var msgAndArgs []interface{}
	if args != nil {
		page = args[0].(internal.PageObject)
		msgAndArgs = args[1:]
	}
	if !suite.Assert().GreaterOrEqual(e1, e2, msgAndArgs) && page != nil {
		err := suite.doSome(page)
		if err != nil {
			suite.Assert().Error(err)
		}
	}
}

// LessOrEqual use args[0] as internal.Page, args[1:] as msgAndArgs
func (suite Snake) LessOrEqual(e1 interface{}, e2 interface{}, args ...interface{}) {
	var page internal.PageObject
	var msgAndArgs []interface{}
	if args != nil {
		page = args[0].(internal.PageObject)
		msgAndArgs = args[1:]
	}
	if !suite.Assert().LessOrEqual(e1, e2, msgAndArgs) && page != nil {
		err := suite.doSome(page)
		if err != nil {
			suite.Assert().Error(err)
		}
	}
}

//TODO:make report and send mail
func (suite Snake) doSome(page internal.PageObject) error {
	err := suite.doPicture(page)
	if err != nil {
		return err
	}
	return nil
}

func (suite Snake) doPicture(page internal.PageObject) error {
	path := filepath.Join(BxhUiTest.Config.ReportPath, fmt.Sprintf("%v", time.Now().Format("2006-01-02")))
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0777)
		if err != nil {
			return err
		}
	}
	filename := strings.Split(suite.T().Name(), "/")[1] + ".png"
	data, err := page.Screenshot()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(path, filename), data, 0777)
	if err != nil {
		return err
	}
	return nil
}
