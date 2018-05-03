package controladores

import (
	"net/http"
	"github.com/golang/GoComentarios/modelos"
	"encoding/json"
	"fmt"
	"github.com/golang/GoComentarios/comun"
	"github.com/golang/GoComentarios/configuracion"
	"errors"
)

// RegistroVoto registra los votos realizados por un usuario
func RegistroVoto(w http.ResponseWriter, r *http.Request) {
	voto := modelos.Voto{}
	usuario := modelos.Usuario{}
	currentVoto := modelos.Voto{}
	m := modelos.Mensaje{}

	usuario, _ = r.Context().Value("usuario").(modelos.Usuario)
	err := json.NewDecoder(r.Body).Decode(&voto)
	if err != nil {
		m.Mensaje = fmt.Sprintf("Error al leer el voto a registrar: %s", err)
		m.CodigoEstado = http.StatusBadRequest
		comun.MonitoreoMensajes(w, m)
		return
	}
	voto.UsuarioID = usuario.ID
	db := configuracion.GetConexion()
	defer db.Close()

	db.Where("comentario_id = ? and usuario_id = ?", voto.ComentarioID, voto.UsuarioID).First(&currentVoto)

	// si no existe
	if currentVoto.ID == 0 {
		db.Create(&voto)
		err = actualizarComentariosVotos(voto.ComentarioID, voto.Valor, false)
		if err != nil {
			m.Mensaje = err.Error()
			m.CodigoEstado = http.StatusBadRequest
			comun.MonitoreoMensajes(w, m)
			return
		}
		m.Mensaje = "voto registrado "
		m.CodigoEstado = http.StatusCreated
		comun.MonitoreoMensajes(w, m)
		return
	} else if currentVoto.Valor != voto.Valor {
		currentVoto.Valor = voto.Valor
		db.Save(&currentVoto)
		err := actualizarComentariosVotos(voto.ComentarioID, voto.Valor, true)
		if err != nil {
			m.Mensaje = err.Error()
			m.CodigoEstado = http.StatusBadRequest
			comun.MonitoreoMensajes(w, m)
			return
		}
		m.Mensaje = "voto actualizado"
		m.CodigoEstado = http.StatusOK
		comun.MonitoreoMensajes(w , m)
		return
	}
	m.Mensaje = "Este voto ya esta registrado"
	m.CodigoEstado = http.StatusBadRequest
	comun.MonitoreoMensajes(w, m)
}
// actualizarComentariosVotos actualiza la cantidad de votos en la tabla comentarios
// esActualizacion indica si es un voto para actualizar
func actualizarComentariosVotos(ComentarioID uint, voto bool, esActualizacion bool)(err error) {
	comentario := modelos.Comentarios{}
	db := configuracion.GetConexion()
	defer db.Close()

	rows := db.First(&comentario, ComentarioID).RowsAffected

	if rows > 0 {
		if voto {
			comentario.Votos ++
			if esActualizacion {
				comentario.Votos ++
			}
		} else {
			comentario.Votos --
			if esActualizacion {
				comentario.Votos --
			}
		}
		db.Save(&comentario)
	} else {
		err = errors.New("no se encontro un registro de comentario para asignarle el voto")
	}
	return
 }
