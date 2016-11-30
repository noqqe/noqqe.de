package main

import (
  "os"
  "os/exec"
  "log"
  "encoding/json"
  "strings"
  "time"
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

  log.Printf("%s\n", string(output))
}

// Create posts for sammelsurium using json export
// of rvo and persisting them into the posts
func sammelsurium(c Config) {

  // change directory
  os.Chdir(c.homedir)

  // convert cmd from string to array
  cmdline := strings.Split(c.rvocmd, " ")
  command := cmdline[0]
  args := cmdline[1:]
  cmd := exec.Command(command, args...)

  // check for errors
  output, err := cmd.CombinedOutput()
  if err != nil {
    log.Fatal("rvo crashed and burned.")
  }

  // convert to bytes and read json
  bytes := []byte(output)
  var documents []Document
  json.Unmarshal(bytes, &documents)

  // test output
  for d := range documents {
    convert_timestamp(documents[d].Created.Date)
    log.Printf("title: %s, created: %d", documents[d].Title, documents[d].Created.Date)
  }
}

// convert unix timestamp with 13 chars
// to human readable date
func convert_timestamp(ts int64) {
  corr_ts := ts / 1000
  tm := time.Unix(corr_ts, 0)
  log.Println(tm.Format(time0)
  return tm
}


func main() {

  // configuration is stored in this
  // tiny little struct made with <3
  c := Config{
    rvocmd: "rvo export -c docs",
    hugocmd: "hugo",
    homedir: "/home/noqqe/Code/noqqe.de"}

  log.Println("(Amaze. Wow)Oo.")

  if os.Args[1] == "build" {
    build(c)
  } else if os.Args[1] == "sammelsurium" {
    sammelsurium(c)
  }
}

