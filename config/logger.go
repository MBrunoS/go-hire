package config

import (
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
}

func NewLogger() *Logger {
	writer := io.Writer(os.Stdout)

	return &Logger{
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

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

func (l *Logger) DebugF(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) InfoF(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) WarningF(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l *Logger) ErrorF(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
