package logger

import (
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

type DefaultLogger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
}

func NewDefaultLogger() *DefaultLogger {
	writer := io.Writer(os.Stdout)

	return &DefaultLogger{
		debug:   log.New(writer, colorString("[DEBUG] ", color.FgCyan), log.Ldate|log.Ltime),
		info:    log.New(writer, colorString("[INFO] ", color.FgGreen), log.Ldate|log.Ltime),
		warning: log.New(writer, colorString("[WARNING] ", color.FgYellow), log.Ldate|log.Ltime),
		err:     log.New(writer, colorString("[ERROR] ", color.FgRed), log.Ldate|log.Ltime),
	}
}

func colorString(s string, c color.Attribute) string {
	col := color.New(c).SprintFunc()
	return col(s)
}

func (l *DefaultLogger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *DefaultLogger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *DefaultLogger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *DefaultLogger) Error(v ...interface{}) {
	l.err.Println(v...)
}

func (l *DefaultLogger) DebugF(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *DefaultLogger) InfoF(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *DefaultLogger) WarningF(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l *DefaultLogger) ErrorF(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
