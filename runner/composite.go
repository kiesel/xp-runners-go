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

func NewComposite(sources ...ConfigSource) CompositeConfigSource {
  instance := CompositeConfigSource{}

  for _, add := range sources {
    if sources.Valid() {
      instance.sources = append(instance.sources, add)
    }
  }

  return instance
}