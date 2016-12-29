---
comments:
- author: Stefan
  content: 1s/^Ich bei/Bei/g LG Stefan
  date: '2016-10-27T17:10:39.577637'
- author: Anonymous
  content: "Vielleicht ist DN42.net f\xFCr dich interessant, ein privates Netz f\xFCr
    Leute die BGP nicht in einem \xF6ffentlichen Netz testen wollen oder k\xF6nnen.\n\nDas
    Wiki dort ist ausf\xFChrlich was Konfiguration der einzelnen BGPd's angeht."
  date: '2016-10-31T16:19:01.773276'
date: '2016-10-27T13:32:00'
tags:
- openbsd
- network
- spam
- administration
- bgp
- opensource
- bsd
- mail
title: "Spamlisten \xFCber BGP"
---

Ich bei meiner Suche nach Antispam Lösungen stolperte ich über
[http://bgp-spamd.net](http://bgp-spamd.net). Jemand den ich vergessen
habe, hatte mir davon schonmal erzählt und eigentlich macht das Sinn.

Wie der Zufall so will habe ich überhaupt keine Ahnung von BGP. Leider.
Solche Sachen sind aber auch schwer zu lernen, wenn man nicht gerade bei
einem ISP arbeitet :( Wenn dafür jemand Tools/Anleitungen hat die
lesenswert sind, bitte her damit.

Wie auch immer. Ich konfigurierte erstmals einen BGP Daemon. Die
idiotensichere Anleitung auf obiger Website macht es einem auch schwer das
falsche zu tun. Die Listen bzw. IPs über ein dafür gedachtes
Protokoll zu übertragen und nicht alle 20 min eine ganzes `gzip` File über
HTTP zu ziehen ist schon irgendwie charmant. Eigentlich ist es auch nichts
anderes als 3 Server in einer Config zu hinterlegen die dann ihr eigenes
Protokoll sprechen.

Mittels folgendem Befehl lassen sich die übertragenen "Routen" dann
anzeigen.

```
bgpctl show rib community 65066:666
```

Im Endeffekt gebe ich nun trotzdem Greylisting noch eine Chance, denn es
sprechen sich nicht nur Blacklisted IPs herum. In der Standardkonfiguration
aus `bgp-spamd` werden auch whitelisted IPs in einen `pf` Table
geschrieben.

```
match from group spamd-bgp community $spamdAS:42  set pftable "bgp-spamd-bypass"
```

Somit kommen Mails von der Bahn oder Paypal usw. trotzdem gleich am
Greylisting vorbei, böse Hosts werden weggeblockt und der Rest darf nach
einer kurzen Zeit im Greylisting nochmal wiederkommen.