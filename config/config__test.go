package config

import (
    "log"
    "testing" //import go package for testing related functionality
)

// func TestGetLoggerReturnType(t *testing.T) {
//     var xlogger interface{}
//     xlogger = GetLogger()
//     switch xlogger.(type) {
//     default:
//         t.Error("GetLogger should return *log.Logger.")
//     case *log.Logger:
//         t.Log("passed")
//     }
// }

// func TestGetLoggerSingletone(t *testing.T) {
//     var l1, l2 *log.Logger
//     l1 = GetLogger()
//     l2 = GetLogger()
//     if l1 != l2 {
//         t.Error("objects, returned by GetLogger is not an singletone.")
//     } else {
//         t.Log("passed")
//     }
// }
