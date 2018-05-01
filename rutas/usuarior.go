package rutas

import (
	"github.com/gorilla/mux"
	"github.com/golang/GoComentarios/controladores"
	"github.com/urfave/negroni"
)
// SetUserRouter ruta para el registro de usuario
func SetUserRouter(router *mux.Router){
	prefix := "/api/usuarior"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controladores.CrearUsuario).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			// negroni.HandlerFunc(controladores.ValidarToken)
			negroni.Wrap(subRouter),
		),
	)
}
