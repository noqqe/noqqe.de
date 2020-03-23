---
title: "Golang"
date: 2020-03-16T14:53:12+01:00
tags:
- Programming
---

## Datentypen

* Boolean types
* Numeric types
* String types
* Derived types
  * Pointer types
  * Array types
  * Structure types
  * Union types and
  * Function types
  * Slice types
  * Interface types
  * Map types
  * Channel Types

## builtin's

```
len()
```

```
make()
```

## new

`new()` ist eine Builtin Funktion
Um einen `int` ohne Namen zu erzeugen und aus einer Funktion
zurück zu geben.

```
new()
```

So bewirkt

```
func newInt() *int {
  return new(int)
}
```

das Gleiche wie

```
func newInt() *int {
  var dummy int
  return &dummy
}
```

Aber ist schöner und kürzer, weil man sich keinen Namen ausdenken muss der
unnötig wäre.

## fmt

Mal ein paar Notizen zu `fmt`

Um eine einzelne Zeile auszugeben, wie `echo` in der Normalbenutzung

    fmt.Println("foo")

Wenn ich noch den Filedescriptor angeben will, muss ich eine andere Funktion
benutzen

    fmt.Fprintln(os.Stderr, "foo")

Und wenn ich diverse Formatzeichen nutzen will, ebenso.

    fmt.Printf("%s %d", lol, foo)


## bufio

`bufio` ist eine Standard Library mit sich unter anderen Files lesen lassen.
Buffered IO halt.

    fhandle, err := os.Open("/path/to/file")
    input := bufio.NewScanner(fhandle)

Um den Input dann zu verarbeiten

    for input.Scan() {
      fmt.Println(input.Text())
    }


## os

OS hat ein paar Toole Bindings

Eine Datei oeffnen und ein Filehandle zurück bekommen

    os.Open("file")

Von Stdin lesen

    os.Stdin

`os` beherrscht auch Argumente des Programms

    os.Args[1]

## strings

Stringmanipluationen aller Art bietet `strings`. Zum Beispiel ein `Slice` zu
einem `String` verwandeln

    strings.join(Slice, sep)

## flag

Flag ist sozusagen das `os.Args` auf Steroiden. Der eingebaute
Argumentenparser kann auch Flags von Argumenten unterscheiden

```
var n = flag.Bool("n", false, "no newline")

func main() {
  flag.Parse()
  fmt.Print(strings.Join(flag.Args(), " "))
  if !*n {
    fmt.Println()
  }
}
```

## Fehlerbehandlung

Die meisten Funktionen geben auch einen `err` zurück beim Aufruf

```
file, err := os.Open("/file/that/does/not/exist.txt")
if err != nil {
  fmt.Printf("%v\n", err)
}
```

`nil` ist dabei ein eingebauter Datentyp speziell für Fehler in Go.

## const

Konstanten werden speziell markiert. Sie bleiben zur Compilezeit gleich
und werden niemals verändert.

```
const (
  foo = "lol"
  bar = 42
)
```

oder auch in Funtkionen

```
func countlines () {
  const (
    foo = "lol"
    bar = 42
  )
  [...]
}
```

## Logging

Besser als einfach nur fmt.Println mit

``` golang
import log

log.Fatal("Process crashed and burned. Error: ", err)
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

## Command Execution

Natürlich kann ich auch einfach Shell Commands wrappen in Go.

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

## Slices

Wie finde ich einen String in einem `Slice`

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

## json

``` golang
// convert to bytes and read json
bytes := []byte(output)
var documents []Document
json.Unmarshal(bytes, &documents)

log.Printf("Entries found: %d", len(documents))
```
