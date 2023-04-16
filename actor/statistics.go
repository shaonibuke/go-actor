package actor

import "time"

// 用于每个actor 的统计信息
type Statistics struct {
	StartAt int64 // 开始时间
	// 用于统计每个actor的消息处理时间
	// 消息名 -> 消息处理时间ms
	MsgHandleTime map[string]int64
	// 用于统计每个actor的消息处理次数
	// 消息名 -> 消息处理次数ms
	MsgHandleCount map[string]int64
}

// NewStatistics 创建一个新的Statistics
func NewStatistics() *Statistics {
	return &Statistics{
		StartAt:        time.Now().Unix(),
		MsgHandleTime:  make(map[string]int64),
		MsgHandleCount: make(map[string]int64),
	}
}

// AddMsgHandleTime 添加消息处理时间
func (s *Statistics) AddMsgHandleTime(msgName string, handleTime int64) {
	s.MsgHandleTime[msgName] += handleTime
}

// AddMsgHandleCount 添加消息处理次数
func (s *Statistics) AddMsgHandleCount(msgName string) {
	s.MsgHandleCount[msgName]++
}

// GetMsgHandleTime 获取消息处理时间
func (s *Statistics) GetMsgHandleTime(msgName string) int64 {
	return s.MsgHandleTime[msgName]
}

// GetMsgHandleCount 获取消息处理次数
func (s *Statistics) GetMsgHandleCount(msgName string) int64 {
	return s.MsgHandleCount[msgName]
}

// GetRunTime 获取运行时间
func (s *Statistics) GetRunTime() int64 {
	return time.Now().Unix() - s.StartAt
}

// GetMsgHandleTimeAvg 获取消息处理时间平均值
func (s *Statistics) GetMsgHandleTimeAvg(msgName string) int64 {
	return s.MsgHandleTime[msgName] / s.MsgHandleCount[msgName]
}

// 获取每种消息各自的平均处理时间
func (s *Statistics) GetMsgHandleTimeAvgAllByMsg() map[string]int64 {
	msgHandleTimeAvg := make(map[string]int64)
	for k, v := range s.MsgHandleTime {
		msgHandleTimeAvg[k] = v / s.MsgHandleCount[k]
	}
	return msgHandleTimeAvg
}

// GetMsgHandleCountAll 获取所有消息的处理次数
func (s *Statistics) GetMsgHandleCountAll() int64 {
	var sum int64
	for _, v := range s.MsgHandleCount {
		sum += v
	}
	return sum
}

// 获取所有消息的平均处理时间
func (s *Statistics) GetMsgHandleTimeAvgAll() int64 {
	var sum int64
	for _, v := range s.MsgHandleTime {
		sum += v
	}
	return sum / s.GetRunTime()
}
