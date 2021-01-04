package decorator

import "DesignPatterns/creational-design-patterns/factory"

type DragonStealDecorator struct {
	Weapon factory.IWeapon
}

func (dragonStealDecorator *DragonStealDecorator) DecorateWeapon() {
	dragonStealDecorator.Weapon.SetPower(dragonStealDecorator.Weapon.GetPower() + 30)
	dragonStealDecorator.Weapon.SetName(dragonStealDecorator.Weapon.GetName() + " decorated with dragon steal")

}
