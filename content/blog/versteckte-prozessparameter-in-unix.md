---
comments:
- author: Clemens
  content: "C99 sagt zwar, dass argv editierbar ist \u2013 allerdings muss das noch
    nicht hei\xDFen, dass damit auch die Kommandozeile, die in ps angezeigt wird auch
    ver\xE4ndert wird. In Linux ist das zwar so implementiert (man bekommt also quasi
    \xFCber /proc/$pid/cmdline eine Sicht in den Adressraum des Prozesses), andere
    (POSIX-kompatible) Systeme k\xF6nnten aber durchaus bei jedem execve(2) eine Kopie
    anlegen und die zur\xFCckgeben. Unter OS X funktioniert das allerdings auch, unter
    BSD und Linux hast dus vermutlich selbst getestet.\n\nDie sicherste Variante,
    die die Race-Condition beim Start vermeidet, ist und bleibt aber das Lesen des
    Passworts von einem Filedeskriptor, wie z.B. bei gpg --passphrase-fd."
  date: '2014-05-23T23:14:33.352542'
- author: noqqe
  content: "Yep, sowas hab ich auch beim Googeln gefunden. Siehe\_http://stackoverflow.com/questions/3830823/hiding-secret-from-command-line-parameter-on-unix.
    Total gut! Danke!"
  date: '2014-05-23T23:21:35.819351'
- author: tux.
  content: "Und seit wann ist \"f\" ein valider Parameter f\xFCr \"ps\"? Unter OpenBSD
    jedenfalls ist er das nicht: http://www.openbsd.org/cgi-bin/man.cgi?query=ps\n\nAuch
    ohne \"f\" zeigt \"ps\" jedenfalls zumindest den Namen des Programms an:\n\nhttps://www.dropbox.com/s/n65ikfx0idde4zw/2014-05-24-004028_1024x768_scrot.png\n\nSch\xF6n,
    wenn es unter Linux geht, aber UNIX mag das halt nicht.\nTja."
  date: '2014-05-23T23:50:59.692086'
- author: noqqe
  content: muss an dem komischen editor liegen den du da nutzt ;)
  date: '2014-05-23T23:54:55.581969'
- author: Vain
  content: Indeed, again what learned! :)
  date: '2014-05-24T10:02:16.255188'
- author: Anonymous
  content: "nice one!\nAber \"Again what learned\", srsly?! Das ist selbst f\xFCr
    einen Witz zu schmerzhaft. :("
  date: '2014-05-24T16:09:30.897238'
- author: Shogun
  content: "Passw\xF6rter in der Cmdline sind generell eine schlechte Idee. Entweder
    sollte, wie schon kommentiert, mit Filedeskriptoren gearbeitet werden oder z.B.
    das f\xFCr MySQL daf\xFCr vorgesehene Defaults-File verwendet werden."
  date: '2014-05-25T14:00:27.553601'
- author: noqqe
  content: "Das kommt drauf an. Bevor das Passwort 365 Tage im Jahr in einem File
    steht, \xFCbergeb ichs lieber readline beim Aufruf. Da ists dann nur im Memory
    und evtl. mal kurz in der Prozessliste. Wohingegen iwelche default-files jeder
    kennt und einsehen kann."
  date: '2014-05-26T09:42:57.416607'
- author: Geier
  content: '"StackOverflow-Driven-Development" -- nice.'
  date: '2014-05-26T20:22:21.112065'
- author: Anonymous
  content: "Es ist immer wieder erstaunlich was mit C so alles m\xF6glich ist..."
  date: '2014-05-26T21:16:33.342711'
date: '2014-05-23T20:14:00'
tags:
- ps
- prozesslisten
- process
- mongodb
- c99
- databases
- administration
- c
- unix
- posix
- mysql
title: Versteckte Prozessparameter in UNIX
---

Passwörter für Datenbanken beispielsweise sind Optionen die sich als Commandline Argument direkt im Aufruf mitgeben lassen.
Bei MySQL oder MongoDB ist das angegebene Passwort aber in der Prozessliste durch `xxxx` ersetzt.

```
$ mysql -u noqqe -ppassw0rd -h localhost
$ ps auxfww
sshd: noqqe@pts/0
 \_ -bash
     \_ mysql -u noqqe -px xxxxx -h localhost
     \_ ps auxfww
```

Irgendwie blieb ich die Woche an dieser Tatsache hängen. Verstand ich nicht.
Das OS bekommt doch den Aufruf des Programms und das Binary parst die bereits
vorher übergebenen Paramter. Ist das Binary nur eine Art Wrapper, der einen neuen
Prozess spawnt? Oder Linux Kernel API wie [hidepid](https://www.kernel.org/doc/Documentation/filesystems/proc.txt)?
Wer filtert hier?

### Passwort Parameter in MySQL

Nachdem MySQL ja OpenSource ist, kann man ja mal etwas `grep`en im [Source](http://bazaar.launchpad.net/~mysql/mysql-server/5.5/view/head:/client/mysql.cc#L1734).
Wurde schliesslich auch fündig.

``` cpp
case 'p':
  if (argument == disabled_my_option)
    argument= (char*) "";     // Don't require password
  if (argument)
  {
    char *start= argument;
    my_free(opt_password);
    opt_password= my_strdup(argument, MYF(MY_FAE));
    while (*argument) *argument++= 'x';     // Destroy argument
    if (*start)
      start[1]=0 ;
    tty_password= 0;
  }
  else
    tty_password= 1;
  break;
```

Das MySQL Client Binary wird also gestartet, initialisiert und die Variable `argument`, die aus dem Parameter-Parser von MySQL fällt, kopiert und
direkt an der entsprechenden Speicheraddresse mit `x`en überschrieben.

{{< figure src="/uploads/2014/05/memoryloss.jpg" >}}

Im Endeffekt eine coole Lösung, aha-Effekt war da. Bedeutet aber auch, dass beim
Start des Programms für eine gewisse Zeit das Passwort in der Prozessliste
steht. So lang bis der Argumentenparser an der entsprechenden Stelle angekommen ist und
den Memory überschreibt. Ab jetzt also immer `-p` Parameter ganz am Anfang hinschreiben :P

### Nachbau

Hört sich etwas nach zurecht gehackt an, fand ich. Dabei ist die Anpassbarkeit
durchaus im C99 Standard vorgesehen.

> The parameters argc and argv and the strings pointed to by the argv array
> shall be modifiable by the program, and retain their last-stored values
> between program startup and program termination.

Ausprobieren lässt sich das eigentlich mit einfach ein bisschen C, welches ich mir
via StackOverflow-Driven-Development zusammen geklaut habe.

``` c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char **argv) {

    int i = 0;
    int k=0;
    int len = 0;

    len = strlen(argv[0]);
    for(k=0;k<len;k++) {
        argv[0][k] = 'y';
    }

    len = strlen(argv[1]);
    for(i=0;i<len;i++) {
        argv[1][i] = 'x';
    }

    system("ps f");

    return 0;
}
```

Dabei kann auch der eigentliche Name des Programms überschrieben werden. Total
evil-haxx0r.

```
$ gcc hide.c -o hide
$ ./hide tolorlerolero
  PID TTY      STAT   TIME COMMAND
12384 pts/1    Ss     0:01 -bash
23512 pts/1    S+     0:00  \_ yyyyyy xxxxxxxxxxxxx
23513 pts/1    S+     0:00      \_ sh -c ps f
23514 pts/1    R+     0:00          \_ ps f
```

Again what learned.
