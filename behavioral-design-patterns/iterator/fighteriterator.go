package iterator

import "DesignPatterns/creational-design-patterns/models"

type fighterIterator struct {
	index    int
	fighters []*models.Fighter
}

func (f *fighterIterator) HasNext() bool {
	if f.index < len(f.fighters) {
		return true
	}
	return false
}

func (f *fighterIterator) GetNext() *models.Fighter {
	if f.HasNext() {
		fighter := f.fighters[f.index]
		f.index++
		return fighter
	}
	return nil
}
