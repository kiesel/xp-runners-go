package runner

import (
  "strings"
  "gopkg.in/ini.v1"
)

type IniConfigSource struct {
  filename string
  ini *ini.File
}

func (this *IniConfigSource) Valid() bool {
  return true
}

func (this *IniConfigSource) GetUse() []string {
  return this.ini.Section("").Key("use").Strings(";")
}

func (this *IniConfigSource) GetExecutable() string {
  return this.ini.Section("").Key("rt").String()
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
    panic("Cannot load ini from " + path)
  }

  instance := &IniConfigSource{
    filename: path,
    ini: ini,
  }

  return instance
}