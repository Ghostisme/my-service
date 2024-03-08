package global

import (
	"my-project-admin-service/pkg/setting"
)

var (
	DBSettings     *setting.DatabaseSettings
	AppSettings    *setting.AppSettings
	ServerSettings *setting.ServerSettings
	// JWTSettings    *setting.JWTSettings
	// ZMQSettings    *setting.ZMQSettings
	// MosqtSettings  *setting.MosqtSettings
)
