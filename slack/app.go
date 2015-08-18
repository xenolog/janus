package slack

import (
    "fmt"
    slacklib "github.com/abourget/slack"
    gangstalib "github.com/codegangsta/cli"
    //"gopkg.in/yaml.v2"
    //"io"
)

type DialogAppType struct {
    gangstalib.App
    SlackMsg        *slacklib.Msg
    SlackClient     *Slack
    responseMsgText *[]byte
}

// io.Writer compatible writer for
func (e *DialogAppType) Write(p []byte) (n int, err error) {
    *e.responseMsgText = append(*e.responseMsgText, p...)
    return len(p), nil
}

// Return response message text
func (e *DialogAppType) GetResponseMsgText() string {
    return fmt.Sprintf("%s", e.responseMsgText)
}

func NewDialogApp(msg *slacklib.Msg, client *Slack) *DialogAppType {
    // App := DialogAppType{
    //     Name:         "Janus",
    //     Usage:        "slack communication BOT",
    //     Version:      "0.0.0",
    //     BashComplete: false,
    //     Action:       helpCommand.Action,
    //     Compiled:     compileTime(),
    // }
    App := new(DialogAppType)
    App.Name = "Janus"
    App.Usage = "slack communication BOT"
    App.Version = "0.0.0"
    App.BashComplete = gangstalib.DefaultAppComplete
    // App.Action = helpCommand.Action
    // App.Compiled = compileTime()
    App.Writer = App
    // App.Flags = []cli.Flag{
    //     cli.BoolFlag{
    //         Name:  "debug",
    //         Usage: "Enable debug mode. Show more output",
    //     },
    // }
    // App.Commands = []DialogAppType.Command{{
    //     Name:  "irc",
    //     Usage: "irc-related commands",
    //     Action: func(c *DialogAppType.Context) {

    //     },
    // }, {
    //     Name:   "private",
    //     Usage:  "Create private channel for participants",
    //     Action: runBot,
    // }, {
    //     Name:   "test",
    //     Usage:  "Just run smal test.",
    //     Action: runTest,
    // },
    // }

    // App.Before = func(c *gangstalib.Context) error {
    //     //Log.Println("Janus started.")
    //     return nil
    // }
    // GangstaApp.CommandNotFound = func(c *gangstalib.Context, cmd string) {
    //     Log.Printf("Wrong command '%s'", cmd)
    //     os.Exit(1)
    // }    return App
    return App
}

//vim: set ts=4 sw=4 et :
