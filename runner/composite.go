package runner

type CompositeConfigSource struct {
  sources []ConfigSource
}

func (this *CompositeConfigSource) Valid() bool {
  return true
}

func (this *CompositeConfigSource) GetUse() []string {
  return []string{}
}

func (this *CompositeConfigSource) GetExecutable() string {
  return ""
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