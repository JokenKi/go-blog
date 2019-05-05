package model

import (
	"encoding/json"
)

// NotifyInfo notify info.
type CustomerInfo struct {
	Cid string          `json:"cid"`
	Phone  string          `json:"phone"`
	Password    json.RawMessage `json:"password"`
	NickName    json.RawMessage `json:"nick_name"`
}
