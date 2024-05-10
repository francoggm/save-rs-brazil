package models

type Cleaning struct {
	Photos             []string `json:"photos,omitempty"`
	HaveCleaningSupply bool     `json:"have_cleaning_supply,omitempty"`
	Description        string   `json:"description,omitempty"`
	Addr               Address  `json:"address"`
}
