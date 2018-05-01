package controladores

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/GoComentarios/comun"
	"github.com/golang/GoComentarios/configuracion"
	"github.com/golang/GoComentarios/modelos"
	"log"
	"net/http"
)

// Login es el controlador de login
func Login(w http.ResponseWriter, r *http.Request) {
	usuario := modelos.Usuario{}
	// mapear los datos
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		fmt.Fprintf(w, "error: %s\n", err)
		return
	}
	// conexion a la base de datos
	db := configuracion.GetConexion()
	defer db.Close()

	// encriptando contraseña con sha 256
	c := sha256.Sum256([]byte(usuario.Contrasenia))
	// otra forma
	pwd := base64.URLEncoding.EncodeToString(c[:32])
	// pwd := fmt.Sprintf("%x", c)

	// mapeando el resultado
	db.Where("email = ? and contrasenia = ?", usuario.Email, pwd).First(&usuario)
	if usuario.ID > 0 {
		usuario.Contrasenia = ""
		token := comun.GenerateJWT(usuario)
		j, err := json.Marshal(modelos.Token{Token: token})
		if err != nil {
			log.Fatalf("error al convertir el token a json: %s", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		// enviaremos a traves de un json si no esta autorizado
		m := modelos.Mensaje{
			Mensaje:      "usuario o clave no valido",
			CodigoEstado: http.StatusUnauthorized,
		}
		comun.MonitoreoMensajes(w, m)
	}
}

// CrearUsuario permite registrar usuarios
func CrearUsuario(w http.ResponseWriter, r *http.Request) {
	// al estar sin parametros se incrusta con sus valores por defecto
	usuario := modelos.Usuario{}
	m := modelos.Mensaje{}
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		m.Mensaje = fmt.Sprintf("Error al leer el usuario a registrar %s", err)
		m.CodigoEstado = http.StatusBadRequest
		comun.MonitoreoMensajes(w, m)
		return
	}
	if usuario.Contrasenia != usuario.ConfirmaContrasenia {
		m.Mensaje = "las contraseñas no coinciden"
		m.CodigoEstado = http.StatusBadRequest
		comun.MonitoreoMensajes(w, m)
		return
	}

	c := sha256.Sum256([]byte(usuario.Contrasenia))
	// otra forma pwd := base64.URLEncoding.EncodeToString(c[:32])
	pwd := fmt.Sprintf("%x", c)
	usuario.Contrasenia = pwd

	// codificando en md5 el email
	picmd5 := md5.Sum([]byte(usuario.Email))
	picstr := fmt.Sprintf("%x", picmd5)
	usuario.Imagen = "https://gravatar.com/avatar/" + picstr + "?s=100"

	// crear una conexion para guardar el usuario
	db := configuracion.GetConexion()
	defer db.Close()

	err = db.Create(&usuario).Error
	if err != nil {
		m.Mensaje = fmt.Sprintf("Error al crear el registro %s", err)
		m.CodigoEstado = http.StatusBadRequest
		comun.MonitoreoMensajes(w, m)
		return
	}
	m.Mensaje = "Usuario creado con exito"
	m.CodigoEstado = http.StatusCreated
	comun.MonitoreoMensajes(w, m)

}
