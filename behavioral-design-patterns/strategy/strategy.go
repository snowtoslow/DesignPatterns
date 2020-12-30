package strategy

import (
	"DesignPatterns/creational-design-patterns/factory"
)

type Strategy interface {
	CreateWeapon() factory.IWeapon
}

var Strategies = map[string]Strategy{
	"sword":       factory.Sword{},
	"spear":       factory.Spear{},
	"theurgy":     factory.Dagger{},
	"dagger":      factory.Dagger{},
	"bowandarrow": factory.BowAndArrow{},
}
