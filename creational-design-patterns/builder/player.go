package builder

import (
	"DesignPatterns/creational-design-patterns/models"
	"log"
)

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

func (p *Player) BuildFighter(name string, role string, weaponName string) models.Fighter {
	p.builder.SetName(name)
	p.builder.SetRole(role)
	if err := p.builder.GetWeapon(weaponName); err != nil {
		log.Println(err)
	}
	return p.builder.GetFighter()
}
