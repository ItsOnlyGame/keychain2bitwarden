package utils

import "github.com/guregu/null"

type BitwardenFile struct {
	Items []*BitwardenItem `json:"items"`
}

type BitwardenItemLoginUris struct {
	Match null.String `json:"match"`
	Uri   string      `json:"uri"`
}

type BitwardenItemLogin struct {
	Uris     []BitwardenItemLoginUris `json:"uris"`
	Username string                   `json:"username"`
	Password string                   `json:"password"`
	OTPAuth  string                   `json:"totp"`
}

type BitwardenItem struct {
	Type  uint16             `json:"type"`
	Name  string             `json:"name"`
	Login BitwardenItemLogin `json:"login"`
}
