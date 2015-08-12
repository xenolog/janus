// entry point to the Janus ()
package main

import (
    "github.com/codegangsta/cli"
    "github.com/xenolog/janus/config"
    "github.com/xenolog/janus/logger"
    "os"
    "path/filepath"
)

const (
    Version = "0.0.1"
)

var (
    Log *logger.Logger
    App *cli.App
    err error
)

func runBot(c *cli.Context) {
    Log.Printf("Not implemented :(")
}

func init() {
    // Setup logger
    Log = logger.New()

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
            Name:  "config, c",
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
    }, {
        Name:   "test",
        Usage:  "Just run smaLog test.",
        Action: runTest,
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

func runTest(c *cli.Context) {
    var err error
    var abs_path string
    Log.Log("Test started")
    // Log.Printf("config file is: '%s'", c.GlobalString("config"))
    if abs_path, err = filepath.Abs(c.GlobalString("config")); err != nil {
        Log.Error("Wrong config path: '%s'", err)
        return
    } else {
        Log.Printf("Config '%s' will be loaded.", c.GlobalString("config"))
    }

    cfg, err := config.New(abs_path)
    if err != nil {
        Log.Error("Config processing error: %s", err)
    } else {
        Log.Info("config: '%s'", cfg)
        Log.Info("xxx: '%s'", cfg.Users["xenolog"].Irc.Username)
        Log.Info("xxx: '%s'", cfg.Users["xenolog"].Irc.Password)
    }
    Log.Log("Test completed")
}
