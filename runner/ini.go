package runner

type IniConfigSource struct {
  filename string
}

func (this *IniConfigSource) Valid() bool {
  return true
}

func (this *IniConfigSource) GetUse() []string {
  return []string{}
}

func NewIniConfigSource(filename string) IniConfigSource {
  instance := IniConfigSource{
    filename: filename
  }

  return instance
}