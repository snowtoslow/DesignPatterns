package models

import (
	"DesignPatterns/creational-design-patterns/factory"
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

func (m *MagicianBuilder) GetWeapon(w string) (err error) {
	m.Weapon, err = factory.GetWeapon(w)
	return
}

func (m *MagicianBuilder) SetRole(s string) {
	m.Role = s
}

func (m *MagicianBuilder) GetRole() string {
	return m.Role
}

func (m *MagicianBuilder) SetName(s string) {
	m.Name = s
}

func (m *MagicianBuilder) GetName() string {
	return m.Name
}
