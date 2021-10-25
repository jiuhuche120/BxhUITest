package luka

type Meta struct {
	Code uint64 `json:"code,omitempty"`
	Data Data   `json:"data"`
}
type Data struct {
	AppchainCount     uint64 `json:"appchainCount,omitempty"`
	BlockHeight       uint64 `json:"blockHeight,omitempty"`
	CurrentDayTxCount uint64 `json:"currentDayTxCount,omitempty"`
	InterchainCount   uint64 `json:"interchainCount,omitempty"`
	ServiceCount      uint64 `json:"serviceCount,omitempty"`
}
