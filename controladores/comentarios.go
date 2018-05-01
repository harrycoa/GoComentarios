package controladores

import (
	"net/http"
	"github.com/golang/GoComentarios/modelos"
	"encoding/json"
	"fmt"
	"github.com/golang/GoComentarios/comun"
	"github.com/golang/GoComentarios/configuracion"
)
// CrearComentario crea un comentario
func CrearComentario(w http.ResponseWriter, r *http.Request){
	comentario := modelos.Comentarios{}
	m := modelos.Mensaje{}
	err := json.NewDecoder(r.Body).Decode(&comentario)
	if err != nil {
		m.CodigoEstado = http.StatusBadRequest
		m.Mensaje = fmt.Sprintf("error al leer el comentario: %s", err)
		comun.MonitoreoMensajes(w ,m)
		return
	}
	db := configuracion.GetConexion()
	defer db.Close()
	// crear el registro
	err = db.Create(&comentario).Error
	if err != nil {
		m.CodigoEstado = http.StatusBadRequest
		m.Mensaje = fmt.Sprintf("error al registrar el comentario: %s", err)
		comun.MonitoreoMensajes(w ,m)
		return
	}
	m.CodigoEstado = http.StatusCreated
	m.Mensaje = "Comentario creado con exito"
	comun.MonitoreoMensajes(w, m)
}