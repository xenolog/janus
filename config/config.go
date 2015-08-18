package config

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os"
    "os/user"
    "regexp"
    "strings"
    "time"
)

type JanusConfig struct {
    // We used underscored field names here, because config stored in yaml format
    // and yaml parser required conformity between fields into yaml file and structure
    Slack struct {
        Api_token               string
        Nickname                string
        Public_channel_prefix   string
        Private_channel_prefix  string
        Direct_channel_prefix   string
        Channel_update_interval time.Duration
    }
}
type UserConfig struct {
    Slack struct {
        Nickname string
    }
    Irc struct {
        Username  string
        Password  string
        Nicknames []string
    }
}

type Config struct {
    path       string // given path to the config file
    raw_config []byte // temporary buffer for storing raw config
    //loaded from yaml
    Janus JanusConfig
    Users map[string]UserConfig
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
    cfg.path, _ = ExpandHomedir(path)
    err := cfg.reload()
    return cfg, err
}

func init() {
    cfg = new(Config)
}

// Expand '~'-based homedif from the given path
func ExpandHomedir(s string) (string, error) {
    const (
        slash = string(os.PathSeparator)
        re1   = "~%s"            // regex: /~\//
        re2   = "~([\\w\\-]+)%s" // regex: /~([\w\-]+)\//
    )
    var (
        err error
        re  *regexp.Regexp
        u   *user.User
        rv  string
    )

    if strings.HasPrefix(s, fmt.Sprintf(re1, slash)) {
        u, _ = user.Current()
        rv = fmt.Sprintf("%s", u.HomeDir+s[1:])
        err = nil
    } else if re = regexp.MustCompile(fmt.Sprintf(re2, slash)); re.MatchString(s) {
        uname := re.FindStringSubmatch(s)[0]
        uname = uname[1 : len(uname)-1]
        if u, _ = user.Lookup(uname); u == nil {
            rv = s
            err = nil
        } else {
            rv = u.HomeDir + slash + strings.Join(strings.Split(s, slash)[1:], slash)
            err = nil
        }
    } else if err != nil {
        rv = s
    } else {
        rv = s
        err = nil
    }
    return rv, err
}

//vim: set ts=4 sw=4 et :
