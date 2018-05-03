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
	"github.com/olahol/melody"
	//"github.com/gorilla/websocket"
	"golang.org/x/net/websocket"
)

// Melody permite utilizar realtime
var Melody *melody.Melody
// CrearComentario crea un comentario

func init(){
	Melody = melody.New()
}

func CrearComentario(w http.ResponseWriter, r *http.Request){
	comentario := modelos.Comentarios{}
	usuario := modelos.Usuario{}
	m := modelos.Mensaje{}

	usuario, _ = r.Context().Value("usuario").(modelos.Usuario)
	err := json.NewDecoder(r.Body).Decode(&comentario)
	if err != nil {
		m.CodigoEstado = http.StatusBadRequest
		m.Mensaje = fmt.Sprintf("error al leer el comentario: %s", err)
		comun.MonitoreoMensajes(w ,m)
		return
	}
	comentario.IDUsuario = usuario.ID


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

	db.Model(&comentario).Related(&comentario.Usuarios)
	comentario.Usuarios[0].Contrasenia = ""

	j, err := json.Marshal(&comentario)
	if err != nil {
		m.Mensaje = fmt.Sprintf("no se pudo convertir el comentario a json %s", err)
		m.CodigoEstado = http.StatusInternalServerError
		comun.MonitoreoMensajes(w, m)
		return
	}
	origin := fmt.Sprintf("http://localhost:%d/", comun.Port)
	url := fmt.Sprintf("ws://localhost:%d/ws", comun.Port)
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := ws.Write(j); err != nil {
		log.Fatal(err)
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
	voto := modelos.Voto{}

	usuario, _ = r.Context().Value("usuario").(modelos.Usuario)
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
	// recorrer el slice

	for i := range comentarios {

		db.Model(&comentarios[i]).Related(&comentarios[i].Usuarios)
		comentarios[i].Usuarios[0].Contrasenia = ""
		// buscar comentarios hijos utilizando la funcion listar hijos
		comentarios[i].Hijos = listarHijos(comentarios[i].ID)
		// se busca el voto del usuario en sesion
		voto.ComentarioID = comentarios[i].ID
		voto.UsuarioID = usuario.ID
		// almacenamos si existe registros
		count := db.Where(&voto).Find(&voto).RowsAffected
		if count > 0 {
			if voto.Valor {
				comentarios[i].CantVotos = 1
			} else {
				comentarios[i].CantVotos = -1
			}
		}
	}


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

}
func listarHijos(id uint) (hijos []modelos.Comentarios){
	db := configuracion.GetConexion()
	defer db.Close()
	db.Where("id_padre = ?", id).Find(&hijos)
	for i := range hijos {
		db.Model(&hijos[i]).Related(&hijos[i].Usuarios)
		hijos[i].Usuarios[0].Contrasenia = ""
	}
	return
}

