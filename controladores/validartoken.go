package controladores

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/golang/GoComentarios/comun"
	"github.com/golang/GoComentarios/modelos"
	"net/http"
	"context"
)

// ValidarToken validar el token del cliente
func ValidarToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var m modelos.Mensaje
	token, err := request.ParseFromRequestWithClaims(
		r,
		request.OAuth2Extractor,
		&modelos.Solicitud{},
		func(t *jwt.Token) (interface{}, error) {
			return comun.PublicKey, nil
		},
	)
	if err != nil {
		m.CodigoEstado = http.StatusUnauthorized
		switch err.(type) {
		case *jwt.ValidationError:
			vError := err.(*jwt.ValidationError)
			switch vError.Errors {
			case jwt.ValidationErrorExpired:
				m.Mensaje = "su token a expirado"
				comun.MonitoreoMensajes(w, m)
				return
			case jwt.ValidationErrorSignatureInvalid:
				m.Mensaje = "LA firma del token no coincide"
				comun.MonitoreoMensajes(w, m)
				return
			default:
				m.Mensaje = "su token no es valido"
				comun.MonitoreoMensajes(w, m)
				return
			}
		}
	}
	if token.Valid {
		// crear un contexto
		ctx := context.WithValue(r.Context(), "usuario", token.Claims.(*modelos.Solicitud).Usuario)
		next(w, r.WithContext(ctx))
	} else {
		m.CodigoEstado = http.StatusUnauthorized
		m.Mensaje = "su token no es valido"
		comun.MonitoreoMensajes(w, m)
	}
}
