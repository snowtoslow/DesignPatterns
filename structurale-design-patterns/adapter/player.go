package adapter

type Player struct {
}

func (p *Player) ChangeFighterRole(changer RoleChanger, newRole string) {
	changer.ChangeFighterRole(newRole)
}
