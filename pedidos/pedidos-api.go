package pedidos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/snsilvam/webhook-go/models"
)

// API de pedidos.
func IniciarServidorPedidos() {
	http.HandleFunc("/pedidos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Leer el cuerpo de la solicitud
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
				return
			}
			defer r.Body.Close()

			// Parsear el JSON del cuerpo
			var pedido models.Pedido
			err = json.Unmarshal(body, &pedido)
			if err != nil {
				http.Error(w, "Error al parsear el JSON", http.StatusBadRequest)
				return
			}

			// Almacenar el pedido
			pedidos[pedido.ID] = pedido

			// Responder con el pedido almacenado
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "Pedido almacenado: %+v", pedido)
		}
	})

	http.HandleFunc("/webhook-pago", func(w http.ResponseWriter, r *http.Request) {
		// Solo se permite solicitudes POST
		if r.Method != http.MethodPost {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		// Leer el cuerpo de la solicitud
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// Parsear el JSON del cuerpo
		var notificacion models.NotificacionPago
		err = json.Unmarshal(body, &notificacion)
		if err != nil {
			http.Error(w, "Error al parsear el JSON", http.StatusBadRequest)
			return
		}

		// Procesar la notificación
		if pedido, exists := pedidos[notificacion.PedidoID]; exists {
			log.Printf("Pago recibido para el pedido ID %s: %s", pedido.ID, notificacion.Estado)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Pago procesado correctamente para el pedido: %v", pedido)
		} else {
			http.Error(w, "Pedido no encontrado", http.StatusNotFound)
		}
	})

	// Iniciar el servidor de pedidos en el puerto 3000
	log.Println("Servidor de pedidos escuchando en el puerto 3000")
	fmt.Println("Servidor de pedidos escuchando en el puerto 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// Almacenamiento de pedido en memoria.
var pedidos = make(map[string]models.Pedido)
