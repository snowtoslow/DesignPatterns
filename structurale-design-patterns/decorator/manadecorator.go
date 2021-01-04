package decorator

import "DesignPatterns/creational-design-patterns/factory"

type ManaDecorator struct {
	Weapon factory.IWeapon
}

func (manaDecorator *ManaDecorator) DecorateWeapon() {
	manaDecorator.Weapon.SetPower(manaDecorator.Weapon.GetPower() + 10)
	manaDecorator.Weapon.SetName(manaDecorator.Weapon.GetName() + " decorated with mana")
}
