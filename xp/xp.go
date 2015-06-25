package main 

import (
  "os"
  "path/filepath"
  "fmt"
  "github.com/xp-framework/xp-runners-go/runner"
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
  base, err := dir(os.Args[0])
  if err != nil {
    panic(err.Error())
  }

  fmt.Println(base)
  runner.Execute(base, "class", "xp.runtime.Version", []string{ "." }, []string{})
}