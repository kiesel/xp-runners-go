package main 

import (
  "os"
  "github.com/kiesel/xp-runners-go/runner"
)

func main() {
  context := runner.Context {
    BaseDir : runner.Base(),
    Runner : "class",
    Includes : []string { "." },
  }
  parseArgs(&context, os.Args)

  runner.Execute(context)
}

func parseArgs(c *runner.Context, in []string) {
  if 1 == len(in) {
    c.Tool = "xp.runtime.ShowResource"
    c.Args = []string { "usage.txt", "255" }
    return
  }

  shift := 0
  skip := 0
  loop := in[1:]
ArgsLoop:
  for i, val := range loop {
    if skip > 0 {
      skip -= 1
      continue
    }

    switch val {
      case "-v":
        c.Tool = "xp.runtime.Version"
        shift += 1
        continue

      case "-e":
        c.Tool = "xp.runtime.Evaluate"
        shift += 1
        continue

      case "-w", "-d":
        c.Tool = "xp.runtime.Dump"
        continue

      case "-r":
        c.Tool = "xp.runtime.Reflect"
        shift += 1
        continue

      case "-xar":
        c.Tool = "xp.runtime.Xar"
        shift += 1
        continue

      case "-cp":
        c.Includes = append(c.Includes, loop[i + 1])
        shift += 2
        skip = 1
        continue

      default:
        if val[0] == '-' {
          panic("*** Invalid argument " + val)
        }

        break ArgsLoop
    }
  }

  c.Args = loop[shift:]
}