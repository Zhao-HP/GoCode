package test

//
//import (
//	"fmt"
//	"github.com/gookit/slog"
//	"github.com/gookit/slog/handler"
//	"testing"
//)
//
//var h = handler.NewConsoleHandler(slog.AllLevels)
//
//var LOG = slog.NewWithHandlers(h).WithFields(slog.M{
//	"category": "service",
//	"IP":       "127.0.0.1",
//})
//
//func init() {
//	fmt.Println("slog_test.init exec start")
//
//	// 使用JSON格式化日志
//	slog.SetFormatter(slog.NewJSONFormatter())
//
//	slog.Configure(func(logger *slog.SugaredLogger) {
//
//		l := logger.Formatter.(*slog.TextFormatter)
//		// 彩色日志开关
//		//l.EnableColor = true
//
//		// 时间格式设置
//		//l.TimeFormat = "2006-01-02 15:04:05"
//
//		l.Template = slog.NamedTemplate
//
//	})
//
//	fmt.Println("slog_test.init exec end")
//}
//
//func TestSlog(t *testing.T) {
//
//	LOG.Trace("trace log message")
//	LOG.Debug("debug log message")
//	LOG.Info("info log message")
//	LOG.Notice("notice log message")
//	LOG.Warn("warn log message")
//	LOG.Error("error log message")
//	LOG.Fatal("fatal log message")
//	LOG.Panic("panic log message")
//}
