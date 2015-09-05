package main 

import (
  "os"
  "path/filepath"
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

  var tool string
  var args []string

  shift := 0
  includes := []string { "." }

  if 1 == len(os.Args) {
    tool = "xp.runtime.ShowResource"
    args = []string { "usage.txt", "255" }
  } else {
ArgsLoop:
    for i, val := range os.Args {
      if 0 == i {
        continue
      }

      switch val {
        case "-v":
          tool = "xp.runtime.Version"
          shift += 1
          break

        case "-e":
          tool = "xp.runtime.Evaluate"
          shift += 1
          break

        case "-w", "-d":
          tool = "xp.runtime.Dump"
          break

        case "-r":
          tool = "xp.runtime.Reflect"
          shift += 1
          break

        case "-xar":
          tool = "xp.runtime.Xar"
          shift += 1
          break

        case "-cp":
          includes = append(includes, os.Args[i + 1])
          shift += 2
          break

        default:
          if val[0] == '-' {
            panic("*** Invalid argument " + val)
          }

          break ArgsLoop
      }
    }

    args = os.Args[shift:]
  }

  runner.Execute(base, "class", tool, includes, args)
}