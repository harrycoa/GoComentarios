package modelos

import jwt "github.com/dgrijalva/jwt-go"

type Solicitud struct {
	Usuario            `json:"usuario"`
	jwt.StandardClaims // campos estandar
}
