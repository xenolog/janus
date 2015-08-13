package slack

import (
    "fmt"
    slacklib "github.com/abourget/slack"
    "github.com/xenolog/janus/config"
    "github.com/xenolog/janus/logger"
    "sync"
)

type Slack struct {
    configured       bool
    eventLoopRunning bool
    slackConfig      *config.SlackConfig
    Api              *slacklib.Client
    Rtm              *slacklib.RTM
    rtmCall          *sync.Mutex
    apiCall          *sync.Mutex
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
        case ev := <-s.Rtm.IncomingEvents:
            log.Info("Event Received: %v", ev.Data)
            switch evt := ev.Data.(type) {
            case *slacklib.HelloEvent:
                // Ignore hello
                log.Info("Hello event: %v", ev.Data)

            case *slacklib.ConnectedEvent:
                log.Info("Infos:", evt.Info)
                log.Info("Connection counter: %d", evt.ConnectionCount)
                //s.Rtm.SendMessage(s.Rtm.NewOutgoingMessage("Hello world", "#general"))
                s.Rtm.SendMessage(s.Rtm.NewOutgoingMessage("Hello world", "C08RDQTFY"))

            case *slacklib.MessageEvent:
                log.Info("Message: %v", evt)
                // if private message given
                log.Info("Presence Change: %v", evt)

            case *slacklib.LatencyReport:
                //log.Info("Current latency: %v", evt.Value)

            case *slacklib.SlackWSError:
                log.Warn("Slack error: %d - %v", evt.Code, evt.Msg)

            default:
                // Ignore other events..
                log.Warn("Unexpected event: %v", ev.Data)
            }
        }
    }
}

func (s *Slack) addChannelToList() error {

    return nil
}

func (s *Slack) updateChannelList() error {
    s.ApiCall.Lock()
    defer s.ApiCall.Unlock()
    s.Api.Ge
    return nil
}

// Periodically update Channel-list
func (s *Slack) ChannelLoop() error {
    // get Channels from s.Rtm.GetInfo
    for {
        time.sleep(60)
        go s.updateChannelList()
    }
    return nil
}

func (s *Slack) MessageLoop() error {
    s.eventLoop()
    return nil
}

func (s *Slack) Connect() error {
    //todo: check for alredy connected
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
    return mainSlack
}

func init() {
    log = logger.New()
    mainSlack = new(Slack)
}
