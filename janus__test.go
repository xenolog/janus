package main

import (
	"gopkg.in/xenolog/go-tiny-logger.v1"
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
