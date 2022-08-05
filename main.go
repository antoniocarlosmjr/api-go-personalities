package main

import (
	"fmt"
	"github.com/antoniocarlosmjr/api-go-personalities/models"
	"github.com/antoniocarlosmjr/api-go-personalities/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Tobias Barreto", Story: "He was a philosopher, poet, critic, etc..."},
		{Name: "Silvio Romero", Story: "He was a lawyer, literature critic, journalist, etc..."},
	}
	fmt.Println("Initing server....")
	routes.HandleRequest()
}
