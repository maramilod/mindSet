package database

import (
	c "mind-set/config"
	"sync"
)

var once sync.Once

func InitData() {
	once.Do(func() {
		if c.Config.Mysql.Enable {
			// mysql
			initMysql()
		}
		if c.Config.Redis.Enable {
			// redis
			initRedis()
		}
		if c.Config.Firebase.Enable {
			// redis
			initFirebase()
		}
	})
}
