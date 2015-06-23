package runner

type ConfigSource interface {
  Valid() bool
  GetUse() []string
  GetExecutable() string
}