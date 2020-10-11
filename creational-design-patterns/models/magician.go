package models

import (
	"DesignPatterns/creational-design-patterns/factory"
	"DesignPatterns/creational-design-patterns/prototype"
)

type MagicianBuilder struct {
	Role   string
	Weapon factory.IWeapon
	Name   string
}

func NewMagicianBuilder() *MagicianBuilder {
	return &MagicianBuilder{}
}

func (m MagicianBuilder) GetFighter() Fighter {
	return Fighter{
		Name:   m.Name,
		Role:   m.Role,
		Weapon: m.Weapon,
	}
}

func (m MagicianBuilder) SetWeapon(s string) {
	m.Weapon.SetName(s)
}

func (m MagicianBuilder) GetWeapon() factory.IWeapon {
	return m.Weapon
}

func (m MagicianBuilder) SetRole(s string) {
	m.Role = s
}

func (m MagicianBuilder) GetRole() string {
	return m.Role
}

func (m MagicianBuilder) SetName(s string) {
	m.Name = s
}

func (m MagicianBuilder) GetName() string {
	return m.Name
}

func (m *MagicianBuilder) Clone() prototype.ClonePrototyper {
	return &MagicianBuilder{Name: m.Name + "_clone",
		Role:   m.Role + "_clone",
		Weapon: m.Weapon}
}
