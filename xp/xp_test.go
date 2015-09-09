package main 

import (
  "testing"
  "github.com/kiesel/xp-runners-go/runner"
)

func expectTool(c runner.Context, actual string, t *testing.T) {
  if (c.Tool != actual) {
    t.Errorf("Unexpected tool '%s', was looking for '%s'", c.Tool, actual)
  }
}

func TestParseArgs_noargs(t *testing.T) {
  c := runner.Context {}
  parseArgs(&c, []string {"xp.go"})

  expectTool(c, "xp.runtime.ShowResource", t)
  if (len(c.Args) != 2) {
    t.Error("Invalid number of arguments.")
  }
}

func TestParseArgs_version(t *testing.T) {
  c := runner.Context {}
  parseArgs(&c, []string {"xp.go", "-v"})
  expectTool(c, "xp.runtime.Version", t)
}

func TestParseArgs_evaluate(t *testing.T) {
  c := runner.Context {}
  parseArgs(&c, []string {"xp.go", "-e", "var_dump('Hello World');"})
  expectTool(c, "xp.runtime.Evaluate", t)
}