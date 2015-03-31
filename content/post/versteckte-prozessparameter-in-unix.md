---
type: post
title: "Versteckte Prozessparameter in UNIX"
date: 2014-05-23T22:14:00+02:00
comments: true
categories:
- ubuntuusers
- osbn
- mongodb
- mysql
- c
- posix
- c99
- unix
- ps
- process
- prozesslisten
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

{% blockquote C99 Standard %}
The parameters argc and argv and the strings pointed to by the argv array shall be modifiable by the program, and retain their last-stored values between program startup and program termination.
{% endblockquote %}

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
