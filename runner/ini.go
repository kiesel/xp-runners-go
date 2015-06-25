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

func (this *IniConfigSource) GetArgs(runtime string) map[string]string {
  args := make(map[string]string, 0)

  mergeWith(args, this.argsInSection("runtime@" + runtime))
  mergeWith(args, this.argsInSection("runtime"))

  return args
}

func (this *IniConfigSource) argsInSection(section string) map[string]string {
  args := make(map[string]string, 0)

  for _, key := range this.ini.Section(section).Keys() {
    if "default" == key.Name() || strings.HasPrefix(key.Name(), "extension.") {
      continue
    }

    args[key.Name()]= key.String()
  }

  return args
}

func mergeWith(args map[string]string, with map[string]string) {
  for key, value := range with {
    if _, ok := args[key]; ok == false {
      args[key]= value
    }
  }
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