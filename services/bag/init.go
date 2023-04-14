package bag

import "github.com/shaonibuke/go-actor/actor"

var bag *Bag

func init() {
	bag = &Bag{}
	bag.actor = actor.GetActorManager().NewActor("bag")
	bag.actor.RegisterRouter("addItem", bag.addItem)
	bag.actor.RegisterRouter("getItem", bag.getItem)
}
