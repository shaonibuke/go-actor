package mail

type MsgType int

const (
	MsgTypeTo    MsgType = iota // 消息发送给某个Actor
	MsgTypeReply                // 消息回复给某个Actor
)

type Mail struct {
	Msg             interface{} // 消息
	MsgName         string      // 消息名 用于路由
	MsgType         MsgType     // 消息类型
	FormID          string      // 消息来自哪个Actor
	FormServiceType string      // 消息来自哪个Actor的类型
	ToID            string      // 消息要发送给哪个Actor
	ToServiceType   string      // 消息要发送给哪个Actor的类型
	ReplyID         string      // 消息回复给哪个Actor
}
