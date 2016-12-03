package main

import (
  "os"
  "strings"
  "time"
)

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
