package modelos

import "github.com/jinzhu/gorm"

// Voto permite controlar que un usuario solo vote una  unica vez por cada comentario
type Voto struct {
	gorm.Model
	ComentarioID uint `json:"comentarioID" gorm:"not null"`
	UsuarioID    uint `json:"usuarioID" gorm:"not null"`
	Valor        bool `json:"valor" gorm:"not null"`
}
