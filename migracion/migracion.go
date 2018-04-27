package migracion

import (

	"github.com/golang/GoComentarios/configuracion"
	"github.com/golang/GoComentarios/modelos"
)

func Migrar(){
	db := configuracion.GetConexion()
	defer db.Close()

	db.CreateTable(&modelos.Usuario{})
	db.CreateTable(&modelos.Comentarios{})
	db.CreateTable(&modelos.Voto{})
	// agregando llave
	db.Model(&modelos.Voto{}).AddUniqueIndex("comentario_id_usuario_id_unique", "comentario_id", "usuario_id")


}