package setting

import "time"

type DatabaseSettings struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type AppSettings struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
	UsclientPath    string
}

type RedisSettings struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

type ServerSettings struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type JWTSettings struct {
	AppSecret string        //秘钥
	Issuer    string        //签发者
	Expire    time.Duration //过期时间
}

// 读取配置文件字段对应内容
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.VP.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
