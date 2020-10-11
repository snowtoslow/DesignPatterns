package models

import "DesignPatterns/creational-design-patterns/factory"

type Fighter struct {
	Role   string
	Weapon factory.IWeapon
	Name   string
}
