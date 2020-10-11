package models

import (
	"DesignPatterns/creational-design-patterns/factory"
	"DesignPatterns/creational-design-patterns/prototype"
)

type King struct {
	Role   string
	Weapon factory.IWeapon
	Name   string
}

func (k *King) GetFighter() Fighter {
	return Fighter{
		Name:   k.Name,
		Role:   k.Role,
		Weapon: k.Weapon,
	}
}

func (k *King) GetWeapon() factory.IWeapon {
	return k.Weapon
}

func (k *King) GetRole() string {
	return k.Role
}

func (k *King) SetWeapon(s string) {
	k.Weapon.SetName(s)
}

func (k *King) SetName(s string) {
	k.Name = s
}

func (k *King) SetRole(s string) {
	k.Role = s
}

func (k *King) GetName() string {
	return k.Name
}

func (k *King) Clone() prototype.ClonePrototyper {
	return &King{Name: k.Name + "_clone",
		Role:   k.Role + "_clone",
		Weapon: k.Weapon}
}
