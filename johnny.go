package main

import (
  "os"
  "os/signal"
  "syscall"
  "fmt"

  "github.com/mattludwigs/johnny-cache/server"
)

func main() {
  sigs := make(chan os.Signal)
  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
  go server.ServerRun(":2000")
   <- sigs
  fmt.Println("received signal; exiting")
}
