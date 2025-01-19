package models

type NotificacionPago struct {
	PedidoID string `json:"pedido_id"`
	Estado   string `json:"estado"`
}
