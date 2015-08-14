package data

import (
    "testing"
)

func TestCreateAndGetRoom(t *testing.T) {
    r := NewRooms()
    r.CreateOrUpdateRoom("P540215", "Test_room_1", 'G')
    r.CreateOrUpdateRoom("P675434", "Test_room_2", 'G')
    m, err := r.GetBySlackId("P675434")
    if err != nil {
        t.Fatalf("Can't get Room by ID: %v", err)
    }
    if m.SlackName != "Test_room_2" {
        t.Error("Wrong Room was got by ID.")
    } else {
        t.Log("Pass")
    }
}

func TestGettingNonExistingRoom(t *testing.T) {
    r := NewRooms()
    r.CreateOrUpdateRoom("P540215", "Test_room_1", 'G')
    r.CreateOrUpdateRoom("P675434", "Test_room_2", 'G')
    _, err := r.GetBySlackId("XXXXXXX")
    if err != nil {
        t.Log("Pass")
    } else {
        t.Error("Found non-existing Room by ID")
    }
}

func TestDeleteExistingRoom(t *testing.T) {
    r := NewRooms()
    r.CreateOrUpdateRoom("P675434", "Test_room_2", 'G')
    err := r.DeleteBySlackId("P675434")
    if err != nil {
        t.Error("Can't delete Room by ID: %v", err)
    } else {
        t.Log("Pass")
    }
}

func TestDeleteNonExistingRoom(t *testing.T) {
    r := NewRooms()
    r.CreateOrUpdateRoom("P675434", "Test_room_2", 'G')
    err := r.DeleteBySlackId("XXXXXXX")
    if err != nil {
        t.Log("Pass")
    } else {
        t.Error("Try to delete Room by non-existing ID was successful")
    }
}
