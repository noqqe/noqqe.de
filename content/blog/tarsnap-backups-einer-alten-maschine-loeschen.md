---
categories:
- administration
- bash
- bsd
- crypto
- opensource
- osbn
- planetenblogger
- shell
- ubuntuusers
date: 2015-06-14T18:25:54+02:00
tags:
- Backup
- Tarsnap
- remove
title: Tarsnap Backups einer alten Maschine löschen
---

Seit einiger Zeit verwende ich [tarsnap](https://tarsnap.com) als Backup auf
meinen VMs. Die Crypto, der Preis, das System. Gefällt mir. Nun habe ich vor ein
paar Wochen ein Stück Hardware gegen eine kleine VM ausgetauscht. Die Backups
auf Tarsnap habe ich vorher noch nicht gelöscht.

{{< figure src="/uploads/2015/06/backup.jpg" >}}

Da die Maschine nicht mehr existent ist, die Backups aber maschinen-gebunden,
muss ich diese noch loswerden um nicht jeden Monat dafür mit zuzahlen. Nicht
dass es sonderlich teuer wäre, aber trotzdem.

Dafür erforderlich ist natürlich der Key der Maschine.

Die Config erzeugen ist der erste Schritt.

```
cachedir /tmp/tarsnapc
keyfile /tmp/key
humanize-numbers
print-stats
totals
normalmem
```

Danach ein `mkdir /tmp/tarnsapc` und den Key nach `/tmp/key` kopieren. Bevor
das eigentliche löschen losgeht, muss das Cacheverzeichnis mit den Metadaten
befüllt werden:

``` bash
$ tarsnap --config conf --fsck
Phase 1: Verifying metadata validity
Phase 2: Verifying metadata/metaindex consistency
Phase 3: Reading chunk list
Phase 4: Verifying archive completeness
  Archive 1/56...
  Archive 2/56...
  Archive 3/56...
  [...]
  Archive 56/56...
Phase 5: Identifying unreferenced chunks
```

Jetzt können die Backups des alten Servers die sich noch bei Tarsnap befinden
angeschaut werden.

``` bash
$ tarsnap --config conf --list-archives
20150402-001026-daily-/home
20150404-001026-daily-/home
20150316-001024-weekly-/etc
20150401-001024-monthly-/etc
20150330-001026-weekly-/root
[...]
```

Um die Backups zu entfernen, eine einfache `for`-Loop über alle gefundenen

``` bash
for x in $(tarsnap --config conf --list-archives); do
  tarsnap --config conf -d -f $x
done
```

Geld gespart.
