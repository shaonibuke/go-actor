package player

import (
	"main/actor/mail"
	"main/services/bag"
	"time"

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
