package main

import (
	"gin/database"
	"gin/handler"
	"gin/routers"
	"log"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Nao foi possivel conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	h := handler.NewHandler(db)

	routers.Initialize(h)
}
