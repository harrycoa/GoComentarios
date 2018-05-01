package rutas

import (
	//"github.com/gorilla/mux"
	//"github.com/urfave/negroni"
	// "github.com/gola"

	"github.com/gorilla/mux"
)
func SetUserRouter(router *mux.Router){
	prefix := "/api/users"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", Control)
}
