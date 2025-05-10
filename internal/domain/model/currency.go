package model

type Currency struct {
	Code   string  `json:"code"`
	Name   string  `json:"name"`
	Symbol *string `json:"symbol,omitempty"`
}
