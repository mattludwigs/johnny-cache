package server

import (
  "fmt"
  "os"
)

func HandleError(err error, from string) {
  if err != nil {
    fmt.Printf("[%s]: %v", from, err)
    os.Exit(0)
  }
}
