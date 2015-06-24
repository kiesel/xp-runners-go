package runner

import (
  "os"
  "strings"
  "reflect"
)

type EnvironmentConfigSource struct {
}

func (this *EnvironmentConfigSource) Valid() bool {
  return true
}

func (this *EnvironmentConfigSource) GetUse() []string {
  use := make([]string, 10)

  wd, err := os.Getwd()
  if err != nil {
    panic(err.Error())
  }
  use = append(use, wd)

  for _, element := range strings.Split(os.Getenv("USE_XP"), ";") {
    use = append(use, element)
  }

  return use
}

func (this *EnvironmentConfigSource) GetRuntime() string {
  return os.Getenv("XP_RT")
}

func (this *EnvironmentConfigSource) GetExecutable(runtime string) string {
  return ""
}

func (this *EnvironmentConfigSource) String() string {
  return reflect.TypeOf(this).String() + "{}"
}

func NewEnvironmentConfigSource() *EnvironmentConfigSource {
  return &EnvironmentConfigSource{}
}