---
comments:
- author: killermoehre
  content: "<p>Zur besseren Lastverteilung k\xF6nntest du noch <a href=\"http://cdn.debian.net\"
    rel=\"nofollow\">cdn.debian.net</a> als Paketquelle eintragen, das Content Delivering
    Network von Debian.</p>"
  date: '2013-08-31T18:33:40'
- author: Martin
  content: <p></p><pre>pip install https://github.com/rodjek/puppet-pygments-lexer/archive/master.zip<br>pbpaste
    | pygmentize -l puppet -f terminal -g /dev/stdin<br></pre><p></p><p>SCNR ;)</p>
  date: '2013-08-31T19:14:36'
- author: noqqe
  content: <p>Du bist so gut zu mir ;D</p>
  date: '2013-09-01T11:24:05'
- author: noqqe
  content: "<p>Ich war der Meinung der routed einen nur Location-Based zum n\xE4chstm\xF6glichen
    Mirror. Nicht Last-Bezogen.</p>"
  date: '2013-09-01T11:24:54'
- author: Michael
  content: "<p>apt-cacher-ng (siehe auch Wiki bei <a href=\"http://ubuntuusers.de\"
    rel=\"nofollow\">ubuntuusers.de</a>) setze ich hier ebenfalls ein, damit sich
    meine 4-5 Ubuntu-Computer im Haushalt m\xF6glichst viele Updates teilen. Funktioniert
    alles gut, einzig f\xFCr die mobilen Ger\xE4te m\xFCsste ich mal ein kleines Skript
    schreiben, damit sie sich direkt mit den Servern verbinden, wenn sie ausser Hause
    sind...</p>"
  date: '2013-09-03T07:18:42'
- author: senden9
  content: <p>"apt-cacher-ng" kannte ich schon. Das Thema Puppet/Massenkonfiguration
    welches du kurz angerissen hast finde ich allerdings interessant.</p>
  date: '2013-09-07T17:41:52'
- author: MKzero
  content: "Man kann, wenn man eh schon einen Proxy braucht(weil man z.B. keine direktes
    Routing ins Internet hat) auch einfach einen Squid o.\xE4. betreiben. Der kann
    entsprechend Cachen und die Clients brauchen auch nur die Proxy-Config f\xFCr
    Apt. "
  date: '2014-01-09T23:31:53.206764'
date: '2013-08-31T17:01:00'
tags:
- linux
- kernel
- cache
- linuxcontaier
- apt-cacher-ng
- administration
- devops
- apt
- puppet
- updates
- mirror
- lxc
- debian
title: A Story about Mirrors and Caches
---

[LinuxContainer](http://lxc.sourceforge.net/) sind für mich Wegwerfartikel.
Benutze sie als Sandboxes um neue Software auszuprobieren, bestimmte
Software zu Betreiben und um der
[Dependency-Hölle](https://en.wikipedia.org/wiki/Dependency_hell) zu
entkommen.  Auch dieser Blogpost wurde in so einer LXC-VM gebaut.

### Mirroring

Natürlich fällt bei (im Schnitt) 20-25 VMs zusätzlicher Verwaltungsaufwand an,
was ich bisher sehr gut mit Puppet erschlagen konnte. Ein anderes Problem sind
aber die OS Updates der Maschinen. Dazu nutze ich alle Monat
[clusterssh](https://github.com/duncs/clusterssh). Klar, wenn alle Maschinen
gleichzeitig Updates ziehen macht das weder meiner Internetanbindung noch dem
Mirror-Betreiber Spaß.

{{< figure src="/uploads/2013/08/w00t.jpg" >}}

Ich rechnete etwas hin und her, ob es sich wirklich lohnt einen eigenen
lokalen Mirror zu betreiben.  Für die richtige Architektur von
main/non-free/contrib braucht wohl ~150GB. Letztendlich rentiert sich das
von der Masse und Update-Synchronisierung erst hart spät. Ich tat es aber
dann doch, was mit dem [apt-mirror](http://apt-mirror.github.io/) einfach
und schnell ging.

### Caching

Es gibt Caching Software für Debian Repos! Nach einem Schwatz mit einem Kollegen
kam die Erkenntnis das alles zu syncen garnicht nötig war.
Auf diese Idee war ich nicht gekommen. Nichtmal an die Möglichkeit gedacht. Konkret
möchte man [apt-cacher-ng](https://www.unix-ag.uni-kl.de/~bloch/acng/) nutzen.

`apt-cacher-ng` agiert jetzt quasi als Proxy zwischen den Containern und dem
konfigurierten Debian Mirror.

``` bash
sudo aptitude install apt-cacher-ng
sudo /etc/init.d/apt-cacher-ng start
```

Möchte ein Client Updates herunterladen, bemüht sich apt-cacher-ng nach
Möglichkeit bereits heruntergeladene Pakete aus
`/var/cache/apt-cacher-ng/$mirror/debian` auszuliefern. Das spart
Bandbreite und Speicherplatz.  Clientseitig lässt sich die Nutzung des
Proxys auf zweierlei Varianten konfigurieren.

Entweder durch das `apt` HTTP Proxy Feature

    Acquire::http { Proxy "http://10.10.0.10:3142"; };

oder durch `sources.list` URLs

    deb http://10.10.0.10:3142/debian.cs.fau.de/debian stable main contrib non-free

Mir gefiel die erste Variante allerdings wesentlich besser. Ich kann es nicht
begründen, fühlt sich aber besser an.

### Puppet

Ausrollen konnte ich die Konfiguration gegen apt-cacher-ng einfach über das
Puppetlabs [apt](https://forge.puppetlabs.com/puppetlabs/apt) Modul.

``` puppet
class { 'apt':
  always_apt_update    => false,
  proxy_host           => '10.10.0.10',
  proxy_port           => '3142',
}

apt::source { 'wheezy-mirror':
  location    => 'http://debian.cs.fau.de/debian/',
  release     => 'wheezy',
  include_src => false,
  repos       => 'main contrib non-free',
}
```

Außerdem war dies der erste Post den ich mit der [Cherry G80-3000LSCDE-2](http://www.cherry.de/cid/b2b_keyboards_G80-3000.htm?rdeLocaleAttr=en&cpssessionid=SID-837EAC29-341CE33E&WT.mc_id=)
verfasste. Es ist schöner als vorher.
