package decorator

import "DesignPatterns/creational-design-patterns/factory"

type GoldDecorator struct {
	Weapon factory.IWeapon
}

func (goldDecorator *GoldDecorator) DecorateWeapon() {
	goldDecorator.Weapon.SetPower(goldDecorator.Weapon.GetPower() + 15)
	goldDecorator.Weapon.SetName(goldDecorator.Weapon.GetName() + " decorated with gold")
}
