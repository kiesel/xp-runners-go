package runner

type ConfigSource interface {
  Valid() bool
  GetUse() []string
  GetRuntime() string
  GetExecutable(runtime string) string
  String() string
}