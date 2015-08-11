package config

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

type configType struct {
}

type Config struct {
    path       string // given path to the config file
    raw_config []byte // temporary buffer for storing raw config
    //loaded from yaml
    Janus struct {
        Slack_username  string
        Slack_api_token string
    }
    Users map[string]struct {
        Slack struct {
            Nickname string
        }
        Irc struct {
            Username  string
            Password  string
            Nicknames []string
        }
    }
}

var cfg *Config

// read config from yaml file and store in into intermediate beffer as raw
func (c *Config) read() error {
    var (
        err error
    )
    if c.raw_config, err = ioutil.ReadFile(c.path); err != nil {
        return fmt.Errorf("Can't read config '%s'", c.path)
    }
    return nil
}

// get config from intermediate buffer, parse it and store in *Config
func (c *Config) parse() error {
    var err error

    if len(c.raw_config) == 0 {
        return fmt.Errorf("Can't parse empty config")
    }

    if err = yaml.Unmarshal(c.raw_config, c); err != nil {
        return fmt.Errorf("Can't parse config file: %s", err)
    }
    c.raw_config = nil
    return nil
}

// load config from yaml file and parse it
func (c *Config) reload() error {
    if err := cfg.read(); err != nil {
        return err
    }
    if err := cfg.parse(); err != nil {
        return err
    }
    return nil
}

// return singletone config data structure
func New(path string) (*Config, error) {
    cfg.path = path
    err := cfg.reload()
    return cfg, err
}

func init() {
    cfg = new(Config)
}
