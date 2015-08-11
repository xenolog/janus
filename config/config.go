package config

import (
    "github.com/xenolog/janus/logger"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type Config struct {
    path       string // given path to the config file
    raw_config []byte // temporary buffer for storing raw config
    C          struct {
        janus struct {
            slack_username  string
            slack_api_token string
        }
        users map[string]struct {
            slack struct {
                nickname string
            }
            irc struct {
                username  string
                password  string
                nicknames []string
            }
        }
    }
}

var (
    ll  *log.Logger
    cfg Config
)

// func (c *Config) Users() *UserConfig {
//     return c.Config.users
// }

func (c *Config) read() error {
    var err error
    if c.raw_config, err = ioutil.ReadFile(c.path); err != nil {
        return err
    }
    return nil
}

func (c *Config) parse() error {
    var err error

    if len(c.raw_config) == 0 {
        ll.Printf("Can't parse empty config")
        return err //error{"xxx"}
    }
    //ll.Printf("X: %s", c.raw_config)

    if err = yaml.Unmarshal(c.raw_config, &c.C); err != nil {
        ll.Printf("Can't parse config file: %s", err)
        return err
    }
    c.raw_config = nil
    ll.Printf("C: %s", c.C)
    return nil
}

func (c *Config) reload() error {
    if err := cfg.read(); err != nil {
        return err
    }
    if err := cfg.parse(); err != nil {
        return err
    }
    return nil
}

func New(path string) (*Config, error) {
    cfg.path = path
    err := cfg.reload()
    return &cfg, err
}

func init() {
    ll = logger.GetLogger()
    cfg = Config{} // &Config{path: path}
}
