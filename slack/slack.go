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
            log.Debug("Event Received: %v", ev.Data)
            switch evt := ev.Data.(type) {
            case *slacklib.HelloEvent:
                // Ignore hello
                log.Debug("Hello event: %v", ev.Data)

            case *slacklib.ConnectedEvent:
                // init room structure by data, given at connection
                s.updateChannelList(&evt.Info.Channels)
                s.updateGroupList(&evt.Info.Groups)
                log.Debug("Rooms: \n%s", s.rooms)

                log.Debug("Connection counter: %d", evt.ConnectionCount)
                //s.Rtm.SendMessage(s.Rtm.NewOutgoingMessage("Hello world", "#general"))
                //s.Rtm.SendMessage(s.Rtm.NewOutgoingMessage("Hello world", "C08RDQTFY"))

            case *slacklib.MessageEvent:
                log.Debug("Message: %v", evt)
                // if private message given
                log.Debug("Presence Change: %v", evt)

            case *slacklib.LatencyReport:
                //log.Debug("Current latency: %v", evt.Value)

            case *slacklib.SlackWSError:
                log.Warn("Slack error: %d - %v", evt.Code, evt.Msg)

            default:
                // Ignore other events..
                log.Debug("Unexpected event: %v", ev.Data)
            }
        }
    }
}

func (s *Slack) updateGroupList(groups *[]slacklib.Group) {
    for _, gr := range *groups {
        s.rooms.PutBySlackId(gr.Id, gr.Name, data.ACCESS_PRIVATE)
    }
}

func (s *Slack) updateChannelList(channels *[]slacklib.Channel) {
    for _, ch := range *channels {
        s.rooms.PutBySlackId(ch.Id, ch.Name, data.ACCESS_GROUP)
    }
}

// fetch channels and groups from
func (s *Slack) fetchChannelList() error {
    s.apiCall.Lock()
    chs, errC := s.Api.GetChannels(true)
    grs, errG := s.Api.GetGroups(true)
    s.apiCall.Unlock()
    if errC == nil {
        s.updateChannelList(&chs)
    }
    if errG == nil {
        s.updateGroupList(&grs)
    }
    log.Debug("Rooms: \n%s", s.rooms)

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
        go s.fetchChannelList()
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
    // start RTM main loop
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
