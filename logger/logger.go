// entry point to the Janus ()
package logger

import (
    "log"
    "os"
)

var main_logger *log.Logger

func init() {
    main_logger = log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

func GetLogger() *log.Logger {
    return main_logger
}
