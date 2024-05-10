package models

type Donation struct {
	Photos      []string `json:"photos,omitempty"`
	Itens       []string `json:"itens,omitempty"`
	Description string   `json:"description,omitempty"`
	Addr        Address  `json:"address"`
}
