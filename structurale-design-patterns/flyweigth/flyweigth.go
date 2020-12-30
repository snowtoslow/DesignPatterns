package flyweigth

import "DesignPatterns/creational-design-patterns/factory"

var weaponFactorySingleInstance = &WeaponFactory{
	WeaponFactory: make(map[string]factory.IWeapon),
}

//flyweigth
type WeaponFactory struct {
	WeaponFactory map[string]factory.IWeapon
}

func GetWeaponFactorySingleInstance() *WeaponFactory {
	return weaponFactorySingleInstance
}
