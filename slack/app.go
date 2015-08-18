package slack

import (
    //"fmt"
    gangsta "github.com/codegangsta/cli"
    //"io"
)

type SlackEvent struct {
    Xxx string
}

// io.Writer compotible writer for
func (e *SlackEvent) Write(p []byte) (n int, err error) {
    return 0, nil
}

type SlackApp struct {
    gangsta.App
    Event *SlackEvent
}

// func NewApp(ev *SlackEvent) *App {
//     return &SlackApp{
//         Name:         "Janus",
//         Usage:        "slack communication BOT",
//         Version:      "0.0.0",
//         BashComplete: false,
//         Action:       helpCommand.Action,
//         Compiled:     compileTime(),
//         Writer:       ev, // was os.Stdout, SlackEvent is compotible io.Writer
//         Event:        ev,
//     }
// }

// func (a *SlackApp) appInit(ev *SlackEvent) {
//     // Configure CLI flags and commands
//     a.App = gangsta.NewApp(ev)
//     a.App.Flags = []cli.Flag{
//         cli.BoolFlag{
//             Name:  "debug",
//             Usage: "Enable debug mode. Show more output",
//         },
//     }
//     a.App.Commands = []gangsta.Command{{
//         Name:  "irc",
//         Usage: "irc-related commands",
//         Action: func(c *gangsta.Context) {

//         },
//     }, {
//         Name:   "private",
//         Usage:  "Create private channel for participants",
//         Action: runBot,
//     }, {
//         Name:   "test",
//         Usage:  "Just run smal test.",
//         Action: runTest,
//     },
//     }
//     // a.App.Before = func(c *gangsta.Context) error {
//     //     //Log.Println("Janus started.")
//     //     return nil
//     // }
//     // a.App.CommandNotFound = func(c *gangsta.Context, cmd string) {
//     //     Log.Printf("Wrong command '%s'", cmd)
//     //     os.Exit(1)
//     // }
// }

// // func main() {
// //   cli.NewApp().Run(os.Args)
// // }

// // io.Writer:
// // type Writer interface {
// //     Write(p []byte) (n int, err error)
// // }

// func RunApp(msg, client *Slack) {

// }

//vim: set ts=4 sw=4 et :
