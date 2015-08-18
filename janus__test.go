package main

import (
    "github.com/xenolog/janus/logger"
    "log"
    "testing" //import go package for testing related functionality
)

func TestGetLogger(t *testing.T) {
    var xlogger interface{}
    xlogger = logger.GetLogger()
    switch xlogger.(type) {
    default:
        t.Error("Something wring was been imported instead logger.")
    case *log.Logger:
        t.Log("passed")
    }
}

//vim: set ts=4 sw=4 et :
