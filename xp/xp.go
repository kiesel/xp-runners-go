package main 

import (
  "github.com/xp-framework/xp-runners-go/runner"
)

func main() {
  runner.Execute(".", "class", "xp.runtime.Version", []string{ "." }, []string{})
}