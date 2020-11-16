package models

import (
	"DesignPatterns/creational-design-patterns/factory"
)

type QueenBuilder struct {
	Role     string
	Weapon   factory.IWeapon
	Name     string
	King     King
	Magician []MagicianBuilder
}

func NewQueenBuilder() *QueenBuilder {
	return &QueenBuilder{}
}

func (q *QueenBuilder) GetFighter() Fighter {
	panic("implement me")
}

func (q *QueenBuilder) GetWeapon(w string) (err error) {
	q.Weapon, err = factory.GetWeaponFactorySingleInstance().GetWeapon(w)
	return
}

func (q *QueenBuilder) SetRole(s string) {
	q.Role = s
}

func (q *QueenBuilder) GetRole() string {
	return q.Role
}

func (q *QueenBuilder) SetName(s string) {
	q.Name = s
}

func (q *QueenBuilder) GetName() string {
	return q.Name
}

func (q *QueenBuilder) SetWeapon(weapon factory.IWeapon) {
	q.Weapon = weapon
}
