package global

import (
	"my-service/pkg/setting"
)

var (
	DBSettings     *setting.DatabaseSettings
	AppSettings    *setting.AppSettings
	RedisSettings  *setting.RedisSettings
	ServerSettings *setting.ServerSettings
	JWTSettings    *setting.JWTSettings
	// ZMQSettings    *setting.ZMQSettings
	// MosqtSettings  *setting.MosqtSettings
)
