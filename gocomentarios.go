package main

import (
	"flag"
	"github.com/golang/GoComentarios/migracion"
	"github.com/golang/GoComentarios/rutas"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"github.com/golang/GoComentarios/comun"
	"fmt"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion a la base de datos")
	flag.IntVar(&comun.Port, "port", 9000, "puerto definido")
	flag.Parse()
	if migrate == "yes" {
		log.Println("comenzo la migracion ...")
		migracion.Migrar()
		log.Println("finalizo la migracion ...")
	}
	// inicia las rutas
	router := rutas.InitRoutes()

	// inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// inicia el servidor
	servidor := &http.Server{
		Addr:    fmt.Sprintf(":%d",comun.Port),
		Handler: n,
	}
	log.Printf("iniciando el servidor en http://localhost:%d", comun.Port)
	log.Println(servidor.ListenAndServe())
	log.Println("finalizo la ejecucion del programa")
}

// ./gocomentarios --migrate yes
