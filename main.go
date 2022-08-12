package main

import (
  "log"
  "os"
  "net/http"
  "github.com/sosedoff/gitkit"
)

const GitDir = "/git/repos"

func main() {
  if err := os.MkdirAll(GitDir, 0o755); err != nil {
    log.Fatal("error creating git dir: %v", err)
  }
  // Configure git service
  service := gitkit.New(gitkit.Config{
    Dir:        GitDir,
    AutoCreate: true,
  })

  // Configure git server. Will create git repos path if it does not exist.
  // If hooks are set, it will also update all repos with new version of hook scripts.
  if err := service.Setup(); err != nil {
    log.Fatal(err)
  }

  http.Handle("/", service)

  // Start HTTP server
  if err := http.ListenAndServe(":5000", nil); err != nil {
    log.Fatal(err)
  }
}
