package main

import (
  "os"
  "os/exec"
  "fmt"
  "log"
  "encoding/json"
)

type Config struct {
  hugocmd string
  homedir string
  rvocmd string
}

type Document struct {
  Updated struct {
    Date int64 `json:"$date"`
  } `json:"updated"`
  Created struct {
    Date int64 `json:"$date"`
  } `json:"created"`
  Encrypted bool `json:"encrypted"`
  Tags []string `json:"tags"`
  Content string `json:"content"`
  Title string `json:"title"`
  ID struct {
    Oid string `json:"$oid"`
  } `json:"_id"`
  Categories []string `json:"categories"`
}

func build(c Config) {
  os.Chdir(c.homedir)
  cmd := exec.Command(c.hugocmd)

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
  var documents []Document
  json.Unmarshal(bytes, &documents)

  for d := range documents {
    log.Printf("title: %s, created: %d", documents[d].Title, documents[d].Created.Date)
  }
}

func main() {
  c := Config{
    rvocmd: "rvo export -c docs",
    hugocmd: "hugo",
    homedir: "/home/noqqe/Code/noqqe.de"}

  log.Println("Amaze - Blogmanagement")
  build(c)
  sammelsurium(c.homedir, "lol")
}

