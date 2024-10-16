---
title: "Golang"
date: 2020-03-16T14:53:12+01:00
tags:
- Programming
---

## Warum go?

Ein Standard, nicht 100 Tools Eco System.

* test system
* module system
* `go vet`
* go routines
* `go fmt`

## Ordner Struktur

Leeres Git Repo

    go mot init practice

Folder angenommen:

```
- practice
    - go.mod
    - app.go
    - models
       - a.go
       - b.go
    - routers
       - a.go
       - b.go
```

Kann dann so abgerufen werden

```
import (
  "practice/routers"
  "practice/models"
  ...
)
```

Access zu Methoden

```
func main() {
  http.HandleFunc("/a", models.AHandler)
  http.HandleFunc("/b", models.BHandler)
  http.ListenAndServe(":8080", nil)
}
```

## Slice

Content für Element setzen

    foo[1] = "rofl"

Slice initialisieren mit default werten

    var foo []string = []{"foo", "bar"}
    var foo []int = []{1, 2, 3}

Ein bestehendes Slice erweitern

    s = append(s, "e", "f")

Slice mit Dimensionierung initialisieren

    s := make([]string, 3)

Slice mit Custom types initialisieren

    type Bar struct {
      x, y int
    }

    var foo []Bar

    foo[]Bar{Bar{1,1}}

Länge eines `Slice` bestimmen

    len(foo)

Wie finde ich einen String in einem `Slice`

``` go
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

## strings

Stringmanipluationen aller Art bietet `strings`. Zum Beispiel ein `Slice` zu
einem `String` verwandeln

    strings.join(Slice, sep)

## Konstanten const

Konstanten werden speziell markiert. Sie bleiben zur Compilezeit gleich
und werden niemals verändert.

``` go
const (
  foo = "lol"
  bar = 42
)
```

oder auch in Funtkionen

``` go
func countlines () {
  const (
    foo = "lol"
    bar = 42
  )
  [...]
}
```

## Structs

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

Fun Fact: Konstaten können wesentlich mehr Nachkommastellen enthalten als
Variablen

## Fehlerbehandlung

Die meisten Funktionen geben auch einen `err` zurück beim Aufruf

``` go
file, err := os.Open("/file/that/does/not/exist.txt")
if err != nil {
  fmt.Printf("%v\n", err)
}
```

`nil` ist dabei ein eingebauter Datentyp speziell für Fehler in Go.

## fmt

Mal ein paar Notizen zu `fmt`

Um eine einzelne Zeile auszugeben, wie `echo` in der Normalbenutzung

    fmt.Println("foo")

Wenn ich noch den Filedescriptor angeben will, muss ich eine andere Funktion
benutzen

    fmt.Fprintln(os.Stderr, "foo")

Und wenn ich diverse Formatzeichen nutzen will, ebenso.

    fmt.Printf("%s %d", lol, foo)

Formatstrings (wird noch fortgeführt)

```
%.2f float64 with 2 zeros
%s string
%i int
%v interface
%T show type of var
```

## bufio

`bufio` ist eine Standard Library mit sich unter anderen Files lesen lassen.
Buffered IO halt.

    fhandle, err := os.Open("/path/to/file")
    input := bufio.NewScanner(fhandle)

Um den Input dann zu verarbeiten

    for input.Scan() {
      fmt.Println(input.Text())
    }

Buffer zu String Konvertieren

    // response.Body = io.ReadCloser
    response, _ := http.Get("https://golangcode.com/")
    buf := new(bytes.Buffer)
    buf.ReadFrom(response.Body)
    newStr := buf.String()

## os

OS hat ein paar Toole Bindings

Eine Datei oeffnen und ein Filehandle zurück bekommen

    os.Open("file")

Von Stdin lesen

    os.Stdin

`os` beherrscht auch Argumente des Programms

    os.Args[1]

## strconv

Einen `string` zum `int` braucht einen Fehlerfall (Input: aaa?)

``` go
port, err := strconv.Atoi("8080")
```

## flag

Flag ist sozusagen das `os.Args` auf Steroiden. Der eingebaute
Argumentenparser kann auch Flags von Argumenten unterscheiden

``` go
var n = flag.Bool("n", false, "no newline")

func main() {
  flag.Parse()
  fmt.Print(strings.Join(flag.Args(), " "))
  if !*n {
    fmt.Println()
  }
}
```

## Logging

Besser als einfach nur fmt.Println mit

``` golang
import log

log.Fatal("Process crashed and burned. Error: ", err)
```

## Command Execution

Natürlich kann ich auch einfach Shell Commands wrappen in Go.

``` go
// execute commands
cmd := exec.Command(command, args...)
cmd := exec.Command("ls", "-lah", "-r")
output, err := cmd.CombinedOutput()

// check for errors
if err != nil {
  log.Fatal("rvo crashed and burned.")
}

```

## new

`new()` ist eine Builtin Funktion
Um einen `int` ohne Namen zu erzeugen und aus einer Funktion
zurück zu geben.

``` go
new()
```

So bewirkt

``` go
func newInt() *int {
  return new(int)
}
```

das Gleiche wie

``` go
func newInt() *int {
  var dummy int
  return &dummy
}
```

Aber ist schöner und kürzer, weil man sich keinen Namen ausdenken muss der
unnötig wäre.

## Testing

Go hat eine eigene Test Suite dabei. Dafür muss man nur ein File
`xxx_test.go` in das gleiche Verzeichnis legen und eine Funktion mit
`TestXxx` benennen. Beispiel:

``` go
import testing

// Simple substract Test
func TestSubInt(t *testing.T) {
  var want uint32 = 20
  if got := subInt(60, 40); got != want {
    t.Errorf("%d, want %d", got, want)
  }
```

```
go test -v
```

## Coverage

Auch Coverage kann man gleich im eingebauten Tool abfrühstücken:

```
> go test -cover -v
=== RUN   TestBrighten
--- PASS: TestBrighten (0.00s)
=== RUN   TestDarken
--- PASS: TestDarken (0.00s)
=== RUN   TestIso
    TestIso: nept_test.go:122: 20655, want 44247>20000
    TestIso: nept_test.go:122: 37364, want 10>32777
--- FAIL: TestIso (0.00s)
=== RUN   TestFlatten
--- PASS: TestFlatten (0.00s)
FAIL
coverage: 34.9% of statements
```

## Go Routines

Man kann sehr einfach Go Routines einbauen die für das bestehende Programm
Paralellität mit Threads einbaut.

``` go

import (
  "sync"
)

func main() {
  var wg sync.WaitGroup

  for a := 0; a < 10; a++ {
    wg.Add(1)
    go edit(x, y)
  }

  wg.Wait()
}

func edit() {
  defer wg.Done()
  fmt.Println(x+y)
  time.sleep(1)
}
```

## json

``` go
// convert to bytes and read json
bytes := []byte(output)
var documents []Document
json.Unmarshal(bytes, &documents)

log.Printf("Entries found: %d", len(documents))
```


## yaml

Beispiel YAML

```yaml
challenges:
  - Och ja... ich weiss auch nicht
  - "Test"
  - Ich weiss nicht
```

Und zum Ausgeben der Daten:

```go
package main

import (
  "fmt"
  "log"
  "io/ioutil" "gopkg.in/yaml.v2"
)

type Challenges struct {
    Challenge []string `challenges`
}

func (c *Challenges) getConf() *Challenges{

    yamlFile, _ := ioutil.ReadFile("challenges.yml")

    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}

func main() {

    var c Challenges
    c.getConf()

    for _, x := range c.Challenge{
      fmt.Println(x)
    }
}
```

Das `Struct` Challenges bekommt also eine Methode `getConf()` zum Laden und
gibt ein befülltes `Struct` zurück.

