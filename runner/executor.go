package runner 

import (
  "fmt"
  "strings"
  "os"
  "path/filepath"
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
    "-d", fmt.Sprintf("include_path=\".%[2]s%[1]s%[2]s%[2]s%[3]s\"",
      strings.Join(useXp, string(os.PathListSeparator)),
      string(os.PathListSeparator),
      strings.Join(includes, string(os.PathListSeparator)),
    ),
    "-d", "magic_quotes_gpc=0",
  }

  for key, value := range configs.GetArgs(runtime) {
    argv = append(argv, "-d", key + "=\"" + value + "\"")
  }

  var runnerPath string
  if runnerPath = locate(useXp, "tools/" + runner + ".php"); runnerPath != "" {
    // noop
  } else if runnerPath = locate([]string { baseDir }, runner + "-main.php"); runnerPath != "" {
    argv = append(argv, "-d", "encoding=\"utf-7\"")
  } else {
    panic("Cannot find tool in " + strings.Join(useXp, string(os.PathListSeparator)))
  }

  fmt.Println("argv   := ", argv)
  fmt.Println("runt   := ", runtime)
  fmt.Println("exec   := ", executable)
  fmt.Println("usexp  := ", useXp)
  fmt.Println("runner := ", runnerPath)

  // cmd := exec.Command("")

  // return cmd
}

func locate(paths []string, entry string) string {
  for _, path := range paths {
    abs := filepath.Join(path, entry)
    fmt.Println("Checking", abs)
    stat, err := os.Stat(abs)
    if err != nil {
      continue
    }

    if !stat.IsDir() {
      return abs
    }
  }

  return ""
}