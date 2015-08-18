package logger

import (
    "fmt"
    "os"
)

// Log message with default facility
func (l *Logger) Log(format string, v ...interface{}) {
    l.Output(2, l.defaultFacility, fmt.Sprintf(format, v...))
}

// Log message with 'I' facility
func (l *Logger) Info(format string, v ...interface{}) {
    l.Output(2, 'I', fmt.Sprintf(format, v...))
}

// Log message with 'W' facility
func (l *Logger) Warn(format string, v ...interface{}) {
    l.Output(2, 'W', fmt.Sprintf(format, v...))
}

// Log message with 'E' facility
func (l *Logger) Error(format string, v ...interface{}) {
    l.Output(2, 'E', fmt.Sprintf(format, v...))
}

// Log message with 'F' facility and execute os.Exit(1)
func (l *Logger) Fail(format string, v ...interface{}) {
    l.Output(2, 'F', fmt.Sprintf(format, v...))
    os.Exit(1)
}

//vim: set ts=4 sw=4 et :
