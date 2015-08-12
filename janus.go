// entry point to the Janus ()
package main

import (
    "fmt"
    "github.com/codegangsta/cli"
    "github.com/xenolog/janus/config"
    "github.com/xenolog/janus/logger"
    "github.com/xenolog/janus/slack"
    "os"
    "path/filepath"
    "time"
)

const (
    Version = "0.0.1"
)

var (
    Log *logger.Logger
    App *cli.App
    err error
)

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

func getConfigAbsName(cfgname string) (string, error) {
    abs_path, err := filepath.Abs(cfgname)
    if err != nil {
        return "", fmt.Errorf("Wrong config path: '%s'", err)
    }
    return abs_path, nil
}

func runBot(c *cli.Context) {
    var (
        Sapi     *slack.Slack
        cfg      *config.Config
        abs_path string
        err      error
    )
    abs_path, err = getConfigAbsName(c.GlobalString("config"))
    if err != nil {
        Log.Error("Wrong config path: '%s'", err)
        return
    } else {
        Log.Printf("Config '%s' will be loaded.", c.GlobalString("config"))
    }

    cfg, err = config.New(abs_path)
    if err != nil {
        Log.Error("Config processing error: %s", err)
        return
    }

    Sapi = slack.New(&cfg.Janus)
    Sapi.Connect()
    Sapi.MainLoop()
    time.Sleep(30 * time.Second)
}

func runTest(c *cli.Context) {
    var err error
    var abs_path string
    Log.Log("Test started")

    if abs_path, err = getConfigAbsName(c.GlobalString("config")); err != nil {
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
