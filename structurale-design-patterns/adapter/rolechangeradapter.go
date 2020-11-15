package adapter

import (
	"DesignPatterns/creational-design-patterns/models"
)

type RoleChangerAdapter struct {
	Fighter *models.Fighter
}

func (w *RoleChangerAdapter) ChangeFighterRole(newRoleName string) {
	w.Fighter.ChangeFighterRole(newRoleName)
}
