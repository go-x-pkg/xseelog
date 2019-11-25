package xseelog

import (
	"sync"

	"github.com/cihub/seelog"
)

type Loggers struct {
	sync.RWMutex
	m map[string]seelog.LoggerInterface
}

func (l *Loggers) ReplaceRoot() {
	l.RLock()
	defer l.RUnlock()

	for name, logger := range l.m {
		if name == "app" || name == "root" || name == "base" {
			seelog.ReplaceLogger(logger)
			return
		}
	}

	// pick first logger
	for _, logger := range l.m {
		seelog.ReplaceLogger(logger)
		return
	}
}

func (l *Loggers) Close() {
	l.Lock()
	defer l.Unlock()

	for _, logger := range l.m {
		logger.Close()
	}
}

func (l *Loggers) Logger(name string) seelog.LoggerInterface {
	l.RLock()
	defer l.RUnlock()

	if logger, ok := l.m[name]; ok {
		return logger
	}

	return seelog.Default
}

/* impl log.Loggers for Loggers */
func (l *Loggers) ByName(name string) interface{} {
	return l.Logger(name)
}

func NewLoggers() *Loggers {
	l := new(Loggers)
	l.m = make(map[string]seelog.LoggerInterface)
	return l
}
