package main

import (
	"log"
	"github.com/OscarMauricio666/twitor/handlers"
	"github.com/OscarMauricio666/twitor/dataBase"
)

func main() {

	if dataBase.ChequeoConnection()== 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()



}