// Singletone logger
package logger

import (
    "fmt"
    "io"
    "os"
    "sync"
)

const (
    LOG_D  = 0
    LOG_I  = 2
    LOG_W  = 4
    LOG_E  = 6
    LOG_F  = 8
    LOG_Dl = "D"
    LOG_Il = "I"
    LOG_Wl = "W"
    LOG_El = "E"
    LOG_Fl = "F"
)

type Logger struct {
    sync.Mutex           // extend for ensures atomic writes; protects the following fields
    configured      bool // flag whether logger configured
    defaultFacility int8 // may be LOG_D..LOG_F
    minFacility     int8 // minimal allowed facility for output
    //syslogServers     []string        // todo
    //logFiles          []string        // todo
    flag int // properties
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

func (l *Logger) checkFacility(facility int8, format string) error {
    if LOG_D > facility || LOG_F < facility {
        err := fmt.Sprintf(format, facility)
        l.Error(err)
        return fmt.Errorf(err)
    }
    return nil
}

// Setup default facility for logging
func (l *Logger) SetDefaultFacility(facility int8) error {
    if err := l.checkFacility(facility, "Try to set disallowed facility as default '%s'"); err != nil {
        return err
    }
    l.defaultFacility = facility
    return nil
}

// Setup minimal facility for logging
func (l *Logger) SetMinimalFacility(facility int8) error {
    if err := l.checkFacility(facility, "Try to set disallowed facility as minimal '%s'"); err != nil {
        return err
    }
    l.minFacility = facility
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

func New() *Logger {
    if !main_logger.configured {
        main_logger.defaultFacility = LOG_I
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

//vim: set ts=4 sw=4 et :
