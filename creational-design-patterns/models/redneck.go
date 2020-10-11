package models

import (
	"DesignPatterns/creational-design-patterns/factory"
	"DesignPatterns/creational-design-patterns/prototype"
)

type RedNeckBuilder struct {
	Role   string
	Weapon factory.IWeapon
	Name   string
}

func NewRedNeckBuilder() *RedNeckBuilder {
	return &RedNeckBuilder{}
}

func (r RedNeckBuilder) GetFighter() Fighter {
	return Fighter{
		Name:   r.Name,
		Role:   r.Role,
		Weapon: r.Weapon,
	}
}

func (r RedNeckBuilder) SetWeapon(s string) {
	r.Weapon.SetName(s)
}

func (r RedNeckBuilder) GetWeapon() factory.IWeapon {
	return r.Weapon
}

func (r RedNeckBuilder) GetRole() string {
	return r.Role
}

func (r RedNeckBuilder) SetRole(s string) {
	r.Role = s
}

func (r RedNeckBuilder) SetName(s string) {
	r.Name = s
}

func (r RedNeckBuilder) GetName() string {
	return r.Name
}

func (r *RedNeckBuilder) Clone() prototype.ClonePrototyper {
	return &RedNeckBuilder{Name: r.Name + "_clone",
		Role:   r.Role + "_clone",
		Weapon: r.Weapon}
}
