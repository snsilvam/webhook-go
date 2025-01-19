package models

type Pedido struct {
	ID     string  `json:"id"`
	Plato  string  `json:"plato"`
	Precio float64 `json:"precio"`
}
