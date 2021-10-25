package test_dir

import (
	"BxhUiTest/pages"
	"strconv"
	"time"

	"github.com/meshplus/bitxhub-model/constant"
	"github.com/meshplus/bitxhub-model/pb"
)

type Model1 struct {
	*Snake
}

func (suite Model1) SetupTest() {
	suite.T().Parallel()
}

func (suite Model1) NewPage() pages.OverViewPage {
	page := pages.NewOverviewPage()
	err := page.Get("http://172.16.30.87/overview")
	suite.Nil(err)
	time.Sleep(time.Second)
	return page
}

func (suite Model1) Test0101_ShowAppchainCount() {
	page := suite.NewPage()
	defer page.Close()
	bxh := suite.NewClient(page)
	res, err := bxh.Client.InvokeBVMContract(constant.AppchainMgrContractAddr.Address(), "CountAppchains", nil)
	suite.Nil(err)
	chainNum, err := page.AppchainCount.Text()
	suite.Nil(err)
	suite.Equal(pb.Receipt_SUCCESS, res.Status)
	suite.Equal(string(res.Ret), chainNum, page)
}

func (suite Model1) Test0102_ShowInterchainCount() {
	page := suite.NewPage()
	defer page.Close()
	bxh := suite.NewClient(page)
	meta, err := bxh.Client.GetChainMeta()
	suite.Nil(err)
	txNum, err := page.InterchainCount.Text()
	suite.Nil(err)
	suite.Equal(strconv.FormatUint(meta.InterchainTxCount, 10), txNum, page)
}

func (suite Model1) Test0103_ShowBlockHeight() {
	page := suite.NewPage()
	defer page.Close()
	bxh := suite.NewClient(page)
	meta, err := bxh.Client.GetChainMeta()
	suite.Nil(err)
	blockHeight, err := page.BlockHeight.Text()
	suite.Nil(err)
	suite.Require().GreaterOrEqual(strconv.FormatUint(meta.Height, 10), blockHeight, page)
}

func (suite Model1) Test0104_ShowCurrentDayTxCount() {
	page := suite.NewPage()
	defer page.Close()
	bxh := suite.NewClient(page)
	meta, err := bxh.Luka.GetMeta()
	suite.Nil(err)
	count, err := page.CurrentDayTx.Text()
	suite.Nil(err)
	suite.Equal(strconv.FormatUint(meta.Data.CurrentDayTxCount, 10), count, page)
}

func (suite Model1) Test0105_ShowServiceCount() {
	page := suite.NewPage()
	defer page.Close()
	bxh := suite.NewClient(page)
	meta, err := bxh.Luka.GetMeta()
	suite.Nil(err)
	count, err := page.ServiceCount.Text()
	suite.Nil(err)
	suite.Equal(strconv.FormatUint(meta.Data.ServiceCount, 10), count, page)
}

func (suite Model1) Test0106_ShowLastedBlocks() {
	page := suite.NewPage()
	defer page.Close()
	length := page.LastedBlocks.Length()
	suite.NotEqual(0, length, page)
}

func (suite Model1) Test0107_ShowLastedTxs() {
	page := suite.NewPage()
	defer page.Close()
	length := page.LastedTxs.Length()
	suite.NotEqual(0, length, page)
}
