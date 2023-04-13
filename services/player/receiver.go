package player

import (
	"github.com/charmbracelet/log"
	"github.com/shaonibuke/go-actor/actor/mail"
)

func (p *Player) GetPlayerName(m *mail.Mail) {
	//p.actor.SendMessage("bag", "", "GetPlayerName", p.Name)

	p.count += 1
	log.Debugf("GetPlayerName %s %d\n", m.ReplyID, p.count)
	p.actor.ReplyMessage(m.FormServiceType, m.FormID, m.ReplyID, p.Name)
}
