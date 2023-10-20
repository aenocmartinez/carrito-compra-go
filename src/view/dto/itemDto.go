package dto

type ItemDto struct {
	Id     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Amount int    `json:"amount,omitempty"`
}
