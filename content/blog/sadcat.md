---
date: 2017-02-03T17:00:02+01:00
tags:
- ssh
- config
- python
title: sadcat
---

Vor Kurzem schrieb ich einen Config Generator für `ssh`. Das hört sich
unsinnig an und ist es wahrscheinlich auch. Der Grund warum ich das tat
liegt nichtmal an der Tatsache, dass Funktionalität im Konfigurationsformat
fehlen würde. Nein. Vielmehr daran das ich mir angewöhnte kleine Aliase für
Server zu benutzen statt deren vollständigen FQDNs.

{{< figure src="/uploads/2017/01/sadcat.jpg" >}}

Hostnames mit denen ich üblicherweise konfrontiert werde sehen (leicht
obfuscated) so aus: `nyc-cexapsdrap21.company.com`. Ich gewöhnte mir an einfach
`drap21` als Alias zu benutzen.

```
Host nyc-cexapsdrap21.company.com drap21 pdrap21
  Hostname nyc-cexapsdrap21.company.com
  User myuser
  Port 22
  IdentiyFile ~/.ssh/project_id_rsa
```

Mit der Zeit wurde meine SSH Config sehr lang (Überraschung!) und fast
unwartbar. Ich half mir immer wieder mit irgendwelchen Shell Loops a la

```
for x in {1..9}; do
  echo 'Host fra-foo0${x} drap$x'
  echo '  Hostname fra-foo0${x}'
  echo ...
done
```

und Edits mit `sed`. Wie man das eben als fauler Admin so macht.

Als ich dann irgendwann Zeile 1000 Überschritt war auch bei mir mal die
maximale Leidensfähigkeit erschöpft und irgendwas musste anders werden.

Die `PATTERNS` Sektion in der `ssh_config` Manpage gab nichts her und
umgewöhnen wollte ich mich einfach nicht. Die Fähigkeit dynamische Aliases
zu generieren gehört nicht zu Notwendigkeiten so einer Software. Das würde
eine Art Templating innerhalb der ssh Config vorraussetzen und das gehört
da nicht hin. Deshalb fand ich es auch völlig okay mir selber etwas zu
bauen.

Daraus entstand dann [sadcat](https://github.com/noqqe/sadcat). Was
`sadcat` tut ist, ein `toml` Config File annehmen und daraus eine
(hoffentlich) valide SSH Config bauen.

```
[hosts]

[hosts.nyc-cexapsdrap]
hostname = "nyc-cexapsdrap[15-36]"
alias = [ "drap", "pdrap" ]
template = "nyc-admins"

[templates]

[templates.nyc-admins]
user = "myuser"
port = "22"
identityfile = "~/.ssh/project_id_rsa"
```

Woraus dann `nyc-cexapsdrap15` bis `nyc-cexapsdrap36` generiert werden und
eben auch die dazugehörigen Aliases.

Technisch für mich interessant das mal zu bauen war Templating und die
"Vererbung" (die eigentlich nur ein naiver Dict Merge ist) und die ein paar
Schritte in Richtung Config Parsing zu machen.
Schön zu sehen was in 150 Zeilen schlechtem Python entstehen kann.

Falls das Eingangs noch nicht klar geworden ist: Besser nicht benutzen. Die
SSH Config bietet ein breites Spektrum an Möglichkeiten seine Config
einfach und klein zu halten. Nur so komische Usecases wie diesen eben
nicht.
