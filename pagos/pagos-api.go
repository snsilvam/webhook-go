package pagos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/snsilvam/webhook-go/models"
)

// API de pagos.
func IniciarServidorPagos() {
	http.HandleFunc("/pagos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Leer el cuerpo de la solicitud.
			var pago models.Pago
			err := json.NewDecoder(r.Body).Decode(&pago)
			if err != nil {
				http.Error(w, "Error al procesar el pago", http.StatusBadRequest)
				return
			}

			// Simula el procesamiento del pago.
			log.Printf("Pago recibido para el pedido %s: $%.2f", pago.PedidoID, pago.Monto)

			// Enviar notificación de pago recibido, al webhook de pedidos.
			notificacion := models.NotificacionPago{
				PedidoID: pago.PedidoID,
				Estado:   "Pago realizado con exito.",
			}

			// Convertir la notificación a JSON.
			notificacionJSON, _ := json.Marshal(notificacion)

			// Hacer la solicitud HTTTP POST al webhook de pedidos.
			_, err = http.Post("http://localhost:3000/webhook-pago", "application/json", bytes.NewBuffer(notificacionJSON))
			if err != nil {
				http.Error(w, "Error al enviar la notificación de pago al webhook", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Pago procesado y notificacion enviada al webhook de pedidos.")
		}
	})

	// Iniciar el servidor de pagos en el puerto 5000.
	log.Println("Servidor de pagos escuchando en el puerto 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
