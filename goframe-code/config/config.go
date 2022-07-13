package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"time"
)

const IniFile = "ini/app.ini"

type Config struct {
	Server Server
	Log    Log
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Log struct {
	Level         string // 级别
	ShowLine      bool   // 显示行
	Format        string // 输出
	Prefix        string // 日志前缀
	Director      string // 日志文件夹
	LinkName      string // 软链接名称
	EncodeLevel   string // 编码级
	StacktraceKey string // 栈名
	LogInConsole  bool   // 输出控制台

}

var GLOBAL_CONF = &Config{}

type ReadCfg interface {
	Read(cfg *ini.File)
}

func (s *Server) Read(cfg *ini.File) {
	section := cfg.Section("server")
	s.RunMode = section.Key("RunMode").String()
	s.HttpPort, _ = section.Key("HttpPort").Int()
	s.ReadTimeout, _ = section.Key("ReadTimeout").Duration()
	s.WriteTimeout, _ = section.Key("WriteTimeout").Duration()
}

func (l *Log) Read(cfg *ini.File) {
	section := cfg.Section("log")
	l.Level = section.Key("level").String()
	l.Format = section.Key("format").String()
	l.Prefix = section.Key("prefix").String()
	l.Director = section.Key("director").String()
	l.LinkName = section.Key("link-name").String()
	l.ShowLine, _ = section.Key("show-line").Bool()
	l.EncodeLevel = section.Key("encode-level").String()
	l.StacktraceKey = section.Key("stacktrace-key").String()
	l.LogInConsole, _ = section.Key("log-in-console").Bool()
}

func init() {
	cfg, err := ini.Load(IniFile)
	if err != nil {
		fmt.Println("读取ini文件失败", err)
	}
	GLOBAL_CONF.Server.Read(cfg)
	GLOBAL_CONF.Log.Read(cfg)
}
