package main

import (
	"DesignPatterns/creational-design-patterns/builder"
	"log"
)

func main() {

	kingBuilder := builder.GetBuilder("king")
	player := builder.NewPlayer(kingBuilder)

	ourShit := player.BuildFighter("Arthur", "King", "sword")

	log.Println(ourShit.Weapon, ourShit.Role, ourShit.Name)
}
