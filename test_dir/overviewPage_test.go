package test_dir

import (
	"BxhUiTest/page"
	"github.com/meshplus/bitxhub-model/constant"
	"strconv"
)

var Page page.OverViewPage

type Model1 struct {
	*Snake
}

func (suite Model1) SetupTest() {
	overviewPage := page.NewOverviewPage()
	err := overviewPage.Get("http://172.16.30.88/overview")
	suite.Require().Nil(err)
	Page = overviewPage
}

func (suite Model1) AfterTest(_, _ string) {
	Page.Close()
}

func (suite Model1) Test0101_ShowChainNum() {
	res, err := suite.client.InvokeBVMContract(constant.AppchainMgrContractAddr.Address(), "CountAppchains", nil)
	suite.Require().Nil(err)
	chainNum, err := Page.ChainNum.Text()
	suite.Require().Nil(err)
	suite.Require().Equal(chainNum, string(res.Ret))
}

func (suite Model1) Test0102_ShowTxNum() {
	meta, err := suite.client.GetChainMeta()
	suite.Require().Nil(err)
	txNum, err := Page.TxNum.Text()
	suite.Require().Nil(err)
	suite.Equal(txNum, strconv.FormatUint(meta.InterchainTxCount, 10))
}

func (suite Model1) Test0103_ShowBlockHeight() {
	meta, err := suite.client.GetChainMeta()
	suite.Require().Nil(err)
	blockHeight, err := Page.BlockHeight.Text()
	suite.Require().Nil(err)
	suite.Equal(blockHeight, strconv.FormatUint(meta.Height, 10))
}
