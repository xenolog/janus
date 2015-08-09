// entry point to the Janus ()
package config

import (
    "github.com/xenolog/janus/logger"
    "log"
)

var (
    ll *log.Logger
)

func init() {
    ll = logger.GetLogger()
}
