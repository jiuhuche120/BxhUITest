package luka

import (
	"BxhUiTest"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Luka struct {
	address string
}

func NewLuck() Luka {
	return Luka{address: BxhUiTest.Config.LukaAddress}
}

func (l Luka) GetMeta() (Meta, error) {
	data, err := l.httpGet(l.address + "/api/v1/meta")
	if err != nil {
		return Meta{}, err
	}
	var meta Meta
	err = json.Unmarshal(data, &meta)
	if err != nil {
		return Meta{}, err
	}
	return meta, err
}

func (l Luka) httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(data))
	}
	return data, nil
}
