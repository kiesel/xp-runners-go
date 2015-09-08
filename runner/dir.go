package runner 

import (
  "os"
  "path/filepath"
)

func RealDir(path string) (string, error) {
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

func Base() (string) {
  base, err := RealDir(os.Args[0])

  if err != nil {
    panic(err.Error())
  }

  return base
}

