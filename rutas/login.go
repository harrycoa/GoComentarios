package rutas

import (
	"github.com/gorilla/mux"
	"github.com/golang/GoComentarios/controladores"
)

// SetLoginRouter router para login
func SetLoginRouter(router *mux.Router){
	router.HandleFunc("/api/login", controladores.Login).Methods("POST")
}
