package factory

import (
	"fmt"
)

//flywiegth
var weaponFactorySingleInstance = &WeaponFactory{
	WeaponFactory: make(map[string]IWeapon),
}

//flywiegth
type WeaponFactory struct {
	WeaponFactory map[string]IWeapon
}

func (wf *WeaponFactory) GetWeapon(weaponType string) (IWeapon, error) {
	if weaponType == "sword" {
		wf.WeaponFactory[weaponType] = NewSword()
		return NewSword(), nil
	}

	if weaponType == "spear" {
		wf.WeaponFactory[weaponType] = NewSpear()
		return NewSpear(), nil
	}

	if weaponType == "theurgy" {
		wf.WeaponFactory[weaponType] = NewTheurgy()
		return NewTheurgy(), nil
	}

	if weaponType == "dagger" {
		wf.WeaponFactory[weaponType] = NewDagger()
		return NewDagger(), nil
	}

	if weaponType == "bowandarrow" {
		wf.WeaponFactory[weaponType] = NewBowAndArrow()
		return NewBowAndArrow(), nil
	}

	return nil, fmt.Errorf("wrong weapon type")
}

// this is for flyweigth
func GetWeaponFactorySingleInstance() *WeaponFactory {
	return weaponFactorySingleInstance
}
