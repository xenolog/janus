package data

import (
    "fmt"
    "sync"
)

////////
// Chat room structure
type RoomT struct {
    sync.Mutex
    SlackName string
    SLackID   string
    Access    byte // may be 'Group', 'Private', 'Direct'
    // IrcName   string
    // IrcId     string
}

func (r *RoomT) Update(id string, name string, access byte) {
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
    ByName        map[string]*RoomT
    BySlackId     map[string]*RoomT
    AllowedAccess map[byte]bool
    // IrcToName   map[string]*RoomT
    // NameToIrc   map[string]*RoomT
}

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

func (r *RoomsType) CreateOrUpdateRoom(id string, name string, access byte) {
    var (
        room *RoomT
        aaa  byte
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
        if r.AllowedAccess[access] {
            aaa = access
        } else {
            aaa = 0
        }
        room.Update(id, name, aaa)
        r.ByName[name] = room
        r.BySlackId[id] = room
    }
}

func (r *RoomsType) init() {
    r.AllowedAccess = map[byte]bool{'G': true, 'P': true, 'D': true}
    r.ByName = make(map[string]*RoomT)
    r.BySlackId = make(map[string]*RoomT)
}

var (
    Rooms *RoomsType
)

func NewRooms() *RoomsType {
    return Rooms
}

func init() {
    Rooms = new(RoomsType)
    Rooms.init()
}
