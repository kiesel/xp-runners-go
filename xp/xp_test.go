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

  if len(c.Args) != 1 {
    t.Errorf("Invalid number of arguments: %d", len(c.Args))
  }

  if c.Args[0] != "var_dump('Hello World');" {
    t.Error("Invalid arguments.")
  }
}

func TestParseArgs_dump(t *testing.T) {
  c := runner.Context {}
  parseArgs(&c, []string {"xp.go", "-d", "true"})
  expectTool(c, "xp.runtime.Dump", t)
}

func TestParseArgs_dump_write(t *testing.T) {
  c := runner.Context {}
  parseArgs(&c, []string {"xp.go", "-w", "true"})
  expectTool(c, "xp.runtime.Dump", t)
}

func TestParseArgs_default_classpath(t *testing.T) {
  c := runner.Context {}
  parseArgs(&c, []string {"xp.go", "-e", "var_dump('Hello.');"})

  if 0 != len(c.Includes) {
    t.Errorf("Expected 0 include path, got: %d", len(c.Includes))
  }
}

func TestParseArgs_extended_classpath(t *testing.T) {
  c := runner.Context {}
  parseArgs(&c, []string {"xp.go", "-cp", "foo.xar", "-e", "var_dump('Hello.');"})

  if 1 != len(c.Includes) {
    t.Errorf("Expected 1 include path, got: %d", len(c.Includes))
  }
}

func TestParseArgs_extended_classpath_twice(t *testing.T) {
  c := runner.Context {}
  parseArgs(&c, []string {"xp.go", "-cp", "foo.xar", "-cp", "bar.xar", "-e", "var_dump(true);"})

  if 2 != len(c.Includes) {
    t.Error("Expected 2 include paths, got:", c.Includes)
  }

  expectTool(c, "xp.runtime.Evaluate", t)
  if 1 != len(c.Args) || c.Args[0] != "var_dump(true);" {
    t.Error("Invalid arguments, expected 'var_dump(true);', got:", c.Args)
  }
}

func TestParseArgs_invalid_command(t *testing.T) {
  c := runner.Context {}
  defer func() {
    if e := recover(); e != nil {

      if e != "*** Invalid argument -foo" {
        // Unexpected error recovered
        panic("Recovered unexpected error.")
      }

      // Otherwise, all good.
    }
  }()
  parseArgs(&c, []string {"xp.go", "-foo", "bar"})

  // Should not be reached:
  t.Error("Invalid command went through:(")
}