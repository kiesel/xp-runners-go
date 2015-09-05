package main 

import (
  "os"
  "path/filepath"
  "fmt"
  "log"
  "github.com/kiesel/xp-runners-go/runner"
)

func dir(path string) (string, error) {
  path, err := filepath.EvalSymlinks(path)

  if err != nil {
    return "", err
  }

  path, err = filepath.Abs(path)
  if err != nil {
    return "", err
  }

  return filepath.Dir(path), nil
}

func main() {
  if dbg := os.Getenv("XP_DBG"); dbg != "" {
    log.SetFlags(log.Lshortfile)
    log.SetOutput(os.Stdout)
  }

  base, err := dir(os.Args[0])
  if err != nil {
    panic(err.Error())
  }

  runner.Execute(base, "class", "xp.runtime.Version", []string{ "." }, []string{})
}