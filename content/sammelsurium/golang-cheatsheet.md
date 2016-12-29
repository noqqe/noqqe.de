---
title: Golang Cheatsheet
date: 2016-12-18T12:04:06
tags: 
- Programming
- go
- golang
---

#### Logging

Besser als einfach nur fmt.Println mit

``` golang
import log

log.Printf("%s\n", string(output))
log.Println("foo")

log.Fatal("Process crashed and burned. Error: ", err)
```

#### Structs

Definieren

``` golang
type Config struct {
  path string
  timeout int64
}
```

Objekt des Typs Config erstellen

``` golang
c := Config{path:"/foo/bar", timeout: 256}
```

Abrufen

``` golang
log.Println(c.path)
```

#### Command Execution

``` golang
// execute commands
cmd := exec.Command(command, args...)
cmd := exec.Command("ls", "-lah", "-r")
output, err := cmd.CombinedOutput()

// check for errors
if err != nil {
	log.Fatal("rvo crashed and burned.")
}

```

#### Slices

In Python kann ich einfach

	if "string" in list:

In Go brauch ich dafuer die Hilfe dieser kleinen Funktion

```
// check if slice contains string
func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
```

#### JSON

Parse JSON from string

``` golang
// convert to bytes and read json
bytes := []byte(output)
var documents []Document
json.Unmarshal(bytes, &documents)

log.Printf("Entries found: %d", len(documents))
```

Ein JSON Objekt mit nested documents

``` golang
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
```


