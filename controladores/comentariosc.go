package controladores

import (
	"net/http"
	"github.com/golang/GoComentarios/modelos"
	"encoding/json"
	"fmt"
	"github.com/golang/GoComentarios/comun"
	"github.com/golang/GoComentarios/configuracion"
	"strconv"
	"log"
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
// ListarComentarios obtiene todos los comentarios
func ListarComentarios(w http.ResponseWriter, r *http.Request) {
	// crear un slice de comentarios
	comentarios := []modelos.Comentarios{}
	m := modelos.Mensaje{}
	usuario := modelos.Usuario{}
	// voto := modelos.Voto{}

	r.Context().Value(&usuario)
	vars := r.URL.Query()

	db := configuracion.GetConexion()
	defer db.Close()
	// consulta ccomentario
	cComentario := db.Where("id_padre = 0")

	if order, ok := vars["order"]; ok {
		if order[0] == "votos" {
			cComentario = cComentario.Order("votos desc, created_at desc")
		}
	} else {
		// caso de que el query no retorna
		if idlimit, ok := vars["idlimit"]; ok {
			registerByPage := 30
			// atoi asCCi to integer
			offset, err := strconv.Atoi(idlimit[0])
			if err != nil {
				log.Println("Error", err)
			}
			cComentario = cComentario.Where("id BETWEEN ? AND ?", offset-registerByPage, offset)
		}
		cComentario = cComentario.Order("id desc")
	}

	// ejecutamos la consulta
	cComentario.Find(&comentarios)
	// buscamos la informacion del usuario que comento
	// devolvemos en json
	j, err := json.Marshal(comentarios)
	if err != nil {
		m.CodigoEstado = http.StatusInternalServerError
		m.Mensaje = "error al convertir los comentarios a json"
		comun.MonitoreoMensajes(w, m)
		return
	}
	if len(comentarios) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		// no hay contenido
		m.CodigoEstado = http.StatusNoContent
		m.Mensaje = "no se encontraron comentarios"
		comun.MonitoreoMensajes(w,m)
	}


	// --- /api/comment/?order=votes&idlimit=10

}