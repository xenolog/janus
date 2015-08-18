// entry point to the Janus ()
package main

import (
    "github.com/codegangsta/cli"
    "github.com/xenolog/janus/config"
    "github.com/xenolog/janus/logger"
    "github.com/xenolog/janus/slack"
    "os"
    "sync"
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
        Log.Info("Janus started.")
        if c.GlobalBool("debug") {
            Log.SetMinimalFacility(logger.LOG_D)
        } else {
            Log.SetMinimalFacility(logger.LOG_I)
        }
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

// func getConfigAbsName(cfgname string) (string, error) {
//     abs_path, err := filepath.Abs(cfgname)
//     if err != nil {
//         return "", fmt.Errorf("Wrong config path: '%s'", err)
//     }
//     return abs_path, nil
// }

func runBot(c *cli.Context) {
    var (
        Sapi     *slack.Slack
        cfg      *config.Config
        cfg_path string
        err      error
        wg       sync.WaitGroup
    )
    cfg_path = c.GlobalString("config")
    Log.Info("Config '%s' will be loaded.", cfg_path)
    cfg, err = config.New(cfg_path)
    if err != nil {
        Log.Error("Config processing error: %s", err)
        return
    }

    Sapi = slack.New(&cfg.Janus)
    Sapi.Connect()
    // start Messsage loop
    wg.Add(1)
    go func() {
        defer wg.Done()
        Sapi.MessageLoop()
    }()
    // start Channel info loop
    wg.Add(1)
    go func() {
        defer wg.Done()
        Sapi.ChannelLoop()
    }()
    // all loops started
    wg.Wait()

}

func runTest(c *cli.Context) {
    Log.Log("Test started")
    Log.Debug("XXX")
    Log.Log("Test completed")
}

//vim: set ts=4 sw=4 et :
