package main

import (
	"DesignPatterns/behavioral-design-patterns/strategy"
)

func main() {

	v := "king"

	if commandStrategy, exists := strategy.Strategies[v]; exists {
		commandStrategy.CreateWeapon()
	}

}
