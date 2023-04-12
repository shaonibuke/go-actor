package player

import (
	"time"

	"github.com/shaonibuke/go-actor/actor/mail"
	"github.com/shaonibuke/go-actor/services/bag"

	"github.com/gookit/goutil/dump"
)

func (p *Player) backGetItem(m *mail.Mail) {
	time.Sleep(time.Second * 1)
	if m.Msg.(bag.Item).ID == "" {
		// 没找到
		dump.P("not find items")
	} else {
		// 找到了
		dump.P("find items", m.Msg)
	}
}
