package player

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/shaonibuke/go-actor/actor/mail"
	"github.com/shaonibuke/go-actor/services/bag"
)

func (p *Player) backGetItem(m *mail.Mail) {
	time.Sleep(time.Second * 1)
	if m.Msg.(bag.Item).ID == "" {
		// 没找到
		log.Debug("not find items")
	} else {
		log.Debug("find items")
	}
}
