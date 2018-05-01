package rutas

import (
	"github.com/golang/GoComentarios/controladores"
	"github.com/gorilla/mux"
)

// SetLoginRouter router para login
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controladores.Login).Methods("POST")
}
