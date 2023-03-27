package actor

import "sync"

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
}

func GetActorManager() *ActorManager {
	return _actorManager
}
