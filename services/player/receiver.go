package player

import (
	"fmt"

	"github.com/shaonibuke/go-actor/actor/mail"
)

func (p *Player) GetPlayerName(m *mail.Mail) {
	//p.actor.SendMessage("bag", "", "GetPlayerName", p.Name)

	p.count += 1
	fmt.Printf("GetPlayerName %s %d\n", m.ReplyID, p.count)
	p.actor.ReplyMessage(m.FormServiceType, m.FormID, m.ReplyID, p.Name)
}
