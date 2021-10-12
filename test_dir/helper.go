package test_dir

import (
	rpcx "github.com/meshplus/go-bitxhub-client"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

var cfg = &config{
	addrs: []string{
		"172.16.30.83:60011",
		"localhost:60012",
		"localhost:60013",
		"localhost:60014",
	},
	logger: logrus.New(),
}

type config struct {
	addrs  []string
	logger rpcx.Logger
}
type Snake struct {
	suite.Suite
	client rpcx.Client
}
