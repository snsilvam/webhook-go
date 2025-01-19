package models

type Pago struct {
	PedidoID string  `json:"pedido_id"`
	Monto    float64 `json:"monto"`
}
