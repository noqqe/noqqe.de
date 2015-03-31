---
date: 2011-04-14T18:55:10+02:00
layout: post
slug: statistical-statistiken-visualisieren-im-terminal
status: publish
comments: true
title: statistical | Statistiken visualisieren im Terminal
alias: /archives/1603
categories:
- Bash
- Coding
- Debian
- General
- git
- PlanetenBlogger
- Spam
- Ubuntu
- Web
tags:
- alice
- bash
- bob
- git
- konsole
- linus
- script
- shell
- Statistik
- Statistiken
- terminal
---

Mir erschien es einen kurzen Moment lang für sinnvoll ein kleines Shell Tool zu haben, welches mir aus einer Liste von Key:Value Paaren eine Balkenstatistik baut und visualisiert. Wie in etwa **$ statistical john:12 alice:5 linus:7 bob:1**. Mir gefiel die Idee einfach alles mögliche in meinem Terminal ansehen zu können.

{{< figure src="/uploads/2011/04/1923_a9bc.jpeg" >}}

Relativ schnell stieß ich aber an eine Grenze. Diese hieß "Windowsize". Ich konnte nicht ohne bedenken eine Schleife die die Value Werte zählt bauen, die dementsprechend viele Zeichen anhängt. Denn bei Zahlen >10000 wird das ziemlich unlesbar :)


     while [ $COUNTER -lt $VALUE ]; do
            ((COUNTER++))
            echo -n "$OUTPUTCHAR"
            if [ $COUNTER -ge $VALUE ]; then
                echo
            fi
        done


Ich brauchte ein Schema, welches alle Werte einließt und eine skalierbare Basis für alle Werte schafft. Ich entschied mich für eine simple Lösung.


    while [ ${FACTORCOUNT} -lt $(( ${#MAXVALUE} - 2 )) ]; do
    FACTOR="${FACTOR}0"
        ((FACTORCOUNT++))
    done



Letztenendes kam dann folgendes Verhalten bei meinem Key:Value Statistik Script raus. Ich mags.


    # Beispiel
    $ statistical john:433 alice:49 linus:12 bob:231
    john    |###########################################
    alice   |####
    linus   |#
    bob     |#######################


Damit lassen sich sogar teilweise sinnvolle Sachen produzieren. Zum Beispiel die Anzahl der Commits innerhalb eines Git-Repositories. Ich habe hier als Beispiel mal [bash-it](http://github.com/revans/bash-it) aufgeführt:


    for a in $(git shortlog -sn --all | cut -f2 | cut -f1 -d' '); do echo -n "$a:" ; git log $LOGOPTS --all --numstat --format="%n" --author=$a | cut -f3 | sort -iu | wc -l; done  | statistical
    Mark          |##################
    Robert        |#########################################################################
    Florian       |##############
    Jesus         |######
    John          |##############
    Rich          |########
    Piotr         |###
    Travis        |####
    Fedyashev     |##
    zerobearing2  |####
    Andy          |###
    Daniel        |####
    Jeff          |##
    Karl          |##
    Robert        |#########################################################################
    Sirupsen      |##



Sollte jemand Interesse daran hegen, das Skript auch mal auszuprobieren es befindet sich wie immer auf Github: [http://github.com/noqqe/statistical](http://github.com/noqqe/statistical)
