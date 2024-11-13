package models

type Stream struct {
	ID    string `json:"id"`
	Data  []byte `json:"data"`
	State string `json:"state"`
}
