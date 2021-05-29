package pkg

import (
	"fmt"
	"go.uber.org/zap"
	"xorm.io/xorm/log"
)

// 给 orm 用
type XormLogger struct {
	level   log.LogLevel
	showSQL bool
}

func (x *XormLogger)Debug(v ...interface{}){
	Logger.WithOptions(zap.AddCallerSkip(2))
	Logger.Debug(fmt.Sprintln(v...))
	Logger.WithOptions(zap.AddCallerSkip(-2))
}
func (x *XormLogger)Debugf(format string, v ...interface{}){
	Logger.Debug(fmt.Sprintf(format,v...))
}
func (x *XormLogger)Error(v ...interface{}){
	Logger.Error(fmt.Sprintln(v...))
}
func (x *XormLogger)Errorf(format string, v ...interface{}){
	Logger.Error(fmt.Sprintf(format,v...))
}
func (x *XormLogger)Info(v ...interface{}){
	Logger.WithOptions(zap.AddCallerSkip(2))
	Logger.Info(fmt.Sprintln(v...))
	Logger.WithOptions(zap.AddCallerSkip(-2))

}
func (x *XormLogger)Infof(format string, v ...interface{}){
	Logger.Info(fmt.Sprintf(format,v...))
}
func (x *XormLogger)Warn(v ...interface{}){
	Logger.Warn(fmt.Sprintln(v...))
}
func (x *XormLogger)Warnf(format string, v ...interface{}){
	Logger.Warn(fmt.Sprintf(format,v...))
}

func (x *XormLogger)Level() log.LogLevel{
	return x.level
}
func (x *XormLogger)SetLevel(l log.LogLevel){
	x.level = l
}

func (x *XormLogger)ShowSQL(show ...bool){
	if len(show) == 0 {
		x.showSQL = true
		return
	}
	x.showSQL = show[0]
}
func (x *XormLogger)IsShowSQL() bool{
	return x.level == log.LOG_DEBUG
}