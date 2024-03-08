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

type ServerSettings struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// 读取配置文件字段对应内容
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.VP.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}