---
title: apt
date: 2012-05-21T08:18:00
tags: 
- Software
- apt
---

Neuen Key aus untrusted Repo importieren

    gpg --keyserver pgpkeys.mit.edu --recv-key 1285491434D8786F
    gpg -a --export 1285491434D8786F | apt-key add -
    aptitude update
    gpg --list-keys
    gpg --list-sigs 34D8786F

    $ gpg --list-sigs 34D8786F
    pub   4096R/34D8786F 2012-03-02
    uid                  Dell Inc., PGRE 2012 (PG Release Engineering Build Group
    2012) <PG_Release_Engineering@Dell.com>
    sig          23B66A9D 2012-03-02  [User ID not found]
    sig 3        92F0FC09 2012-03-02  [User ID not found]
    sig          1019CED6 2012-03-02  [User ID not found]
    sig          86A10CC0 2012-03-02  [User ID not found]
    sig 3        34D8786F 2012-03-02  Dell Inc., PGRE 2012 (PG Release Engineering
    Build Group 2012) <PG_Release_Engineering@Dell.com>
    sub   4096R/79DF80D8 2012-03-02
    sig          34D8786F 2012-03-02  Dell Inc., PGRE 2012 (PG Release Engineering
    Build Group 2012) <PG_Release_Engineering@Dell.com>

Key manuell von Site herunterladen

    wget http://apt.example.com/squeeze.asc
    apt-key
    apt-key add squeeze.asc

## Wenn bei Ubuntu Sources nichts mehr hilft...

Wegen eventuell diesem Fehler

    W: Bei der Überprüfung der Signatur ist ein Fehler aufgetreten. Die Paketquelle ist nicht aktuell, die bisherigen Index-
    Dateien werden weiter verwendet. GPG-Fehler: http://de.archive.ubuntu.com lucid Release Die folgenden Signaturen waren
    ungültig: BADSIG 40976EAF437D05B5 Ubuntu Archive Automatic Signing Key <ftpmaster@ubuntu.com>
    W: Fehlschlag beim Holen von http://de.archive.ubuntu.com/ubuntu/dists/lucid/Release
    W: Einige Indexdateien konnten nicht heruntergeladen werden, sie wurden ignoriert oder alte an ihrer Stelle benutzt.

dann

    sudo apt-get clean
    cd /var/lib/apt
    sudo mv lists lists.old
    sudo mkdir -p lists/partial
    sudo apt-get clean
    sudo apt-get update

Gefunden bei [askubuntu.com](http://askubuntu.com/questions/85641/how-do-i-deal-with-unauthenticated-sources-errors-in-the-software-center)
