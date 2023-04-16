package player

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/shaonibuke/go-actor/actor/mail"
)

func (p *Player) getPlayerName(m *mail.Mail) {
	//p.actor.SendMessage("bag", "", "GetPlayerName", p.Name)
	time.Sleep(time.Second * 1)
	p.count += 1
	log.Debugf("getPlayerName %s %d", m.ReplyID, p.count)
	p.actor.ReplyMessage(m.FormServiceType, m.FormID, m.ReplyID, p.Name)
}
