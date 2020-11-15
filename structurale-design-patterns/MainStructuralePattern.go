package main

import (
	"DesignPatterns/creational-design-patterns/builder"
	"DesignPatterns/structurale-design-patterns/adapter"
	"DesignPatterns/structurale-design-patterns/decorator"
	"log"
)

func main() {

	kingBuilder := builder.GetBuilder("king")
	newPlayerKing := builder.NewPlayer(kingBuilder)
	fighter := newPlayerKing.BuildFighter("Arthur", "King", "sword")
	log.Printf("King weapon power before decorate: %d", fighter.Weapon.GetPower())

	// DECORATOR EXAMPLE!
	kingDragonStealSword := &decorator.DragonStealDecorator{
		Weapon: fighter.Weapon,
	}
	kingDragonStealSword.DecorateWeapon()

	kingWithStealAndManaAdd := &decorator.ManaDecorator{
		Weapon: kingDragonStealSword.Weapon,
	}

	kingWithStealAndManaAdd.DecorateWeapon()

	log.Printf("King decorated with sword made from dragon steal and mana: %d", fighter.Weapon.GetPower())

	//ADAPTER EXAMPLE:

	log.Printf("Role of fighter before changes: %s", fighter.Role)

	fighterAdapter := adapter.RoleChangerAdapter{
		Fighter: &fighter,
	}

	newPlayerKing.ChangeFighterRole(&fighterAdapter, "Queen")

	log.Printf("Role of fighter after changes: %s", fighter)

}
