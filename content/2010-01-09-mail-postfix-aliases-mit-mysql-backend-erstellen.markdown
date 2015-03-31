---
date: '2010-01-09 12:16:10'
layout: post
slug: mail-postfix-aliases-mit-mysql-backend-erstellen
status: publish
comments: true
title: Mail | Postfix-Aliases mit MySQL-Backend erstellen
alias: /archives/839
categories:
- Coding
- Linux
- ubuntuusers
tags:
- address
- alias
- aliases
- courier
- debian
- mail
- mysql
- postfix
- script
- skript
- sql
- ubuntu
---

Ich registrierte mich vor kurzem wiedermal bei einem etwas zwielichtigem Portal. Keine begründete Behauptung, es schien mir aber trotz allem so vorzukommen. Wie üblich loggte ich mich in meinen PHPMyAdmin ein und erstellte (um SpamEmails vorzubeugen) mithilfe meines MySQL-Backends von Postfix einen Alias. In einer Tabelle gesammelt liegen sämtliche aliase und deren Empfänger-Postfach.

```
address | goto
ubuntu@zwetschge.org | mail@zwetschge.org
spam2@zwetschge.org | mail@zwetschge.org
spam3@zwetschge.org | mail@zwetschge.org
spam4@zwetschge.org | mail@zwetschge.org

```


Es mag jetzt mit Sicherheit User geben die MySQL mit Postfix für unnötig halten, da die Steuerung über ConfigFiles ausreicht. Auf kurz oder lang gefällt mir die MySQL einfach besser. Einfach zuhandhaben. Flexibel. Schön. Das ständige eingelogge in HTpasswd, phpmyadmin-login und herumgeklicke war mir grad nur etwas zu blöd. Ich wollte ein kleines Skript basteln das mir das adden von Aliasen per CLI ermöglicht. In etwa so:

```
aliasadd <alias> <recepient>
```


Via echo lässt sich mysql (nach Authentifizierung) einen Befehl übergeben:

```
echo "use maildb; insert into aliases values ('$1', '$2');" | mysql -u <user> --password=<pass>
```


Damit wäre auch schon das gröbste geschafft. Zumindest das Einfügen. Ein Skript zeichnet allerdings mehr aus als nur die Aufgabe die es erledigen soll. Ein Skript muss zuverlässlich sicherstellen das die Aufgabe ausgeführt wurde und dies dem Benutzer nach Möglichkeit auch noch mitteilen.

```
echo "use maildb; select * from aliases where address = '$1';" | mysql -u <user> --password=<pass>
```


Die vorherige Zeile sieht eigentlich nur nach ob der eingegeben Alias wirklich in der Datenbank vorkommt. Freilich(wer findet 'freilich' eigentlich noch seltsam in Sätzen?) könnte ich jetzt noch nach Rückgabewerten mit $? Abfragen und ähnliche if-Vorraussetzungen einbauen. Aber für die 4-5 mal im Monat in denen ich es benutze wäre das übertrieben.

Fertig sieht das ganze dann wie folgt aus:

```
#!/bin/bash
echo "use maildb; insert into aliases values ('$1', '$2');" | mysql -u <user> --password=<pass>
echo "use maildb; select * from aliases where address = '$1';" | mysql -u <user> --password=<pass>
```


nochmal als Plaintext:
[/uploads/2009/09/6](/uploads/2009/09/6)
