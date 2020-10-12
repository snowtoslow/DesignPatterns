package models

import (
	"DesignPatterns/creational-design-patterns/factory"
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

func (k *King) GetWeapon(w string) (err error) {
	k.Weapon, err = factory.GetWeapon(w)
	return nil
}

func (k *King) GetRole() string {
	return k.Role
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
