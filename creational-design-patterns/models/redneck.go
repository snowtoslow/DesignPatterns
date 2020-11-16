package models

import (
	"DesignPatterns/creational-design-patterns/factory"
)

type RedNeckBuilder struct {
	Role   string
	Weapon factory.IWeapon
	Name   string
}

func NewRedNeckBuilder() *RedNeckBuilder {
	return &RedNeckBuilder{}
}

func (r *RedNeckBuilder) GetFighter() Fighter {
	return Fighter{
		Name:   r.Name,
		Role:   r.Role,
		Weapon: r.Weapon,
	}
}

func (r *RedNeckBuilder) GetWeapon(w string) (err error) {
	r.Weapon, err = factory.GetWeaponFactorySingleInstance().GetWeapon(w)
	return
}

func (r *RedNeckBuilder) GetRole() string {
	return r.Role
}

func (r *RedNeckBuilder) SetRole(s string) {
	r.Role = s
}

func (r *RedNeckBuilder) SetName(s string) {
	r.Name = s
}

func (r *RedNeckBuilder) GetName() string {
	return r.Name
}

func (r *RedNeckBuilder) SetWeapon(weapon factory.IWeapon) {
	r.Weapon = weapon
}
