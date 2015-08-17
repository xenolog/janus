package config

import (
    "fmt"
    "os/user"
    "testing" //import go package for testing related functionality
)

const config_test_yaml = `
---
janus:
  slack:
    api_token: xxx-9034xxx3-vArCxxxaa3dm31xxxFgPH
    nickname: janus-bot
    public_channel_prefix: j_
    private_channel_prefix: j_
    direct_channel_prefix: d_
    channel_update_interval: 60
users:
  user1:
    slack:
      nickname: user_1
    irc:
      username: user_1
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
      username: user_2
      password: passwd2222
      nicknames:
        - helicopter
        - user_hel_1
        - user_hel_2
        - user_hel_3
`

// create new type for mock file read
type Config4test struct {
    Config
}

// mock reload() for get pre-defined data fixture
func (c *Config4test) reload() error {
    c.raw_config = []byte(config_test_yaml)
    // emulate parse()
    if err := c.Config.parse(); err != nil {
        return err
    }
    return nil
}

func TestParseConfig(t *testing.T) {
    var (
        cfg *Config4test
        err error
    )

    cfg = new(Config4test)
    cfg.path = "required_but_unused_string"

    if err = cfg.reload(); err != nil {
        t.Errorf("Yaml parse failed. '%s'", err)
    }

    if cfg.Janus.Slack.Nickname != "janus-bot" {
        t.Errorf("Yaml parse failed. '%s'", cfg.raw_config)
    }

    if cfg.Janus.Slack.Channel_update_interval != 60 {
        t.Errorf("Yaml parse failed. Channel_update_interval should be int")
    }

    if len(cfg.Users) != 2 {
        t.Errorf("Yaml parse failed. 'users' section not found.'%s'", cfg.raw_config)
    }
}

func TestExpandHomedir(t *testing.T) {
    var u *user.User
    path1 := "/aaa/bbb/ccc/ddd"
    if rv, err := ExpandHomedir(path1); err != nil {
        t.Errorf("Error while expanding homedir into '%s': %s", path1, err)
    } else if rv != path1 {
        t.Errorf("Wrong value returned while expanding homedir into '%s': %s", path1, rv)
    }
    path2 := "~/aaa/bbb/ccc/ddd"
    u, _ = user.Current()
    path2rv := fmt.Sprintf("%s/aaa/bbb/ccc/ddd", u.HomeDir)
    if rv, err := ExpandHomedir(path2); err != nil {
        t.Errorf("Error while expanding homedir into '%s': %s", path2, err)
    } else if rv != path2rv {
        t.Errorf("Wrong value returned while expanding homedir into '%s': %s", path2, rv)
    }
    path3 := "~username/aaa/bbb/ccc/ddd"
    if rv, err := ExpandHomedir(path3); err != nil {
        t.Errorf("Error while expanding homedir into '%s': %s", path3, err)
    } else if rv != path3 {
        t.Errorf("Wrong value returned while expanding homedir into '%s': %s", path3, rv)
    }
    // path4 := "~/aaa/./bbb/../bbb/ccc/ddd"
    // path4rv := fmt.Sprintf("%s/aaa/bbb/ccc/ddd", u.HomeDir)
    // if rv, err := ExpandHomedir(path4); err != nil {
    //     t.Errorf("Error while expanding homedir into '%s': %s", path4, err)
    // } else if rv != path4rv {
    //     t.Errorf("Wrong value returned while expanding homedir into '%s': %s", path4, rv)
    // }
}
