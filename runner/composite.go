package runner

import (
  "reflect"
)

type CompositeConfigSource struct {
  sources []ConfigSource
}

func (this *CompositeConfigSource) Valid() bool {
  return true
}

func (this *CompositeConfigSource) GetUse() []string {
  uses := []string{}
  for _, member := range this.sources {
    for _, element := range member.GetUse() {
      if element != "" {
        uses = append(uses, element)
      }
    }
  }

  return uses
}

func (this *CompositeConfigSource) GetRuntime() string {
  for _, member := range this.sources {
    if runtime := member.GetRuntime(); runtime != "" {
      return runtime
    }
  }

  return "default"
}

func (this *CompositeConfigSource) GetExecutable(runtime string) string {
  for _, member := range this.sources {
    if executable := member.GetExecutable(runtime); executable != "" {
      return executable
    }
  }

  return "php"
}

func (this *CompositeConfigSource) GetArgs(runtime string) map[string]string {
  args := make(map[string]string, 0)
  for _, member := range this.sources {
    for key, value := range member.GetArgs(runtime) {
      if _, ok := args[key]; ok == false {
        args[key]= value
      }
    }
  }

  return args
}

func (this *CompositeConfigSource) Locate(paths []string, entry string) string {
  for _, member := range this.sources {
    if location := member.Locate(paths, entry); location != "" {
      return location
    }
  }

  return ""
}

func (this *CompositeConfigSource) String() string {
  val := reflect.TypeOf(this).String() + "{\n"
  for _, elem := range this.sources {
    val += "  " + elem.String() + "\n"
  }
  
  return val + "}"
}

func NewCompositeConfigSource(sources ...ConfigSource) ConfigSource {
  instance := &CompositeConfigSource{}

  for _, add := range sources {
    if add.Valid() {
      instance.sources = append(instance.sources, add)
    }
  }

  return instance
}