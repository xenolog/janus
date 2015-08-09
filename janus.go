// entry point to the Janus ()
package main

import (
    "github.com/codegangsta/cli"
    _ "github.com/xenolog/janus/config"
    "github.com/xenolog/janus/logger"
    "log"
    "os"
)

const (
    Version = "0.0.1"
)

var (
    Log *log.Logger
    App *cli.App
    err error
)

func runBot(c *cli.Context) {
    Log.Printf("Not implemented :(")
}

func init() {
    // Setup logger
    Log = logger.GetLogger()

    // Configure CLI flags and commands
    App = cli.NewApp()
    App.Name = "Janus"
    App.Version = Version
    App.EnableBashCompletion = true
    App.Usage = "BOT for Slack--IRC transparent proxying with support multi user"
    App.Flags = []cli.Flag{
        cli.BoolFlag{
            Name:  "debug",
            Usage: "Enable debug mode. Show more output",
        },
        cli.StringFlag{
            Name:  "c, config",
            Usage: "Specify config file (default: ./janus.jaml)",
        },
    }
    App.Commands = []cli.Command{{
        Name:   "runBot",
        Usage:  "Run bot",
        Action: runBot,
    }, {
        Name:   "control",
        Usage:  "Manipulate already started bots.",
        Action: runBot,
    },
    }
    App.Before = func(c *cli.Context) error {
        Log.Println("Janus started.")
        return nil
    }
    App.CommandNotFound = func(c *cli.Context, cmd string) {
        Log.Printf("Wrong command '%s'", cmd)
        os.Exit(1)
    }
}

func main() {
    App.RunAndExitOnError()

}
