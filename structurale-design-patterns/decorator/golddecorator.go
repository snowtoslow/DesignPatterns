package decorator

import "DesignPatterns/creational-design-patterns/factory"

type GoldDecorator struct {
	Weapon factory.IWeapon
}

func (goldDecorator *GoldDecorator) DecorateWeapon() int {
	return goldDecorator.Weapon.GetPower() + 15
}
