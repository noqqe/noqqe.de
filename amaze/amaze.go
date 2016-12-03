package main

import (
  "log"
  "os"
  "os/exec"
  "strings"
  "encoding/json"
)


// Tiny little configuration struct
type Config struct {
  hugocmd string
  homedir string
  sammelsuriumdir string
  rvocmd string
}

// A document coming from rvo
// looks like this.
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

func build(c Config) bool {
  os.Chdir(c.homedir)
  cmd := exec.Command(c.hugocmd)

  output, err := cmd.CombinedOutput()
  if err != nil {
    log.Fatal("Hugo crashed and burned.")
  }

  log.Printf("%s\n", string(output))

  return true
}

// Create posts for sammelsurium using json export
// of rvo and persisting them into the posts
func sammelsurium(c Config) bool {

  // Log message
  log.Println("Sammelsurium calculating...")

  // change directory
  os.Chdir(c.homedir)

  // convert cmd from string to array
  cmdline := strings.Split(c.rvocmd, " ")
  command := cmdline[0]
  args := cmdline[1:]

  // execute commands
  cmd := exec.Command(command, args...)
  output, err := cmd.CombinedOutput()

  // check for errors
  if err != nil {
    log.Fatal("rvo crashed and burned.")
  }

  // convert to bytes and read json
  bytes := []byte(output)
  var documents []Document
  json.Unmarshal(bytes, &documents)

  log.Printf("Entries found: %d", len(documents))

  for d := range documents {

		// dont parse if it has a tag called "private"
		if stringInSlice("private", documents[d].Tags) {
			continue
		}

    // get filename
    filename := generateFilename(documents[d].Title)
    path := c.sammelsuriumdir + "/" + filename

    // strip content
    content := generateContent(documents[d].Content)

    // generate hugo comptabile timestamp
    date := generateDate(documents[d].Created.Date)

    // generate header
    header := generateHeader(documents[d].Title, date, documents[d].Tags)

    // write sammelsurium post to hugo dir
    createSammelsuriumFile(path, header, content)

    // print debugging
    log.Printf("Post: %s (%s) written to %s", documents[d].Title, date, path)
  }

  return true
}

func main() {

  // configuration is stored in this
  // tiny little struct made with <3
  c := Config{
    rvocmd: "rvo export -c docs",
    hugocmd: "hugo",
    sammelsuriumdir: "/home/noqqe/Code/noqqe.de/content/sammelsurium/",
    homedir: "/home/noqqe/Code/noqqe.de"}

  // Some warm welcome
  log.Println(".oO(Amaze - Wow)Oo.")

  if os.Args[1] == "build" {
    build(c)
  } else if os.Args[1] == "sammelsurium" {
    sammelsurium(c)
  }
}

