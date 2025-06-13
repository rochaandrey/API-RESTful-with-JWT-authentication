package main

import (
	"gin/database"
	"gin/handler"
	"gin/repository"
	"gin/routers"
	"log"
)

func main() {
	database.RunMigrations()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Nao foi possivel conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	clienteRepo := repository.NewClienteRepository(db)
	h := handler.NewHandler(clienteRepo)

	routers.Initialize(h)
}
