package models

type Product struct {
	Id          string  `json:"id"`
	Reference   string  `json:"reference"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
}
