package decorator

import "DesignPatterns/creational-design-patterns/factory"

type FenixFeathersDecorator struct {
	Weapon factory.Weapon
}

func (fenixFeathersDecorator *FenixFeathersDecorator) DecorateWeapon() {
	fenixFeathersDecorator.Weapon.SetPower(fenixFeathersDecorator.Weapon.GetPower() + 20)
	fenixFeathersDecorator.Weapon.SetName(fenixFeathersDecorator.Weapon.GetName() + " decorated with fenix feather")
}
