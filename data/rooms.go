package data

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "sync"
)

const (
    ACCESS_DIRECT  = 1
    ACCESS_PRIVATE = 2
    ACCESS_GROUP   = 4
)

////////
// Chat room structure
type RoomT struct {
    sync.Mutex `yaml: mutex, omitempty`
    SlackName  string
    SLackID    string // should be in the Slack-specific format
    Access     int8   // may be ACCESS_[GROUP, PRIVATE, DIRECT]
    // IrcName   string
    // IrcId     string
}

func (r *RoomT) Update(id string, name string, access int8) {
    r.Lock()
    defer r.Unlock()
    if id != "" {
        r.SLackID = id
    }
    if name != "" {
        r.SlackName = name
    }
    if access != 0 {
        r.Access = access
    }
}

////////
// Rooms collection
type RoomsType struct {
    sync.Mutex
    ByName    map[string]*RoomT
    BySlackId map[string]*RoomT
    // IrcToName   map[string]*RoomT
    // NameToIrc   map[string]*RoomT
}

// Stringify for room collection
func (r *RoomsType) String() string {
    rv, _ := yaml.Marshal(r.BySlackId)
    return fmt.Sprintf("%s", rv)
}

// DeleteBySlackId -- remove room record with given SlackID
func (r *RoomsType) DeleteBySlackId(id string) error {
    r.Lock()
    defer r.Unlock()
    room, err := r.GetBySlackId(id)
    if err != nil {
        return err
    }
    delete(r.BySlackId, id)
    tmp := r.ByName[room.SlackName]
    if tmp != nil {
        delete(r.ByName, room.SlackName)
    }
    room = nil // emulate free(...)
    return nil
}

// GetBySlackId -- find room record with given SlackID and return pointer to it.
// Be carefully, use Lock()/Unlock() before and after use it for exclusive access
func (r *RoomsType) GetBySlackId(id string) (*RoomT, error) {
    // Locks should be used before call this method for avoiding deadlocks
    // r.Lock()
    // defer r.Unlock()
    room := r.BySlackId[id]
    if room == nil {
        return nil, fmt.Errorf("Can't find Room by ID '%s'.", id)
    }
    return room, nil
}

// PutBySlackId -- create new room record or modify existing with given SlackID
func (r *RoomsType) PutBySlackId(id string, name string, access int8) {
    var (
        room *RoomT
        err  error
    )
    r.Lock()
    defer r.Unlock()
    room, err = r.GetBySlackId(id)
    if err == nil {
        // update existing
        room.Update(id, name, access)
    } else {
        // create new
        room = new(RoomT)
        room.Update(id, name, access)
        r.ByName[name] = room
        r.BySlackId[id] = room
    }
}

// Make initialization. Should be used once.
func (r *RoomsType) init() {
    r.ByName = make(map[string]*RoomT)
    r.BySlackId = make(map[string]*RoomT)
}

var (
    Rooms *RoomsType
)

// NewRooms() -- return rooms storage. Rooms storage is a singletone,
// that created at start Janus.
func NewRooms() *RoomsType {
    return Rooms
}

////////
func init() {
    Rooms = new(RoomsType)
    Rooms.init()
}
