package runner

type ConfigSource interface {
  Valid() bool
  GetUse() []string
  GetRuntime() string
  GetExecutable(runtime string) string
  GetArgs(runtime string) map[string]string
  String() string
}