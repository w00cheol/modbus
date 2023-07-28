package modbus

import (
	"log"
	"os"
)

type Logger struct {
	debug log.Logger
	err   log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		debug: *log.New(os.Stdout, "[DEBUG] ", 0),
		err:   *log.New(os.Stdout, "[ERROR] ", 0),
	}
}

func (logger *Logger) Debug(s ...string) {
	logger.debug.Println(s)
}

func (logger *Logger) Error(err error) {
	logger.debug.Println(err)
}
