---
title: "dd"
date: 2020-06-23T10:25:04+02:00
tags:
- Software
---

All die Konfusion bzgl `dd`, endlich mal sortiert.

<!--more-->

## Progress

Man kann auch den Fortschritt nativ anzeigen (bei GNU dd)

``` none
# dd if=/dev/zero of=tempswap bs=1M count=1024 status=progress
1010827264 bytes (1.0 GB, 964 MiB) copied, 13 s, 77.8 MB
```

Ich bin mir nicht sicher warum ich das all die Jahre nicht wusste.

## Random

Wenn man unbedingt Zufällige Daten braucht unter Linux lieber `/dev/urandom`
nehmen statt `/dev/random` da dort die Entropie nicht ausgeht.

## Unterschiedliche Filegrößen

Zum Beispiel 10Mb bis 1G

``` fish
for x in 10240 102400 1024000
  dd if=/dev/zero of=/tmp/$x.bin bs=1024 count=$x
end
```

## truncate

Ist eigentlich ein eigenes Tool das viel schneller ist und bessere
human-readable Filegrößen beherrscht.

``` fish
for x in 1M 5M 10M 100M 1G 10G ; truncate -s $x /tmp/$x.bin ; end
```
