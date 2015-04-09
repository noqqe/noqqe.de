---
date: 2011-03-26T14:45:37+02:00
type: post
comments: true
title: 'ZRE | SQL Statistik Modul'
aliases:
- /archives/1566
- /blog/2011/03/26/zre-sql-statistik-modul
categories:
- Bash
- Debian
- Development
tags:
- attack
- bash
- chewbacca
- Events
- game
- github
- han solo
- modul
- module
- mysql
- shell
- sql
- zombie
- zombie revolution environment
- zombies
- zre
---

Ich möchte kurz ein Wort über das SQL Modul verlieren, welches die
Statistiken des permanent laufenden [Zombie
Environments](http://zombies.n0q.org) aufzeichnet. Wie mit Sicherheit schon
überall auf dem Planeten Erde bekannt ist, passieren in ZRE Events. Diese
Hand voll Events lösen (wie der Name vielleicht schon suggeriert)
Geschehnisse in innerhalb des Environments aus, welche es gilt statistisch
auszuwerten und darzustellen.

{{< figure src="/uploads/2011/03/0744_a999_450.gif" >}}

Wie aber erfasse ich solche Informationen am Besten. Ich hatte ehrlich
gesagt keine große Lust, in jedes der Events einzeln eine SQL-Verbindung
aufzubauen und den gewünschten Query abzusetzen. Ganz zu schweigen von der
Problemfallbehandlung (MySQL nicht installiert, Modul in Config nicht
aktiviert, MySQL nicht erreichbar usw.) die ich in jedem Event hätte extra
behandeln müssen.


```
events/
├── attack.humans.event
├── attack.zombies.event
├── born.humans.event
├── born.zombies.event
├── building.humans.event
├── collecting.zombies.event
├── die.humans.event
├── die.zombies.event
├── infos.event
├── stats.humans.event
├── stats.zombies.event
├── support.humans.event
├── support.zombies.event
└── weather.event
```


Stattdessen entschied ich mich ein kleines Modul in Bash zu schreiben, welches eine Hand voll übergebener Daten annimmt und dann in die konfigurierte Datenbank schreibt.  Hier am Beispiel des weather.event

``` bash
case $nature_msg in
1) echo "A volcano explodes." ; zresql weather vulcano ;;
2) echo "An earthquake hits the city." ; zresql weather earthquake ;;
3) echo "A hurricane hits the city." ; zresql weather hurricane ;;
esac
```

Als erstes übergeben wird immer der gewünschte Modus, in den zresql den
Datensatz einordnen soll, danach alle weiteren Informationen die der
Struktur nach erforderlich sind.

```
# Hurricane
$ zresql weather hurricane
# Humans gewinnen einen Kampf und vernichten 643 Zombies
$ zresql kill human 643
```

Das zresql Modul nimmt die Informationen dann entgegen und leitet es unter
den gegebenen Umständen via sqlsend an die Datenbank weiter.

``` bash
[...]
weather) sqlsend "INSERT INTO zre_weather VALUES (NULL, "$2", CURRENT_TIMESTAMP);" ;;
kill) sqlsend "INSERT INTO zre_kills VALUES (NULL , "$2", "$3", CURRENT_TIMESTAMP);" ;;
[...]
```

sqlsend ist im Grunde nur eine weitere kleine Funktion die direkt über
Kommandozeile den SQL Query abschickt.

``` bash
sqlsend() {
mysql -e "use $sqldb; $1" --user=$sqluser --password=$sqlpw --host=$sqlhost
}
```

Mit dieser Lösung bin ich eigentlich relativ zufrieden. Über
Verbesserungsvorschläge und Kritik freue ich mich natürlich wie immer. Das
ganze Modul ist zusehen auf Github:
[https://github.com/noqqe/zombie-revolution-environment/blob/master/lib/sqlstats.library.bash](https://github.com/noqqe/zombie-revolution-environment/blob/master/lib/sqlstats.library.bash)

