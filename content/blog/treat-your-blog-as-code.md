---
title: "Treat Your Blog as Code"
date: 2019-01-23T12:41:44+01:00
tags:
- CI
- CD
- TravisCI
- markdownlint
- hugo
---

Mein Blog besteht schon seit Langem einfach nur aus Text Files. Keine
Datenbank, kein dynamischer Content. Eignet sich das nicht eigentlich
hervorragend um [CI/CD](https://en.wikipedia.org/wiki/CI/CD) zu machen?

{{< figure src="/uploads/2019/01/codefry.png" >}}

Wie ich was und warum gemacht habe, wird hier erklärt.

TL;DR: `markdownlint`, hugo Build und `rsync` Deploy via
[travis-ci.org](https://travis-ci.org)

## Markdown Style Testing

Per Zufall bin ich auf
[markdownlint-cli](https://github.com/igorshubovych/markdownlint-cli)
gestoßen habe es auf meine Posts losgelassen. Tatsächlich ging das ganze
CI/CD für meinen Blog mit dem Fund dieses Stücks Software los.

```
markdownlint content/blog/*.md
```

Anfangs waren das ca.  3150 Errors :/ Das hat mich erstmal ziemlich
schockiert, aber klar viele Markdown Lint
[Rules](https://github.com/DavidAnson/markdownlint) sind nicht unbedingt
Errors sondern eher Style Warnings.

Mit etwas Konfiguration, welche Regeln ich nicht ausgewertet haben möchte,
hab ich dann die `.markdownlintrc` gebaut:

```
{
  "default": true,
  "fenced-code-language": false,
  "first-heading-h1": false,
  "heading-increment": false,
  "no-trailing-punctuation": false,
  "first-line-h1": false,
  "line-length": false,
  "ul-style": false,
  "no-inline-html": {
     "allowed_elements": [ "code" ]
  },
  "header-style": {
    "style": "atx"
  }
}
```

Und war dann runter auf 1150. Diese restlichen Errors erfordeten `sed` und Handarbeit.

{{< figure src="/uploads/2019/01/handarbeit.png" >}}

Sehr viel Handarbeit.

## CI mit Travis

Nachdem das jetzt alles “schön” war hab ich mich von [Poschi3](h) mit
[travis-ci.org](https://travis-ci.org) anfixen lassen.

Ich wollte `markdownlint` bei jedem Commit über mein Repo laufen lassen.
Dafür hab ich mich dort eingeloggt und das Repo `github.com/noqqe/noqqe.de` aktiviert.
Danach die `.travis.yml` mit folgendem Inhalt gefüllt.

```
language: generic

install:
  - npm install -g markdownlint-cli

script:
  - markdownlint content/*.md
  - markdownlint content/blog/*.md
  - markdownlint content/sammelsurium/*.md
```

Und dann lief die Sache auch (nach gefühlt 130 Anläufen) auch schon.

{{< figure src="/uploads/2019/01/testmarkdown.png" >}}

Siehe [Build](https://travis-ci.org/noqqe/noqqe.de/builds/479452544)

## Building

Markdown Syntax Qualität ist jetzt also sichergestellt. Wenn ich meinen
Content richtig befülle, muss ich allerdings auch noch darauf achten, dass
mein `hugo`  Blog auch wirklich compiled. Also aus dem ganzen Markdown + Hugo
Templating auch wirklich lesbares HTML wird.

Das geht nun wieder etwas anders. Diesmal muss ich der `.travis.yml` noch
beibringen, `go` und das zugehörige Paket `hugo` zu installieren und
schliesslich auszuführen.

```
language: go
go:
  - '1.11'

install:
  - go get github.com/gohugoio/hugo

script:
  - hugo
```

Wenn der Blog also jetzt einen neuen Commit bekommt, wird sowohl Markdown
überprüft als auch ob der statische Blog Compiler den Blog wirklich bauen
kann.

{{< figure src="/uploads/2019/01/testhugo.png" >}}

Yay.

## Deploying

Wenn man schon den CI Teil erledigt hat, kann der Continious Deployment Teil
ja auch gleich mit erledigt werden.

Im Grunde brauche ich nur einen `rsync` des Verzeichnisses `public` auf einen
meiner Server. Deployment wäre um ein vielfaches leichter, wären da nicht
**Geheimnisse** notwendig. Ich will in meinem Deployjob nicht public lesbar
haben auf welchen Server, mit welchem User und schon garnicht mit welchem
**Passwort/Key** sich auf dem Server angemeldet wird.

Das haben auch die Leute bei Travis erkannt und dafür ein HowTo verfasst das
ich jetzt hier nicht replizieren muss: [Encrypting Files - Travis
CI](https://docs.travis-ci.com/user/encrypting-files/). Im Grunde erzeuge ich
einen SSH Key den ich via Commandline zu Travis hochlade und  dann noch ein
paar “geheime” Variablen im Build Job setze, die ich nachher verwenden kann.

{{< figure src="/uploads/2019/01/travisvars.png" >}}

Um den Deploy zu erledigen gibt es bei Travis `deploy` und `before_deploy`.
Meine aktuelle Lösung sieht so aus:

```
install:
  - echo $KNOWN_HOSTS_KEY >> $HOME/.ssh/known_hosts

before_deploy:
  - openssl aes-256-cbc -K $encrypted_a01276847ef6_key -iv $encrypted_a01276847ef6_iv -in deploy_rsa.enc -out /tmp/deploy_rsa -d
  - eval "$(ssh-agent -s)"
  - chmod 600 /tmp/deploy_rsa
  - ssh-add /tmp/deploy_rsa

deploy:
- provider: script
  skip_cleanup: true
  script: rsync -avi public/ $PRODUCTION_SERVER_USER@$PRODUCTION_SERVER:public/
  on:
    branch: master
```

**Achtung**: Achtet bitte darauf welchen User mit welchen Rechten ihr auf dem
Zielserver einrichtet, selbst wenn das alles nicht “public” für jedermann
ersichtlich ist, kann sich der Inhaber (oder ein späterer Angreifer) von
TravisCI auf eurem Server mittels dieses Users+Keys einloggen.

Wenn das alles klappt sieht das dann so aus:

{{< figure src="/uploads/2019/01/deploy.png" >}}

Oder siehe hier der
[Build](https://travis-ci.org/noqqe/noqqe.de/builds/483281484)

Anmerkung zur `echo $KNOWN_HOSTS_KEY >> $HOME/.ssh/known_hosts` Sache: Es
gibt zwar ein [Addon](https://docs.travis-ci.com/user/ssh-known-hosts/) für
Known Hosts, das wird aber **immer** VOR dem auswerten der Globalen Variablen
ausgeführt und ist somit nutzlos weil ich dann den Hostname nicht als Travis
Setting einfügen kann.

## Was jetzt?

Tatsächlich ist das der erste Blogpost den ich mit der vollautomatischen
Chain hier deploye.  Ich denke ich werde wenn das alles sich als sinnvoll
herausstellt auch noch [entbehrlich.es](j) automatisieren, da gibt es noch
ein paar Steps die ich bisher immer per Hand mache.

Wenn das Endresultat interessiert kann sich auf [Github](h) die `.travis.yml`
nochmal anschauen, in der alle 3 Steps zusammengefasst sind.

Natürlich auch die Build Logs für den aktuellen Blogpost kann man
[hier](https://travis-ci.org/noqqe/noqqe.de) finden.
