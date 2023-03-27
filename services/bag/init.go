package bag

import "main/actor"

var bag *Bag

func init() {
	bag = &Bag{}
	bag.actor = actor.GetActorManager().NewActor("bag")
	bag.actor.RegisterRouter("AddItem", bag.AddItem)
	bag.actor.RegisterRouter("GetItem", bag.GetItem)
}
