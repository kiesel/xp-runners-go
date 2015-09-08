package main 

import (
  "os"
  "github.com/kiesel/xp-runners-go/runner"
)

func main() {
  runner.Execute(runner.Base(), "class", "xp.command.Runner", []string{ "." }, os.Args[1:])
}