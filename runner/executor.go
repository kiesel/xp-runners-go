package runner 

import (
  "os"
  "os/exec"
)

type Executor struct {
}

func Execute(baseDir, runner, tool string, includes, args []string) int {
  proc := newProcess(baseDir, runner, tool, includes, args)
  proc.Run()

  return 1
}

func newProcess(baseDir, runner, tool string, includes, args []string) *exec.Cmd {
  configs := NewCompositeConfigSource(
    NewEnvironmentConfigSource(),
    IniConfigSourceFromFile(".", "xp.ini"),
    IniConfigSourceFromFile(os.Getenv("HOME"), ".xp.ini"),
    IniConfigSourceFromFile(baseDir, "xp.ini"),
  )

  useXp := configs.GetUse()
  executable := configs.GetExecutable()
  // runtime := configs.GetRuntime()

  cmd := exec.Command("")

  return cmd
}