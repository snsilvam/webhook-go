package main

import (
	"fmt"

	"github.com/snsilvam/webhook-go/pagos"
	"github.com/snsilvam/webhook-go/pedidos"
)

func main() {
	fmt.Println("Hello, World!")
	go pedidos.IniciarServidorPedidos()
	pagos.IniciarServidorPagos()
}
