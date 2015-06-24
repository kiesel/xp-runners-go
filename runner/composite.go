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
  uses := make([]string, 10)
  for _, member := range this.sources {
    uses = append(uses, member.GetUse()...)
  }

  return uses
}

func (this *CompositeConfigSource) GetRuntime() string {
  for _, member := range this.sources {
    if runtime := member.GetRuntime(); runtime != "" {
      return runtime
    }
  }

  panic("Cannot determine runtime from " + this.String())
}

func (this *CompositeConfigSource) GetExecutable(runtime string) string {
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