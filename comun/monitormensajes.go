package comun

import (
	"encoding/json"
	"github.com/golang/GoComentarios/modelos"
	"log"
	"net/http"
)

// MonitoreoMensajes devuelve un mensaje al cliente
func MonitoreoMensajes(w http.ResponseWriter, m modelos.Mensaje) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("error al convertir el mensaje a json: %s", err)
	}
	w.WriteHeader(m.CodigoEstado)
	w.Write(j)
}
