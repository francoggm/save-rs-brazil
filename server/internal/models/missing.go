package models

type Missing struct {
	Photos      []string `json:"photos,omitempty"`
	Name        string   `json:"name"`
	Age         int      `json:"age,omitempty"`
	Description string   `json:"description,omitempty"`
	Addr        Address  `json:"address"`
}
