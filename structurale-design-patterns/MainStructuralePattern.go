package main

import (
	"DesignPatterns/behavioral-design-patterns/memento"
	"DesignPatterns/creational-design-patterns/builder"
	"DesignPatterns/creational-design-patterns/factory"
	"DesignPatterns/structurale-design-patterns/adapter"
	"DesignPatterns/structurale-design-patterns/decorator"
	"log"
)

func main() {

	kingBuilder := builder.GetBuilder("king")
	newPlayerKing := builder.NewPlayer(kingBuilder)
	fighter := newPlayerKing.BuildFighter("Arthur", "King", "sword")
	log.Printf("King weapon power before decorate: %d", fighter.Weapon.GetPower())

	caretaker := &memento.CareTaker{
		MementoArray: make([]*memento.Memento, 0),
	}

	originator := &memento.Originator{
		State: fighter.Weapon.GetName(),
	}

	caretaker.AddMemento(originator.CreateMemento())

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

	weaponFactoryInstance := factory.GetWeaponFactorySingleInstance()
	log.Println(weaponFactoryInstance.WeaponFactory)
	/*for k,v := range weaponFactoryInstance.WeaponFactory{
		log.Println(k,v)
	}*/

	//here is the memento!
	log.Println(caretaker.GetMenento(0))

}
