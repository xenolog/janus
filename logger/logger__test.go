package logger

import (
    "log"
    "testing" //import go package for testing related functionality
)

var (
    result *log.Logger
)

func TestGetLoggerReturnType(t *testing.T) {
    var xlogger interface{}
    xlogger = GetLogger()
    switch xlogger.(type) {
    default:
        t.Error("GetLogger should return *log.Logger.")
    case *log.Logger:
        t.Log("passed")
    }
}

func TestGetLoggerSingletone(t *testing.T) {
    var l1, l2 *log.Logger
    l1 = GetLogger()
    l2 = GetLogger()
    if l1 != l2 {
        t.Error("objects, returned by GetLogger is not an singletone.")
    } else {
        t.Log("passed")
    }
}

func BenchmarkGetLogger(b *testing.B) {
    var l *log.Logger
    for i := 0; i < b.N; i++ {
        l = GetLogger()
    }
    // this trick need for avoid compiler optimization
    result = l
}

//vim: set ts=4 sw=4 et :
