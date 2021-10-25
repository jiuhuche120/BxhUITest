package rpcx

import (
	"BxhUiTest"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/meshplus/bitxhub-kit/crypto"
	"github.com/meshplus/bitxhub-kit/crypto/asym"
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/meshplus/bitxhub-model/pb"
	rpcx "github.com/meshplus/go-bitxhub-client"
	"github.com/sirupsen/logrus"
)

var Nonce1 uint64
var Nonce2 uint64
var Nonce3 uint64
var Nonce4 uint64

type config struct {
	addrs  []string
	logger rpcx.Logger
}

var cfg = &config{
	addrs: []string{
		BxhUiTest.Config.BxhAddress,
	},
	logger: logrus.New(),
}

func NewBxhClient() rpcx.Client {
	pk, err := asym.GenerateKeyPair(crypto.Secp256k1)
	if err != nil {
		panic(err)
	}
	from, err := pk.PublicKey().Address()
	node0 := &rpcx.NodeInfo{Addr: cfg.addrs[0]}
	client, err := rpcx.New(
		rpcx.WithNodesInfo(node0),
		rpcx.WithLogger(cfg.logger),
		rpcx.WithPrivateKey(pk),
	)
	err = TransferFromAdmin(from.String(), "1")
	if err != nil {
		panic(err)
	}
	return client
}

func TransferFromAdmin(address string, amount string) error {
	pk, err := asym.RestorePrivateKey("../config/node1.json", "bitxhub")
	if err != nil {
		return err
	}
	from, err := pk.PublicKey().Address()
	if err != nil {
		return err
	}
	node0 := &rpcx.NodeInfo{Addr: cfg.addrs[0]}
	client, err := rpcx.New(
		rpcx.WithNodesInfo(node0),
		rpcx.WithLogger(cfg.logger),
		rpcx.WithPrivateKey(pk),
	)
	if err != nil {
		return err
	}
	data := &pb.TransactionData{
		Amount: amount + "000000000000000000",
	}
	payload, err := data.Marshal()
	if err != nil {
		return err
	}

	tx := &pb.BxhTransaction{
		From:      from,
		To:        types.NewAddressByStr(address),
		Timestamp: time.Now().UnixNano(),
		Payload:   payload,
	}
	ret, err := client.SendTransactionWithReceipt(tx, &rpcx.TransactOpts{
		From:  from.String(),
		Nonce: atomic.AddUint64(&Nonce1, 1) - 1,
	})
	if err != nil {
		return err
	}
	if ret.Status != pb.Receipt_SUCCESS {
		return fmt.Errorf(string(ret.Ret))
	}
	return nil
}
