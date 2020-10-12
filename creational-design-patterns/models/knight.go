package models

import (
	"DesignPatterns/creational-design-patterns/factory"
)

type KnightBuilder struct {
	Role   string
	Weapon factory.IWeapon
	Name   string
}

func NewKnightBuilder() *KnightBuilder {
	return &KnightBuilder{}
}

func (k KnightBuilder) GetFighter() Fighter {
	return Fighter{
		Name:   k.Name,
		Role:   k.Role,
		Weapon: k.Weapon,
	}
}
func (k *KnightBuilder) GetWeapon(w string) (err error) {
	k.Weapon, err = factory.GetWeapon(w)
	return
}

func (k KnightBuilder) GetRole() string {
	return k.Role
}

func (k KnightBuilder) SetName(s string) {
	k.Name = s
}

func (k KnightBuilder) SetRole(s string) {
	k.Role = s
}

func (k KnightBuilder) GetName() string {
	return k.GetName()
}
