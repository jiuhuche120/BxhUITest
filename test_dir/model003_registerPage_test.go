package test_dir

import (
	"BxhUiTest/pages"
	"github.com/meshplus/bitxhub-kit/crypto"
	"github.com/meshplus/bitxhub-kit/crypto/asym"
	"time"
)

type Model3 struct {
	*Snake
}

func (suite Model3) NewPage() pages.RegisterPage {
	page := pages.NewRegisterPage()
	err := page.Get("http://172.16.30.87/overview")
	suite.Nil(err)
	time.Sleep(time.Second)
	return page
}

func (suite Model3) Test0301_RegisterAppchainIsSuccess() {
	page := suite.NewPage()
	defer page.Close()
	pk, err := asym.GenerateKeyPair(crypto.Secp256k1)
	suite.Nil(err)
	address, err := pk.PublicKey().Address()
	suite.Nil(err)
	err = suite.Login02(pk, page.Driver)
	suite.Nil(err)
	err = page.RegisterChainBtn.Click()
	suite.Nil(err)
	err = page.ChainID.SendKeys("")
	suite.Nil(err)
	err = page.ChainName.SendKeys("")
	suite.Nil(err)
	err = page.ChainType.SendKeys("Flato V1.0.0")
	suite.Nil(err)
}
