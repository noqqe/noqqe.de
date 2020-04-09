---
title: "HTTP Debugging: MaxUploadSize"
date: 2020-04-08T15:44:39+02:00
tags:
- golang
- HTTP
- Loadbalancer
---

Erzfeind ist ein hartes Wort. Das es mir beim Beschreiben meines Gefühls zu
dem auf der Arbeit verwendeten (proprietären) Loadbalancers in den Sinn kommt, kann aber
auch kein Zufall sein.

<!--more-->

User berichteten mir über Verbindungsabbrüche beim Übertragen großer
Requests. Ich hatte den Loadbalancer in Verdacht. Seinen (wenig einfühlsamen)
Umgang mit HTTP Anfrangen hatte er bereits in der Vergangenheit unter Beweis
gestellt. Zum Beispiel als `HTTP GET` Parameter nach einer bestimmten Länge
einfach abgeschnitten wurden. Warum sollte es also bei `HTTP POST` Max Body
Size anders sein?

{{< figure src="/uploads/2020/04/uploadlimit.jpg" caption="Auf der Suche nach MaxUploadSize Memes fand ich dieses schöne Stück :D" >}}

TL;DR: Es war anders.

## Was brauchte ich also?

Ich will das Problem nachstellen. Warum bricht die Übertragung ab?

Entweder ich fange an die API der Software (die ich betreibe) hinter dem
Loadbalancer zu verstehen und baue dann einen `curl` Befehl nach oder...

...oder ich schreibe mir einen kleinen HTTP Daemon. Dieser müsste mir nur
anzeigen wie viel Daten übertragen wurden. Yes.

## HTTPD

Da ich in der COVID-19 Isolation gerade `Go` lerne, war das eine gute
Gelegenheit etwas Praktisches zu tun.

Die Go HTTP Library `net/http` ist einfach zu bedienen und die
[Doku](https://golang.org/pkg/net/http/#Response) hilft
viel.

```golang
func post(w http.ResponseWriter, r *http.Request) {

  // fetch header size
  size, err := strconv.Atoi(r.Header["Content-Length"][0])
  if err != nil {
    log.Fatal(err)
  }

  // fetch body size
  buf := new(bytes.Buffer)
  buf.ReadFrom(r.Body)
  body := len(buf.String())

  // bytes to mb
  mbh := float64(size)/1000000
  mbb := float64(body)/1000000

  w.WriteHeader(200)
  fmt.Fprintf(w, "Header: %.2f MB, Body: %.2f MB!\n", mbh, mbb)
  fmt.Printf("Header: %.2f, Body: %.2f MB!\n", mbh, mbb)
}
```

Kümmern um den eigentlichen Request, tut sich hier keine einzige Zeile. Was
interessant ist, ist das `HTTP Header` Feld `Content-Length`. Darin befindet
sich die Größe des übertragenen Contents. Diese musste
ich auslesen und zu `MB` konvertieren. *Link zum ganzen Sourcecode, siehe unten.*

## Payload erzeugen

Um Files in der notwendigen Größe zum Übertragen bereit zu haben dachte ich
zuerst an das müßige `dd`, was immer ewig braucht um Files von ein paar
wenigen Gigabytes zu erzeugen. Dann fand ich `truncate`, dass leere Files
im Filesystem allokiert. Und das ist um ein vielfaches schneller!

{{< figure src="/uploads/2020/04/truncate.png">}}
{{< figure src="/uploads/2020/04/dd.png">}}

Meine 6 Usecases sind so schnell generiert:

```fish
for x in 1M 5M 10M 100M 1G 10G
  truncate -s $x /tmp/$x.bin
end
```

## Der Test

Jetzt sind alle kompoenenten des Tests beisammen. Also Applikation hinter dem
Loadbalancer herunterfahren, `maxupload.go` kompilieren und an den Start
bringen.

```
> go build maxupload.go
> ./maxupload 8080
```

Auf meinem MacBook warteten bereits die leeren Dateien via `curl` übertragen
zu werden. Ergebnis:

```fish
$ for x in 1M 5M 10M 100M 1G 10G
  curl -X POST --data-binary @/tmp/$x.bin localhost:8000
end

Header: 1.05 MB, Body: 1.05 MB!
Header: 5.24 MB, Body: 5.24 MB!
Header: 10.49 MB, Body: 10.49 MB!
Header: 104.86 MB, Body: 104.86 MB!
Header: 1073.74 MB, Body: 1073.74 MB!
Header: 10737.42 MB, Body: 10737.42 MB!
```

Mist(?). Der Loadbalancer hat also kein Problem mit 10G im `HTTP POST` Body. Mhpf.

Was es ist weiss ich immernoch nicht, aber zumindest hab ich ein Problem mit
der MaxUploadSize jetzt ausgeschlossen. Und ein bisschen Go geübt.

Der Source für dieses kleine Go Projekt ist hier zu finden:
[gist.github.com](https://gist.github.com/noqqe/f5cb617897eb495ea60a9f68be5ca412)

**Update**: [Morris](https://twitter.com/MorrisJobke) hat mich darauf
[hingewiesen](https://twitter.com/MorrisJobke/status/1247874283810029568) das
die ursprüngliche Version nur die HTTP Headers analysiert hat. Diese werden
von `curl` generiert und haben keine Aussage ob der Body des Requests
vollständig ankommt. Daher hab ich den `Go` Teil aktualisiert. Danke!
