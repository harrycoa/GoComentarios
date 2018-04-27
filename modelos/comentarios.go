package modelos

import "github.com/jinzhu/gorm"

// Comentarios del sistema
type Comentarios struct {
	gorm.Model
	IDUsuario uint   `json:"idUsuario"`
	IDPadre   uint   `json:"idPadre"`
	Votos     int32  `json:"votos"`
	Contenido string `json:"contenido"`
	CantVotos int8   `json:"cantVotos" gorm:"-"`
	// hacer un hack para que el usuario pueda consultar
	Usuarios []Usuario     `json:"usuarios, omitempty"` // informacion del usuario que creo el comentario
	Hijos    []Comentarios `json:"hijos,omitempty"`     // comentarios respuesta
}
