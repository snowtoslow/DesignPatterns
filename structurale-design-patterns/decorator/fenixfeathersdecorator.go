package decorator

import "DesignPatterns/creational-design-patterns/factory"

type FenixFeathersDecorator struct {
	Weapon factory.Weapon
}

func (fenixFeathersDecorator *FenixFeathersDecorator) DecorateWeapon() int {
	return fenixFeathersDecorator.Weapon.GetPower() + 20
}
