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

  // maybe start TLS server
  go func() {
    tls_listener := os.Getenv("TLS")
    if tls_listener != "" {
      if err := http.ListenAndServeTLS(":5443", "/tls/tls.crt", "/tls/tls.key", nil); err != nil {
        log.Fatal(err)
      }
    }
  }()

  // Start HTTP server
  if err := http.ListenAndServe(":5000", nil); err != nil {
    log.Fatal(err)
  }
  
}
