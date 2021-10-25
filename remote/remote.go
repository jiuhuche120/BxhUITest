package remote

import (
	l "BxhUiTest/remote/luka"
	rpcx "github.com/meshplus/go-bitxhub-client"
)

type Remote struct {
	Client rpcx.Client
	Page   interface{}
	Luka   l.Luka
	Name   string
}
