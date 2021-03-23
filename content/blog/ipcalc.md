---
title: "ipcalc in Go"
date: 2021-03-23T15:01:35+01:00
tags:
- ipcalc
- golang
---

Als Fachinformatiker werden sich viele der Leser hier an die Zeit erinnern,
in der man in der Berufsschule [Subnetting]() bei gebracht bekam. Entweder
mit Schrecken oder gelangweilt.

<!--more-->

{{< figure src="/uploads/2021/03/subnetmeme.png" >}}

Für den täglichen Gebrauch nutze ich seit Jahren [ipcalc](). Mal kurz ein
Netz in [CIDR]() Notation double-checken.


## Perl aus 2006

{{< figure src="/uploads/2021/03/ipcalc.png" caption="Traditionelles ipcalc (2006)" >}}

Ein bisschen komisch war der Output für mich immer, aber die Infos die ich brauchte waren dabei.
Wenn ich zum Beispiel 8 Subnets einem aufteilen möchte, muss ich das Tool so
bedienen:

```
ipcalc <cidr> -s 30 -s 30 -s 30 -s 30 -s 30 -s 30 -s 30 -s 30 #WTF?
```

Außerdem rechnet/bedient es sich immer mit *usable ip addresses*? Ok?!

Ich begann darüber nachzudenken wie schwer es wohl ist das Tool meinem
"Verständnis von Benutzerfreundlichkeit" anzupassen. Ein kurzes `brew info
ipcalc` zeigt dass dahinter Mitglieder des CCC Göttingen stecken und seit
2006 nichts mehr passiert ist: http://jodies.de/ipcalc-archive/

Aber das schlimme ist: Es ist Perl.

## ipcalc in Go

Ich begann nach Go Libraries zu suchen, welche ich für die ganze IP
Calculation benutzen konnte. Schliesslich ist IP Validation [hart (toller Blogpost)](https://blog.dave.tf/post/ip-addr-parsing/).

Dabei fand ich eine Codebase unter MIT Lizenz, die eigentlich vieles was ich
will schon kann. Nämlich [asenci/ipcalc](https://github.com/asenci/ipcalc).
Cool, einfach, verständlich. Genau was ich wollte?

Sogar mit Features die ich selbst noch implementiert hätte! Zum Beispiel das
Zusammenfassen von mehreren sich aneinanderreihenden Subnetzen.

{{< figure src="/uploads/2021/03/ipcalc-agg.png" caption="ipcalc aggregate" >}}

Aber auch anderes wie *Überlappungen* (`overlap`) oder eine `exclude`
Funktion!

Wie es sich für einen ordentlichen Hacker gehört, hab ich natürlich erstmal
ein bisschen Output eingefärbt ([dracula](https://draculatheme.com) &lt;3) und so formatiert wie ich es mir von dem
traditionellen `ipcalc` gewünscht hätte.

{{< figure src="/uploads/2021/03/ipcalc-show.png" caption="ipcalc in Go" >}}

Die Syntax von `subnet`, einfach toll. Ich möchte bitte einfach 4 Netze
haben. Nicht `-s 30 -s 30...`. DANKE!

{{< figure src="/uploads/2021/03/ipcalc-subnet.png" caption="ipcalc subnet" >}}

Was ich selbst dazu gebaut habe ist einen `--verbose` Mode für den `ipcalc subnet` Split
Befehl.

{{< figure src="/uploads/2021/03/ipcalc-subnet-verbose.png" caption="Subnet splitting mit --verbose" >}}

Somit war ich dann zuerstmal zufrieden. Ich habe für einzelne Features
Pullrequests beim Upstream Repo erstellt, aber bisher scheint der Autor daran
noch kein sonderliches Interesse zu haben.

## Homebrew & Go

"Vielleicht doch mal ein `brew` Formula schreiben?". Eine Installation über
`brew` wäre jedenfalls cool!

Ich las [die Dokumentation über Taps](https://docs.brew.sh/How-to-Create-and-Maintain-a-Tap) und erstellte
das dazugehörige [GitHub Repo](https://github.com/noqqe/homebrew-tap/)

Wie ich jetzt am Besten meine `go` Software baue ist mir komplett abgenommen
worden, da ich bei anderen Github Projekten
[goreleaser](https://goreleaser.com) gefunden habe.

`goreleaser` baut das eigene Go Projekt und besitzt mehrere Adapter. In
meinem Fall den Github und den Homebrew Adapter. Alles definiert in
`.goreleaser.yml`

```yaml
before:
  hooks:
    - go mod tidy
    - go generate ./...
builds:
  - env:
      - CGO_ENABLED=0
    goos:
      - linux
      - windows
      - darwin
archives:
  - replacements:
      darwin: Darwin
      linux: Linux
      windows: Windows
      386: i386
      amd64: x86_64
checksum:
  name_template: 'checksums.txt'
snapshot:
  name_template: "{{ .Tag }}-next"
changelog:
  sort: asc
  filters:
    exclude:
      - '^docs:'
      - '^test:'
brews:
  - name: ipcalc
    goarm: 6
    tap:
      owner: noqqe
      name: homebrew-tap
    download_strategy: CurlDownloadStrategy
    folder: Formula
    homepage: "https://github.com/noqqe/ipcalc"
    description: "ipcalc - written in go"
    license: "MIT"
    dependencies:
      - name: go
    install: |
      bin.install "ipcalc"
```

Ich hab also immernoch keine Ahnung wie man eine Homebrew Formula schreibt,
oder Assets nach Github hochlädt, aber das brauche ich garnicht :) Würde
`goreleaser` jedem ans Herz legen.

```
brew install noqqe/tap/ipcalc
```

Und das wars :) Danke an die Autoren von `ipcalc` und das ihr mir Jahrelang
gute Dienste erwiesen habt, aber es ist zeit weiterzuziehen :)
