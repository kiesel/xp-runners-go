package runner

import (
  "testing"
  "reflect"
)

type Testing struct {
  *testing.T
}

func (this *Testing) equalSlices(a, b []string) {
  if len(a) != len(b) {
    this.Error("Difference slice sizes, expected same length!")
  }

  for position, value := range a {
    if value != b[position] {
      this.Errorf("Inequality at position %d (\"%s\" / \"%s\")", position, value, b[position])
    }
  }
}

type DummyConfig struct {
}

func (this *DummyConfig) Valid() bool {
  return true
}

func (this *DummyConfig) GetUse() []string {
  return []string {"/path/to/xp"}
}

func (this *DummyConfig) GetRuntime() string {
  return "default"
}

func (this *DummyConfig) GetExecutable(runtime string) string {
  return "php"
}

func (this *DummyConfig) GetArgs(runtime string) map[string]string {
  return make(map[string]string, 0)
}

func (this *DummyConfig) Locate(paths []string, entry string) string {
  return paths[0] + "/" + entry
}

func (this *DummyConfig) String() string {
  return reflect.TypeOf(this).String()
}

func TestBuildArgv(t *testing.T) {
  argv := buildArgv(
    new(DummyConfig),
    ".",
    "class",
    "xp.runtime.Version",
    []string {"."},
    []string {},
  )

  expect := []string {"-C", "-q", "-d", "include_path=\".:/path/to/xp::.\"", "-d", "magic_quotes_gpc=0", "/path/to/xp/tools/class.php", "xp.runtime.Version"}

  test := Testing{t}
  test.equalSlices(argv, expect)
}