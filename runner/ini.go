package runner

import (
  "fmt"
  "strings"
  "reflect"
  "gopkg.in/ini.v1"
)

type IniConfigSource struct {
  filename string
  ini *ini.File
}

func (this *IniConfigSource) Valid() bool {
  return this.ini != nil
}

func (this *IniConfigSource) GetUse() []string {
  return this.ini.Section("").Key("use").Strings(";")
}

func (this *IniConfigSource) GetRuntime() string {
  return this.ini.Section("").Key("rt").String()
}

func (this *IniConfigSource) GetExecutable(runtime string) string {
  return this.ini.Section("runtime@" + runtime).Key("default").MustString(
    this.ini.Section("runtime").Key("default").String())
}

func (this *IniConfigSource) String() string {
  return reflect.TypeOf(this).String() + "{ filename: " + this.filename + " }"
}

func IniConfigSourceFromFile(paths ...string) *IniConfigSource {
  for _, element := range paths {

    // Given one path element is empty, return empty struct
    if "" == element {
      return &IniConfigSource{}
    }
  }

  path := strings.Join(paths, "/")

  ini, err := ini.Load(path)
  if err != nil {
    fmt.Println(err.Error())
    return &IniConfigSource{
      filename: path,
      ini: nil,
    }
  }

  instance := &IniConfigSource{
    filename: path,
    ini: ini,
  }

  return instance
}