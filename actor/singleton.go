package actor

import (
	"os"
	"sync"

	"github.com/charmbracelet/log"
)

// 管理者应该成为单利模式
var (
	// ActorManager 单例
	_actorManager *ActorManager
	once          sync.Once
)

func init() {
	once.Do(func() {
		_actorManager = newActorManager()
	})
	logger := log.New(os.Stderr)
	logger.SetLevel(log.DebugLevel)
	logger.SetTimeFormat(log.DefaultTimeFormat)
	logger.SetReportTimestamp(true)
	logger.SetReportCaller(true)
	log.SetDefault(logger)
}

func GetActorManager() *ActorManager {
	return _actorManager
}
