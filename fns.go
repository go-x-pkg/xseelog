package xseelog

import (
	"github.com/cihub/seelog"
	"github.com/go-x-pkg/log"
)

/* impl From<log.Level> for seelog.LogLevel */
func FromLogLevel(lvl log.Level) seelog.LogLevel {
	switch lvl {
	case log.Quiet:
		return seelog.Off
	case log.Trace:
		return seelog.TraceLvl
	case log.Debug:
		return seelog.DebugLvl
	case log.Info:
		return seelog.InfoLvl
	case log.Warn:
		return seelog.WarnLvl
	case log.Error:
		return seelog.ErrorLvl
	case log.Critical:
		return seelog.CriticalLvl
	}

	return seelog.Off
}
