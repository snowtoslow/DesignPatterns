package main

import (
	"DesignPatterns/creational-design-patterns/builder"
	"log"
)

func main() {
	//Builder
	kingBuilder := builder.GetBuilder("king")
	magicianBuilder := builder.GetBuilder("magician")

	newPlayerKing := builder.NewPlayer(kingBuilder)
	newPlayerMagician := builder.NewPlayer(magicianBuilder)

	//Here we build fighters
	kingFighter := newPlayerKing.BuildFighter("Arthur", "King", "sword")
	magicianFighter := newPlayerMagician.BuildFighter("Gearalt", "Magician", "theurgy")

	log.Println("King:", kingFighter.Weapon.GetName(), kingFighter.Role, kingFighter.Name)
	log.Println("Magician:", magicianFighter.Weapon.GetName(), magicianFighter.Name, magicianFighter.Role)

	log.Println("King clone:", kingFighter.Clone())
	log.Println("Magician clone:", magicianFighter.Clone())

}
