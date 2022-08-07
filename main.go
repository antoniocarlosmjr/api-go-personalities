package main

import (
	"fmt"
	"github.com/antoniocarlosmjr/api-go-personalities/database"
	"github.com/antoniocarlosmjr/api-go-personalities/models"
	"github.com/antoniocarlosmjr/api-go-personalities/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Tobias Barreto", Story: "He was a philosopher, poet, critic, etc..."},
		{Id: 2, Name: "Silvio Romero", Story: "He was a lawyer, literature critic, journalist, etc..."},
	}

	database.ConnectDatabase()
	fmt.Println("Initing server....")
	routes.HandleRequest()
}
