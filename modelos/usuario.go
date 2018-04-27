package modelos

import (
	"github.com/jinzhu/gorm"
	//	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Usuario user del sistema
type Usuario struct {
	gorm.Model
	NombreUsuario       string        `json:"nombreUsuario" gorm:"not null; unique"`
	Email               string        `json:"email" gorm:"not null; unique"`
	NombreCompleto      string        `json:"nombreCompleto" gorm:"not null"`
	Contrasenia         string        `json:"contrasenia, omitempty" gorm:"not null; type:varchar(256)"` // omitempty para que la contraseña no quede en blanco
	ConfirmaContrasenia string        `json:"confirmaContrasenia,omitempty" gorm:"-"`                    // el guion es para que el orm omita
	Imagen              string        `json:"imagen"`
	Comentarios         []Comentarios `json:"comentarios,omitempty"`
}
