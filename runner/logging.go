package runner

import (
  "os"
  log "github.com/kdar/factorlog"
)

var (
  Log log.Logger
)

func init() {
  Log = new (log.NullLogger)
  if dbg := os.Getenv("XP_DBG"); dbg != "" {
    Log = log.New(os.Stdout, log.NewStdFormatter("%{File}:%{Line} >> %{Message}"))
  }
}