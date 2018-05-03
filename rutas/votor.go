package rutas

import (
	"github.com/gorilla/mux"
	"github.com/golang/GoComentarios/controladores"
	"github.com/urfave/negroni"
)

// SetVoteRouter es la ruta para el registro de un voto
func SetVoteRouter(router *mux.Router) {
	prefix := "/api/votes"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controladores.RegistroVoto).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controladores.ValidarToken),
			negroni.Wrap(subRouter),
		),
	)
}