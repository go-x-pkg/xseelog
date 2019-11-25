package xseelog

import (
	"fmt"
	"io"
	"path/filepath"
	"time"

	"github.com/cihub/seelog"
	"github.com/go-x-pkg/dumpctx"
	"github.com/go-x-pkg/log"
)

type Config struct {
	Dir string `yaml:"dir"`

	DisableConsole bool `yaml:"disable-console"`
	DisableFile    bool `yaml:"disable-file"`

	Logs map[string]*ConfigLogger `yaml:"loggers"`
}

func (c *Config) Dump(ctx *dumpctx.Ctx, w io.Writer) {
	fmt.Fprintf(w, "%sdir: %s\n", ctx.Indent(), c.Dir)
	fmt.Fprintf(w, "%sdisable-console: %t\n", ctx.Indent(), c.DisableConsole)
	fmt.Fprintf(w, "%sdisable-file: %t\n", ctx.Indent(), c.DisableFile)

	fmt.Fprintf(w, "%sloggers", ctx.Indent())

	if len(c.Logs) != 0 {
		fmt.Fprintf(w, " (x%d):\n", len(c.Logs))

		for _, r := range c.Logs {
			ctx.WrapList(func() { r.Dump(ctx, w) })
		}
	} else {
		fmt.Fprint(w, ": ~\n")
	}
}

func (c *Config) VV() {
	for _, cl := range c.Logs {
		if cl.LevelMin > log.Debug {
			cl.LevelMin = log.Debug
		}
	}
}

func (c *Config) VVV() {
	for _, cl := range c.Logs {
		cl.LevelMin = log.Trace
	}
}

func (c *Config) Quiet() {
	for _, cl := range c.Logs {
		cl.LevelMin = log.Critical
		cl.LevelMax = log.Critical
	}
}

func (c *Config) Ensure(name, prefix string, min, max log.Level) {
	if _, ok := c.Logs[name]; ok {
		return
	}

	if c.Logs == nil {
		c.Logs = make(map[string]*ConfigLogger)
	}

	cl := &ConfigLogger{
		Prefix:   prefix,
		LevelMin: min,
		LevelMax: max,
	}
	c.Logs[name] = cl
}

func (c *Config) Loggers() (*Loggers, error) {
	loggers := NewLoggers()

	for name, cl := range c.Logs {
		if logger, err := cl.logger(c); err != nil {
			return loggers, fmt.Errorf("error create console logger: %w", err)
		} else {
			loggers.m[name] = logger
		}
	}

	return loggers, nil
}

func NewConfig() *Config {
	c := new(Config)
	c.Logs = make(map[string]*ConfigLogger)
	return c
}

type ConfigLogger struct {
	Prefix string `yaml:"prefix"`

	LevelMin log.Level `yaml:"level-min"`
	LevelMax log.Level `yaml:"level-max"`
}

func (cl *ConfigLogger) Dump(ctx *dumpctx.Ctx, w io.Writer) {
	ctx.EmitPrefix(w)

	fmt.Fprintf(w, "prefix: %s\n", cl.Prefix)

	ctx.Enter()
	defer ctx.Leave()

	fmt.Fprintf(w, "%slevel-min: %s\n", ctx.Indent(), cl.LevelMin)
	fmt.Fprintf(w, "%slevel-max: %s\n", ctx.Indent(), cl.LevelMax)
}

func (cl *ConfigLogger) logger(c *Config) (seelog.LoggerInterface, error) {
	var (
		receivers []interface{}
	)

	if !c.DisableConsole {
		consoleWriter, err := seelog.NewConsoleWriter()
		if err != nil {
			return nil, fmt.Errorf("error create console logger: %w", err)
		}

		formatterTrace, err := seelog.NewFormatter(formatTrace(cl.Prefix))
		if err != nil {
			return nil, fmt.Errorf("error gen format-trace: %w", err)
		}

		formatterDebug, err := seelog.NewFormatter(formatDebug(cl.Prefix))
		if err != nil {
			return nil, fmt.Errorf("error gen format-debug: %w", err)
		}

		formatterInfo, err := seelog.NewFormatter(formatInfo(cl.Prefix))
		if err != nil {
			return nil, fmt.Errorf("error gen format-info: %w", err)
		}

		formatterWarn, err := seelog.NewFormatter(formatWarn(cl.Prefix))
		if err != nil {
			return nil, fmt.Errorf("error gen format-warn: %w", err)
		}

		formatterError, err := seelog.NewFormatter(formatError(cl.Prefix))
		if err != nil {
			return nil, fmt.Errorf("error gen format-error: %w", err)
		}

		formatterCritical, err := seelog.NewFormatter(formatCritical(cl.Prefix))
		if err != nil {
			return nil, fmt.Errorf("error gen format-critical: %w", err)
		}

		dispatcherTrace, err := seelog.NewFilterDispatcher(formatterTrace, []interface{}{consoleWriter},
			seelog.TraceLvl)
		if err != nil {
			return nil, fmt.Errorf("error gen dispatcher console for trace: %w", err)
		}
		dispatcherDebug, err := seelog.NewFilterDispatcher(formatterDebug, []interface{}{consoleWriter},
			seelog.DebugLvl)
		if err != nil {
			return nil, fmt.Errorf("error gen dispatcher console for debug: %w", err)
		}
		dispatcherInfo, err := seelog.NewFilterDispatcher(formatterInfo, []interface{}{consoleWriter},
			seelog.InfoLvl)
		if err != nil {
			return nil, fmt.Errorf("error gen dispatcher console for info: %w", err)
		}
		dispatcherWarn, err := seelog.NewFilterDispatcher(formatterWarn, []interface{}{consoleWriter},
			seelog.WarnLvl)
		if err != nil {
			return nil, fmt.Errorf("error gen dispatcher console for warn: %w", err)
		}
		dispatcherError, err := seelog.NewFilterDispatcher(formatterError, []interface{}{consoleWriter},
			seelog.ErrorLvl)
		if err != nil {
			return nil, fmt.Errorf("error gen dispatcher console for error: %w", err)
		}
		dispatcherCritical, err := seelog.NewFilterDispatcher(formatterCritical, []interface{}{consoleWriter},
			seelog.CriticalLvl)
		if err != nil {
			return nil, fmt.Errorf("error gen dispatcher console for critical: %w", err)
		}

		receivers = append(receivers,
			dispatcherTrace,
			dispatcherDebug,
			dispatcherInfo,
			dispatcherWarn,
			dispatcherError,
			dispatcherCritical)
	}

	if !c.DisableFile {
		pathAccess := filepath.Join(c.Dir, fileAccess)
		pathError := filepath.Join(c.Dir, fileError)

		// logrotate writer open here
		fileAccessWriter, err := seelog.NewFileWriter(pathAccess)
		if err != nil {
			return nil, fmt.Errorf("error open file-access writer: %w", err)
		}

		// logrotate writer open here
		fileErrorWriter, err := seelog.NewFileWriter(pathError)
		if err != nil {
			return nil, fmt.Errorf("error open file-error writer: %w", err)
		}

		formatterFile, err := seelog.NewFormatter(formatFile(cl.Prefix))
		if err != nil {
			return nil, fmt.Errorf("error gen format-file: %w", err)
		}

		dispatcherFileAccess, err := seelog.NewFilterDispatcher(formatterFile, []interface{}{fileAccessWriter},
			seelog.TraceLvl, seelog.DebugLvl, seelog.InfoLvl,
			seelog.WarnLvl, seelog.ErrorLvl, seelog.CriticalLvl,
		)
		if err != nil {
			return nil, fmt.Errorf("error gen dispatcher file-access: %w", err)
		}

		dispatcherFileError, err := seelog.NewFilterDispatcher(formatterFile, []interface{}{fileErrorWriter},
			seelog.WarnLvl, seelog.ErrorLvl, seelog.CriticalLvl)
		if err != nil {
			return nil, fmt.Errorf("error gen dispatcher file-error: %w", err)
		}

		receivers = append(receivers,
			dispatcherFileAccess,
			dispatcherFileError)
	}

	constraints, err := seelog.NewMinMaxConstraints(FromLogLevel(cl.LevelMin), FromLogLevel(cl.LevelMax))
	if err != nil {
		return nil, fmt.Errorf("error gen format-critical: %w", err)
	}

	rootFormatter, _ := seelog.NewFormatter("")

	root, err := seelog.NewSplitDispatcher(rootFormatter, receivers)
	if err != nil {
		return nil, fmt.Errorf("error gen root dispatcher: %w", err)
	}

	logger, err := seelog.NewAsyncTimerLogger(seelog.NewLoggerConfig(constraints, nil, root), time.Second)
	if err != nil {
		return nil, fmt.Errorf("error gen logger: %w", err)
	}

	return logger, nil
}
