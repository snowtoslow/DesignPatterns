package builder

import (
	"DesignPatterns/creational-design-patterns/factory"
	"DesignPatterns/creational-design-patterns/models"
)

type IBuilder interface {
	SetName(string)
	SetRole(string)
	GetRole() string
	GetName() string
	SetWeapon(string)
	GetWeapon() factory.IWeapon
	GetFighter() models.Fighter
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "king" {
		return &models.King{}
	}

	if builderType == "queen" {
		return &models.QueenBuilder{}
	}

	if builderType == "magician" {
		return &models.MagicianBuilder{}
	}

	if builderType == "redneck" {
		return &models.RedNeckBuilder{}
	}

	if builderType == "knight" {
		return &models.KnightBuilder{}
	}

	return nil
}
