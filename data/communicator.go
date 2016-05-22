// interface and constants for communication plugins
package data

import (
//"fmt"
//"gopkg.in/yaml.v2"
//"sync"
)

type UniversalMsg struct {
    FromUid   string
    ToUid     []string // Message may be addresssed to dofferent recipients
    ChannelId string
    MsgId     string
}

type UniversalChannel struct {
    ChannelID string
    Public    bool   // True if this channel has public access
    ChannelId string // Channel-ID in transmitter-system format
    MsgId     string
}

type Communicator interface {
    Lock()               // Lock mutex
    Unlock()             // Free mutex
    Connect()            // connect to APIO
    Disconnect()         // disconnect from API
    StartRTM()           // start RealTimeMessaging loop
    SendMessage()        // send message by this way
    SetReceiveCallback() // setup goroutine, which will process incoming message this way
    UserListUpdate()     // fetch userlist from communicator
    AddUserEvent()       // goroutine for process 'new user' event
    ChannelListUpdate()  // fetch userlist from communicator
    AddChannelEvent()    // goroutine for process 'new channel' event
}
