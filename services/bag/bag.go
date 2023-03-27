package bag

import (
	"main/actor"
)

type Item struct {
	ID    string
	Price int
}

type Bag struct {
	items []Item
	actor *actor.Actor
}
