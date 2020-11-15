package models

import (
	"DesignPatterns/creational-design-patterns/factory"
	"DesignPatterns/creational-design-patterns/prototype"
)

type Fighter struct {
	Role   string
	Weapon factory.IWeapon
	Name   string
}

func (f *Fighter) ChangeFighterRole(newRoleName string) {
	f.Role = newRoleName
}

func (f *Fighter) Clone() prototype.ClonePrototyper {
	return &Fighter{
		Name:   f.Name + "_clone",
		Role:   f.Role + "_clone",
		Weapon: f.Weapon,
	}
}
