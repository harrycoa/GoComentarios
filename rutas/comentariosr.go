package rutas

import (
	"github.com/gorilla/mux"
	"github.com/golang/GoComentarios/controladores"
	"github.com/urfave/negroni"
)

func SetCommentRouter(router *mux.Router){
	prefix := "/api/comments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controladores.CrearComentario).Methods("POST")
	subRouter.HandleFunc("/", controladores.ListarComentarios).Methods("GET")
	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controladores.ValidarToken),
			negroni.Wrap(subRouter),
		),
	)
}