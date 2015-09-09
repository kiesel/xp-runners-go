package main 

import (
  "os"
  "github.com/kiesel/xp-runners-go/runner"
)

func main() {
  runner.Execute(runner.Context {
    BaseDir : runner.Base(),
    Runner : "class",
    Tool : "xp.command.Runner",
    Includes : []string { "." },
    Args : os.Args[1:],
  })
}