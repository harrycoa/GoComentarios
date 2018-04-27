package main

import (
	"flag"
	"github.com/golang/GoComentarios/migracion"
	"log"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion a la base de datos")
	flag.Parse()
	if migrate == "yes" {
		log.Println("comenzo la migracion ...")
		migracion.Migrar()
		log.Println("finalizo la migracion ...")
	}
}

// ./gocomentarios --migrate yes
