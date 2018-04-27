package modelos

import "github.com/jinzhu/gorm"

type Voto struct {
	gorm.Model
	ComentarioID uint `json:"comentarioID" gorm:"not null"`
	UsuarioID    uint `json:"usuarioID" gorm:"not null"`
	Valor        bool `json:"valor" gorm:"not null"`
}
