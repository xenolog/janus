// Singletone logger
package logger

import (
    "fmt"
    "io"
    "os"
    "sync"
)

type Logger struct {
    configured        bool          // flag whether logger configured
    defaultFacility   byte          // may be "I", "W", "E", "F"
    allowedFacilities map[byte]bool //
    //syslogServers     []string        // todo
    //logFiles          []string        // todo
    mu   sync.Mutex // ensures atomic writes; protects the following fields
    flag int        // properties
    out  io.Writer
    buf  []byte // for accumulating text to write
}

var main_logger *Logger

// add syslog server for logging
// func (l *Logger) AddSyslogServer(server string) error {
//     const errMsg = "Non implemented :("
//     l.crtLogger.Printf(errMsg)
//     return fmt.Errorf(errMsg)
// }

// Setup default facility for logging
func (l *Logger) SetDefaultFacility(facility []byte) error {
    f := facility[0]
    if !l.allowedFacilities[f] {
        err := fmt.Sprintf("Try to set as default disallowed facility '%s'", facility)
        l.Warn(err)
        return fmt.Errorf(err)
    }
    l.defaultFacility = f
    return nil
}

// Enable logging to console
func (l *Logger) EnableConsoleLog() {
    if l.out == nil {
        l.flag = Ldate | Lmicroseconds | Lshortfile
        l.out = os.Stderr
    }
}

// Disable logging to console
func (l *Logger) DisableConsoleLog() {
    l.out = nil
}

// Log message with default facility
func (l *Logger) Log(format string, v ...interface{}) {
    l.Output(2, l.defaultFacility, fmt.Sprintf(format, v...))
}

// Log message with default facility (obsoleted alias for Log)
func (l *Logger) Printf(format string, v ...interface{}) {
    l.Output(2, 'I', fmt.Sprintf(format, v...))
}

// Log message with default facility (obsoleted alias for Log)
func (l *Logger) Println(format string, v ...interface{}) {
    l.Output(2, 'I', fmt.Sprintf(format, v...))
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

// Log message with 'E' facility (Alias for Error)
func (l *Logger) Errorf(format string, v ...interface{}) {
    l.Output(2, 'E', fmt.Sprintf(format, v...))
}

// Log message with 'F' facility and execute os.Exit(1)
func (l *Logger) Fail(format string, v ...interface{}) {
    l.Output(2, 'F', fmt.Sprintf(format, v...))
    os.Exit(1)
}

func New() *Logger {
    if !main_logger.configured {
        main_logger.defaultFacility = 'I'
        main_logger.allowedFacilities = map[byte]bool{'I': true, 'W': true, 'E': true, 'F': true}
        main_logger.EnableConsoleLog()
        main_logger.configured = true
    }
    //todo: Also may be used form:
    // log.SetOutput(io.MultiWriter(os.Stdout, logFile))
    return main_logger
}

func init() {
    main_logger = new(Logger)
}
