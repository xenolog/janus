package slack

import (
    "fmt"
    slacklib "github.com/abourget/slack"
    gangstalib "github.com/codegangsta/cli"
    "strings"
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

// run created app. May be started as go-routine
func (e *DialogAppType) RunApp() {
    cmd := strings.Split(e.SlackMsg.Text, "\n")[0]
    log.Debug("Slack DialogApp command: '%s'", cmd)
    e.Run(strings.Fields(cmd))
    log.Debug("response> %s", e.responseMsgText)
}

// create App for Dialog req
func NewDialogApp(msg *slacklib.Msg, client *Slack) *DialogAppType {
    App := new(DialogAppType)
    App.Name = "Janus"
    App.Usage = "slack communication BOT"
    App.Version = "0.0.0"
    App.BashComplete = gangstalib.DefaultAppComplete
    // App.Action = helpCommand.Action
    // App.Compiled = compileTime()
    App.Writer = App
    App.Flags = []gangstalib.Flag{
        gangstalib.BoolFlag{
            Name:  "debug",
            Usage: "Enable debug mode. Show more output",
        },
    }
    App.Commands = []gangstalib.Command{{
        Name:  "irc",
        Usage: "irc-related commands",
        Action: func(c *gangstalib.Context) {

        },
    }, {
        Name:  "private",
        Usage: "Create private channel for participants",
        Action: func(c *gangstalib.Context) {

        },
        // }, {
        //     Name:   "test",
        //     Usage:  "Just run smal test.",
        //     Action: runTest,
    },
    }

    App.Before = func(c *gangstalib.Context) error {
        //Log.Println("Janus started.")
        return nil
    }
    // App.CommandNotFound = func(c *gangstalib.Context, cmd string) {
    //     Log.Printf("Wrong command '%s'", cmd)
    //     os.Exit(1)
    // }
    App.SlackClient = client
    App.SlackMsg = msg
    return App
}

//vim: set ts=4 sw=4 et :
