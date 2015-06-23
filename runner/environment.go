package runner

type EnvironmentConfigSource struct {
}

func (this *EnvironmentConfigSource) Valid() bool {
  return true
}

func (this *EnvironmentConfigSource) GetUse() []string {
  return []string{}
}