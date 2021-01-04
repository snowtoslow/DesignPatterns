package iterator

import "DesignPatterns/creational-design-patterns/models"

type iterator interface {
	HasNext() bool
	GetNext() *models.Fighter
}
