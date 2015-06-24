package runner 

import (
  "fmt"
  "strings"
  "os"
  // "os/exec"
)

type Executor struct {
}

func Execute(baseDir, runner, tool string, includes, args []string) int {
  newProcess(baseDir, runner, tool, includes, args)
  // proc.Run()

  return 1
}

func newProcess(baseDir, runner, tool string, includes, args []string) {
  configs := NewCompositeConfigSource(
    NewEnvironmentConfigSource(),
    IniConfigSourceFromFile(".", "xp.ini"),
    IniConfigSourceFromFile(os.Getenv("HOME"), ".xp.ini"),
    IniConfigSourceFromFile(baseDir, "xp.ini"),
  )

  useXp := configs.GetUse()
  runtime := configs.GetRuntime()
  executable := configs.GetExecutable(runtime)

  if 0 == len(useXp) {
    panic("Cannot determine use_xp setting from " + configs.String())
  }

  argv := []string {
    "-C", "-q", 
    "-d", fmt.Sprintf("include_path=\"%[1]s%[0]s%[1]s%[1]s%[2]s\"",
      strings.Join(useXp, ";"),
      ";",
      strings.Join(includes, ";"),
    ),
    "-d", "magic_quotes_gpc=0",
  }

  fmt.Println("argv := ", argv)
  fmt.Println("runt := ", runtime)
  fmt.Println("exec := ", executable)

  // cmd := exec.Command("")

  // return cmd
}