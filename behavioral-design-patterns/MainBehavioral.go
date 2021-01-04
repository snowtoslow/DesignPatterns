package main

import (
	"DesignPatterns/behavioral-design-patterns/strategy"
	"log"
)

func main() {

	v := "sword"

	if commandStrategy, exists := strategy.Strategies[v]; exists {
		//to demonstrate strategy by an input of sword
		log.Println(commandStrategy.CreateWeapon())
	}

}
