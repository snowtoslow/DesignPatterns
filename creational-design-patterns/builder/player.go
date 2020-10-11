package builder

import "DesignPatterns/creational-design-patterns/models"

type Player struct {
	builder IBuilder
}

func NewPlayer(builder IBuilder) *Player {
	return &Player{
		builder: builder,
	}
}

func (p *Player) SetBuilder(builder IBuilder) {
	p.builder = builder
}

func (p *Player) BuildFighter(name string, role string, weapon string) models.Fighter {
	p.builder.SetName(name)
	p.builder.SetRole(role)
	p.builder.SetWeapon(weapon)
	return p.builder.GetFighter()
}
