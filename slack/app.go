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
func (e *DialogAppType) Write(p []byte) (int, error) {
    if e.responseMsgText == nil {
        e.responseMsgText = &p
    } else {
        *e.responseMsgText = append(*e.responseMsgText, p...)
    }
    return len(p), nil
}

// Return response message text
func (e *DialogAppType) GetResponseMsgText() string {
    return fmt.Sprintf("%s", *e.responseMsgText)
}

func (e *DialogAppType) GetSlackClient() *Slack {
    return e.SlackClient
}

func (e *DialogAppType) GetSlackMsg() slacklib.Msg {
    return *e.SlackMsg
}

// type DialogAppInterface interface {
//     GetResponseMsgText()
//     GetSlackClient()
//     GetSlackMsg()
// }

// run created app. May be started as go-routine
func (e *DialogAppType) RunApp(msg slacklib.Msg) {
    cmd := strings.Split(msg.Text, "\n")[0]
    log.Debug("Slack DialogApp command: '%s'", cmd)
    e.Run(strings.Fields(cmd))
    log.Debug("response> %s", *e.responseMsgText)
}

// create App for Dialog req
func NewDialogApp(client *Slack) *DialogAppType {
    App := new(DialogAppType)
    App.Name = "Janus"
    App.Usage = "slack communication BOT"
    App.HideVersion = true
    //App.BashComplete = gangstalib.DefaultAppComplete
    // App.Compiled = compileTime()
    App.Action = func(c *gangstalib.Context) {
        log.Debug("DialogApp 'App.Action: Start'")
        args := c.Args()
        if args.Present() {
            gangstalib.ShowCommandHelp(c, args.First())
        } else {
            gangstalib.ShowSubcommandHelp(c)
        }
        log.Debug("DialogApp 'App.Action: End'")
    }
    App.Writer = App
    // App.Flags = []gangstalib.Flag{
    //     gangstalib.BoolFlag{
    //         Name:  "debug",
    //         Usage: "Enable debug mode. Show more output",
    //     },
    // }
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
        log.Debug("DialogApp 'before'")
        return nil
    }

    // App.CommandNotFound = func(c *gangstalib.Context, cmd string) {
    //     Log.Printf("Wrong command '%s'", cmd)
    //     os.Exit(1)
    // }
    App.SlackClient = client
    return App
}

//vim: set ts=4 sw=4 et :
