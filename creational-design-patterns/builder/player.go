package builder

import (
	"DesignPatterns/creational-design-patterns/factory"
	"DesignPatterns/creational-design-patterns/models"
	"DesignPatterns/structurale-design-patterns/adapter"
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

func (p *Player) ChangeFighterRole(changer adapter.RoleChanger, newRole string) {
	changer.ChangeFighterRole(newRole)
}

func (p *Player) SetBuilder(builder IBuilder) {
	p.builder = builder
}

func (p *Player) BuildFighter(name string, role string, weaponName string) models.Fighter {
	p.builder.SetName(name)
	p.builder.SetRole(role)

	if weapon, err := factory.GetWeaponFactorySingleInstance().GetWeapon(weaponName); err != nil {
		log.Println(err)
	} else {
		p.builder.SetWeapon(weapon)
	}

	return p.builder.GetFighter()
}
