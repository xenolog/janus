package config

import (
    "testing" //import go package for testing related functionality
)

const config = `
---
janus:
  slack_username: janus
  slack_api_token: token-aaaBBBccc
users:
  user1:
    slack:
      nickname: user_1
    irc:
      usename: user_1
      password: passwd1111
      nicknames:
        - user_1
        - user_113
        - user_11
        - user_12
  user2:
    slack:
      nickname: user_2
    irc:
      usename: user_2
      password: passwd2222
      nicknames:
        - helicopter
        - user_hel_1
        - user_hel_2
        - user_hel_3
`

func TestParseConfig(t *testing.T) {
    var config Config

    config =


    //     t.Error("GetLogger should return *log.Logger.")
    // case *log.Logger:
    //     t.Log("passed")
    // }
}

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
