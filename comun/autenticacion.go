package comun

import (
	"crypto/rsa"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/golang/GoComentarios/modelos"
	"io/ioutil"
	"log"
)

var (
	privateKey *rsa.PrivateKey
	// PublicKey se usa para validar el token
	PublicKey *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./llaves/private.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo privado")
	}

	publicBytes, err := ioutil.ReadFile("./llaves/public.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo public")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a privatekey")
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a publicKey")
	}
}

// GenerateJWT Genera el token para el cliente
func GenerateJWT(usuario modelos.Usuario) string {
	claims := modelos.Solicitud{
		Usuario: usuario,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt:time.Now().Add(time.Hour * 2).Unix(),
			Issuer: " Comentarios",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("No se pudo firmar el token")
	}
	return result
}
