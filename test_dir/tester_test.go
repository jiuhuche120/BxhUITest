package test_dir

import (
	"github.com/meshplus/bitxhub-kit/crypto/asym"
	rpcx "github.com/meshplus/go-bitxhub-client"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestTester(t *testing.T) {
	pk, err := asym.RestorePrivateKey("../config/node1.json", "bitxhub")
	require.Nil(t, err)
	node0 := &rpcx.NodeInfo{Addr: cfg.addrs[0]}
	client, err := rpcx.New(
		rpcx.WithNodesInfo(node0),
		rpcx.WithLogger(cfg.logger),
		rpcx.WithPrivateKey(pk),
	)
	require.Nil(t, err)
	suite.Run(t, &Model1{Snake: &Snake{client: client}})
}
