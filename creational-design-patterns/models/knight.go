package models

import (
	"DesignPatterns/creational-design-patterns/factory"
	"DesignPatterns/creational-design-patterns/prototype"
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

func (k KnightBuilder) SetWeapon(s string) {
	panic("implement me")
}

func (k KnightBuilder) GetWeapon() factory.IWeapon {
	panic("implement me")
}

func (k KnightBuilder) GetRole() string {
	panic("implement me")
}

func (k KnightBuilder) SetName(s string) {
	panic("implement me")
}

func (k KnightBuilder) SetRole(s string) {
	panic("implement me")
}

func (k KnightBuilder) GetName() string {
	panic("implement me")
}

func (k *KnightBuilder) Clone() prototype.ClonePrototyper {
	return &KnightBuilder{Name: k.Name + "_clone",
		Role:   k.Role + "_clone",
		Weapon: k.Weapon}
}
