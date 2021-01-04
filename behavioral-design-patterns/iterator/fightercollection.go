package iterator

import "DesignPatterns/creational-design-patterns/models"

type FighterCollection struct {
	Fighters []*models.Fighter
}

func (u *FighterCollection) CreateIterator() iterator {
	return &fighterIterator{
		fighters: u.Fighters,
	}
}
