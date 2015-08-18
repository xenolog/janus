package slack

import (
    "fmt"
    slacklib "github.com/abourget/slack"
    "github.com/xenolog/janus/config"
    "github.com/xenolog/janus/data"
    "github.com/xenolog/janus/logger"
    "sync"
    "time"
)

///
type Slack struct {
    configured       bool
    eventLoopRunning bool
    janusConfig      *config.JanusConfig
    Api              *slacklib.Client
    Rtm              *slacklib.RTM
    rtmCall          sync.Mutex
    apiCall          sync.Mutex
    rooms            *data.RoomsType
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

func (s *Slack) updateChannelList() error {
    s.apiCall.Lock()
    chs, errC := s.Api.GetChannels(true)
    grs, errG := s.Api.GetGroups(true)
    s.apiCall.Unlock()
    if errC == nil {
        for _, ch := range chs {
            s.rooms.PutBySlackId(ch.Id, ch.Name, 'G')
        }
    }
    if errG == nil {
        for _, gr := range grs {
            s.rooms.PutBySlackId(gr.Id, gr.Name, 'P')
        }
    }
    log.Info("Rooms: %v", s.rooms)

    if errC != nil {
        return errC
    } else if errG != nil {
        return errG
    }
    return nil
}

// Periodically update Channel-list
func (s *Slack) ChannelLoop() error {
    // get Channels from s.Rtm.GetInfo
    for {
        time.Sleep(s.janusConfig.Slack.Channel_update_interval * time.Second)
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
    s.Api = slacklib.New(s.janusConfig.Slack.Api_token)
    s.Api.SetDebug(true)
    s.Rtm = s.Api.NewRTM()
    go mainSlack.Rtm.ManageConnection()
    return nil
}

func (s *Slack) init() {
    s.rooms = data.NewRooms()
}

func New(config *config.JanusConfig) *Slack {
    if !mainSlack.configured {
        mainSlack.janusConfig = config
        mainSlack.configured = true
    }
    return mainSlack
}

func init() {
    log = logger.New()
    mainSlack = new(Slack)
    mainSlack.init()
}

//vim: set ts=4 sw=4 et :
