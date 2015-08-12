package slack

import (
    "fmt"
    slacklib "github.com/abourget/slack"
    "github.com/xenolog/janus/config"
    "github.com/xenolog/janus/logger"
)

type Slack struct {
    configured       bool
    eventLoopRunning bool
    slackConfig      *config.SlackConfig
    Api              *slacklib.Client
    Rtm              *slacklib.RTM
}

var (
    mainSlack *Slack
    log       *logger.Logger
)

func (s *Slack) eventLoop() error {
    if s.eventLoopRunning {
        return fmt.Errorf("Event loop already running")
    }
    for {
        select {
        case msg := <-s.Rtm.IncomingEvents:
            log.Info("Event Received: ")
            switch ev := msg.Data.(type) {
            case *slacklib.HelloEvent:
                // Ignore hello
                log.Info("Hello event: %s", msg.Data)

            case *slacklib.ConnectedEvent:
                log.Info("Infos:", ev.Info)
                log.Info("Connection counter: %d", ev.ConnectionCount)
                //s.Rtm.SendMessage(s.Rtm.NewOutgoingMessage("Hello world", "#general"))
                s.Rtm.SendMessage(s.Rtm.NewOutgoingMessage("Hello world", "C08RDQTFY"))

            case *slacklib.MessageEvent:
                log.Info("Message: %v", ev)

            case *slacklib.PresenceChangeEvent:
                log.Info("Presence Change: %v", ev)

            case *slacklib.LatencyReport:
                log.Info("Current latency: %v", ev.Value)

            case *slacklib.SlackWSError:
                log.Warn("Slack error: %d - %s", ev.Code, ev.Msg)

            default:
                // Ignore other events..
                log.Warn("Unexpected event: %v", msg.Data)
            }
        }
    }
}

func (s *Slack) MainLoop() error {
    go s.eventLoop()
    return nil
}

func (s *Slack) Connect() error {
    s.Api = slacklib.New(s.slackConfig.Slack_api_token)
    s.Api.SetDebug(true)
    s.Rtm = s.Api.NewRTM()
    go mainSlack.Rtm.ManageConnection()
    return nil
}

func New(config *config.SlackConfig) *Slack {
    if !mainSlack.configured {
        mainSlack.slackConfig = config
        mainSlack.configured = true
    }
    //todo: Also may be used form:
    // log.SetOutput(io.MultiWriter(os.Stdout, logFile))
    return mainSlack
}

func init() {
    log = logger.New()
    mainSlack = new(Slack)
}
