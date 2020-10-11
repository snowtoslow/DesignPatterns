package models

import (
	"DesignPatterns/creational-design-patterns/factory"
	"DesignPatterns/creational-design-patterns/prototype"
)

type QueenBuilder struct {
	Role   string
	Weapon factory.IWeapon
	Name   string
}

func NewQueenBuilder() *QueenBuilder {
	return &QueenBuilder{}
}

func (q *QueenBuilder) GetFighter() Fighter {
	panic("implement me")
}

func (q *QueenBuilder) SetWeapon(s string) {
	q.Weapon.SetName(s)
}

func (q *QueenBuilder) GetWeapon() factory.IWeapon {
	return q.Weapon
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

func (q *QueenBuilder) Clone() prototype.ClonePrototyper {
	return &QueenBuilder{Name: q.Name + "_clone",
		Role:   q.Role + "_clone",
		Weapon: q.Weapon}
}
