package player

import "main/actor"

var player *Player

func init() {
	player = &Player{}
	player.Name = "player"
	player.actor = actor.GetActorManager().NewActor("player")
	player.actor.RegisterRouter("GetPlayerName", player.GetPlayerName)
}

func GetPlayer() *Player {
	return player
}
