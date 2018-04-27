package configuracion

import (
	"os"
	"log"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

// Estructura que mapeara la funcion de nuestro archivo json
type Configuracion struct {
	Server string
	Port string
	User string
	Password string
	Database string
}

func GetConfiguracion() Configuracion{
	var c Configuracion
	// abrir archivo json
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// paquete json, que devuelve erro
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func GetConexion() *gorm.DB {
	c := GetConfiguracion()
	// url de conexion
	// user:password@tcp(server:port)/database?charset=utf8&parseTime=True&loc=local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User,c.Password,c.Server,c.Port,c.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}