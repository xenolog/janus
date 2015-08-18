package logger

import (
    "fmt"
)

// Log message with default facility (obsoleted alias for Log)
func (l *Logger) Printf(format string, v ...interface{}) {
    l.Output(2, LOG_I, fmt.Sprintf(format, v...))
}

// Log message with default facility (obsoleted alias for Log)
func (l *Logger) Println(format string, v ...interface{}) {
    l.Output(2, LOG_I, fmt.Sprintf(format, v...))
}

// Log message with 'E' facility (Alias for Error)
func (l *Logger) Errorf(format string, v ...interface{}) {
    l.Output(2, LOG_E, fmt.Sprintf(format, v...))
}

//vim: set ts=4 sw=4 et :
