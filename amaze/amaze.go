package main

import (
  "log"
  "os"
  "os/exec"
  "strings"
  "time"
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

// used to write content as a post to hugo sammelsurium directory
func createSammelsuriumFile(path string, header string, content string) bool {

	// create file
  f, err := os.Create(path)
  if err != nil {
      panic(err)
  }

  // convert header string to bytes and write to file opened
  h := []byte(header)
	f.Write(h)

  // convert header string to bytes and write to file opened
  b := []byte(content)
	f.Write(b)

	f.Close()

  return true
}

// Helper to make the filename out of the title
func generateFilename(title string) string {
  name := strings.ToLower(title)
  name = strings.Replace(name, " ", "-", -1)
  name = strings.Replace(name, "/", "", -1)
  name = name + ".md"
  return name
}

func generateContent(content string) string {

  // get rid of the first line of content because
  // it only contains the headline which is not needed
  // in the content anymore.
  c := strings.Split(content,"\n")
  c = c[1:]
  cs := strings.Join(c[:],"\n")

  // replace h1 with h2 for layouting
  cs = strings.Replace(cs, "# ", "##", -1)
  return cs
}

// convert unix timestamp with 13 chars
// to human readable date
func generateDate(ts int64) string {
  corr_ts := ts / 1000
  date := time.Unix(corr_ts, 0).Format("2006-01-02 15:04:05")
  return date
}

// builds hugo header in yaml format
func generateHeader(title string, date string, tags []string) string {
  limit := "---"
  n := "\n"

  // head
  header := limit + n

  // title
  header = header + "title: " + title + n

  // date
  header = header + "date: " + date + n

  // tags
  if len(tags) > 0 {
    header = header + "tags: " + n
    for _, tag := range tags {
      header = header + "- " + tag + n
    }
  }

  // foot
  header = header + limit + n

  return header
}

// check if slice contains string
func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}


func main() {

  // configuration is stored in this
  // tiny little struct made with <3
  c := Config{
    rvocmd: "rvo export -c docs",
    hugocmd: "hugo",
    sammelsuriumdir: "/tmp/foo",
    homedir: "/home/noqqe/Code/noqqe.de"}

  // Some warm welcome
  log.Println(".oO(Amaze - Wow)Oo.")

  if os.Args[1] == "build" {
    build(c)
  } else if os.Args[1] == "sammelsurium" {
    sammelsurium(c)
  }
}

