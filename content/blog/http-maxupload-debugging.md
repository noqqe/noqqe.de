---
title: "HTTP Debugging: MaxUploadSize"
date: 2020-04-08T11:19:43+02:00
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
func handle_post(w http.ResponseWriter, r *http.Request) {

  // fetch http header field and convert to megabytes
  size, _:= strconv.parseFloat(r.Header["Content-Length"][0], 64)
  mb := size/1000000

  // http response
  w.WriteHeader(200)
  fmt.Fprintf(w, "Size: %.2f MB!\n", mb)
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

Size: 1.048576 MB!
Size: 5.24288 MB!
Size: 10.48576 MB!
Size: 104.8576 MB!
Size: 1073.741824 MB!
Size: 10737.41824 MB!
```

Mist(?). Der Loadbalancer hat also kein Problem mit 10G im `HTTP POST` Body. Mhpf.

Was es ist weiss ich immernoch nicht, aber zumindest hab ich ein Problem mit
der MaxUploadSize jetzt ausgeschlossen. Und ein bisschen Go geübt.

Der Source für dieses kleine Go Projekt ist hier zu finden:
[gist.github.com](https://gist.github.com/noqqe/f5cb617897eb495ea60a9f68be5ca412)
