package bag

import (
	"github.com/shaonibuke/go-actor/actor"
)

type Item struct {
	ID    string
	Price int
}

type Bag struct {
	items []Item
	actor *actor.Actor
}
