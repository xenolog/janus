// entry point to the Janus ()
package main

import (
    "fmt"
    //"github.com/codegangsta/cli"
    "github.com/xenolog/janus/logger"
    "log"
    // "log"
    // "os"
)

var (
    Log *log.Logger
)

func init() {
    Log = logger.GetLogger()
}

func main() {
    fmt.Println("qqq!")
    Log.Printf("qqq!")
}
