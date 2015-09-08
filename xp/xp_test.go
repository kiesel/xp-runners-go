package main 

import (
  "testing"
  "testing/quick"
)

func TestparseArgs_version(t *testing.T) {
  c = runner.Context {}
  parseArgs(&c, []string {"xp.go", "-v"})


}