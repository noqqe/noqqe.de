package main

import (
  "os"
  "os/exec"
  "fmt"
  "log"
  "github.com/antonholmquist/jason"
)

func build(home string) {
  os.Chdir(home)
  cmd := exec.Command("hugo")

  output, err := cmd.CombinedOutput()
  if err != nil {
    log.Fatal("Hugo crashed and burned.")
  }

  fmt.Printf("%s\n", string(output))
}

func sammelsurium(home string, destination string) {
  log.Println(destination)

  os.Chdir(home)
  cmd := exec.Command("rvo", "export", "-c", "docs", "-t", "OpenBSD")

  output, err := cmd.CombinedOutput()
  if err != nil {
    log.Fatal("rvo crashed and burned.")
  }

  // convert to bytes and read json
  bytes := []byte(output)
  documents, err := jason.NewObjectFromBytes(bytes)

  for _, friend := range documents {
    name, _ := friend.GetString("title")
    log.Println(name)
  }

  // for d := range documents {
  //   log.Printf("title: %s, created: %s", documents[d].Title, documents[d].Created)
  // }
}

func main() {
  log.Println("Amaze - Blogmanagement")
  var home string = "/home/noqqe/Code/noqqe.de/"
  // build(home)
  sammelsurium(home, "lol")
}

