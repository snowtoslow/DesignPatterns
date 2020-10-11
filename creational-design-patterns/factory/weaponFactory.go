package factory

import (
	"fmt"
)

func GetWeapon(weaponType string) (IWeapon, error) {
	if weaponType == "sword" {
		return NewSword(), nil
	}

	if weaponType == "spear" {
		return NewSpear(), nil
	}

	if weaponType == "theurgy" {
		return NewTheurgy(), nil
	}

	if weaponType == "dagger" {
		return NewDagger(), nil
	}

	if weaponType == "bowandarrow" {
		return NewBowAndArrow(), nil
	}

	return nil, fmt.Errorf("wrong weapon type")
}
