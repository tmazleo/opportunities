package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug *log.Logger
	info *log.Logger
	waning *log.Logger
	err *log.Logger
	writer io.Writer
}

func NewLogger(prefix string) *Logger{

	writer := io.Writer(os.Stdout)
	logger := log.New(writer,prefix,log.Ldate|log.Ltime)

	return &Logger{
		debug: log.New(writer,"DEBUG: ", logger.Flags()),
		info: log.New(writer,"INFO: ", logger.Flags()),
		waning: log.New(writer,"WARNING: ", logger.Flags()),
		err: log.New(writer,"ERROR: ", logger.Flags()),
		writer: writer,
	}
}

func (l *Logger) Debug(v ...interface{}){
	l.debug.Println(v ...)
}
func (l *Logger) Info(v ...interface{}){
	l.info.Println(v ...)
}
func (l *Logger) Warn(v ...interface{}){
	l.waning.Println(v ...)
}
func (l *Logger) Error(v ...interface{}){
	l.err.Println(v ...)
}

func (l *Logger) DebugF(format string, v ...interface{}){
	l.debug.Printf(format, v ...)
}
func (l *Logger) InfoF(format string, v ...interface{}){
	l.info.Printf(format, v ...)
}
func (l *Logger) WarnF(format string, v ...interface{}){
	l.waning.Printf(format, v ...)
}
func (l *Logger) ErrorF(format string, v ...interface{}){
	l.err.Printf(format, v ...)
}
