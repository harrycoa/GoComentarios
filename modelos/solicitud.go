package modelos

import jwt "github.com/dgrijalva/jwt-go"

// Solicitud token de usuario
type Solicitud struct {
	Usuario            `json:"usuario"`
	jwt.StandardClaims // campos estandar
}
