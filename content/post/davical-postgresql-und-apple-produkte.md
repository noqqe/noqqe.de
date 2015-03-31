---
layout: post
title: "DaviCal, PostgreSQL und Apple Produkte"
date: 2012-04-09T20:05:00+02:00
comments: true
categories:
- Linux
- Web
- Apple
tags:
- davical
- caldav
- carddav
- postgres
- ical
- AddressBook
- App
- Contacts
---

Vor ewigen Zeiten hab ich mir mal ein eGroupware installiert. Es war schrecklich
und deshalb hab ich es durch etwas weniger Schreckliches ersetzt.
[DaviCal](http://davical.org).

Nicht das ich spontan Lust gehabt hätte sowas wie Kontakte und so mal zu Ordnen.
Vielmehr hat mich die PostgreSQL Datenbank intressiert die dazu nötig ist.
Installation und
Konfiguration ist so schön dokumentiert das ich das hier nicht widerholen
brauche. Die Integration in Thunderbird verlief ohne Probleme. Dann kamen meine
beiden Geräte aus dem Hause Crapple mit der Synchronisation der Kontakte dran.

## iPhone

Mein iPhone wollte den hingeworfenen Remote für Contacts nicht einfach so
hinnehmen. Ewiges "Verifying" durch den Installationsassistenten und jedesmal
gefühlte 5-10 Minuten später hat er mich dann abgewiesen. Im Apache Log ist
davon nichts
aufgetaucht. Nach bisschen wälzen in der Dokumentation hier und da:


{% blockquote Carddav Clients, http://wiki.davical.org/w/CardDAV/Clients wiki.davical.org %}
CardDav-Accounts only could set up via iPhone configuration utility. It had to
set up a principal url (e.g. https://domain.tld/caldav.php/username/contacts)
for a working account.
{% endblockquote %}

Das [iPhone configuration utility](http://www.apple.com/support/iphone/enterprise/) also.
So ganz leicht fand ich das damit nicht... Ich hab aber gelernt das man super
Debuggen kann über die eingebaute Consolen Funktion in dem Tool. Im iPhone wird
man nach dem ServerPath gefragt. Den ich mit

> http://cal.n0q.org/caldav.php/noqqe/addressbook

Mein Fehler war eben genau die Angabe des Protokolls. Ich meine sorry.
Apple probiert hier gefühlte 1,3 mio (nichtmal wirklich RFC spezifizierte) Orte
an denen meine Kontakte serverseitig wohl liegen könnten durch und am Ende
können die nichtmal ein führendes http:// wegparsen? Egal. Fehler sah wie folgt
aus:

> http://http//cal.n0q.org/caldav.php/noqqe/addressbook


## Macbook

Hier nun der wirklich lustige Teil. Das Address Book des Macbooks. Selbes
"Verifying" Spielchen wieder.

Nachdem ich rausgefunden habe wie ich das intuitive Interface der App wirklich
benutzen soll kam nur folgendes auf meinem Webserver an:

```
77.191.14.156 - - [09/Apr/2012:19:35:53 +0200] "PROPFIND /caldav.php/noqqe/addressbook:0(null)/ HTTP/1.1" 401 401 "-" "Address%20Book/883 CFNetwork/454.12.4 Darwin/10.8.0 (i386) (MacBook2%2C1)"
77.191.14.156 - - [09/Apr/2012:19:35:53 +0200] "PROPFIND /caldav.php/noqqe/addressbook:0(null)/ HTTP/1.1" 400 477 "-" "Address%20Book/883 CFNetwork/454.12.4 Darwin/10.8.0 (i386) (MacBook2%2C1)"
```

Herzlichen Glückwunsch. Woher kommt dieses ":0(null)" ? Ich habe den Serverport angegeben:

{{< figure src="/uploads/2012/04/addressbook.png" >}}

Ich habe mehr als 10 mal versucht das Interface korrekt auszufüllen, aber hatte
keine Chance. Nach einmaligem Anlegen konnte ich den Serverpath nicht mehr
editieren und wenn ich Ihn von vornherein eingetragen habe bekam ich immer das
oben beschriebene ":0(null)".

Ich hab mir schliesslich die Config von Addressbuch mühsam ergrept und gefunden:

```
~/Library/Application Support/AddressBook/Sources/FC24F051-B1B2-4286-9A28-5830207F0DA2/Configuration.plist
```

Dort gibt es unter anderem folgende Paramteri (bereits korrigiert):


```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>_className</key>
    <string>PHXCardDAVSource</string>
    <key>disabled</key>
    <integer>0</integer>
    <key>haveWriteAccess</key>
    <integer>0</integer>
    <key>isSharedABAccount</key>
    <integer>1</integer>
    <key>name</key>
    <string>cal.n0q.org</string>
    <key>refreshInterval</key>
    <integer>0</integer>
    <key>serverSupportsSearch</key>
    <integer>0</integer>
    <key>servername</key>
    <string>http://cal.n0q.org:80/caldav.php/noqqe/addressbook</string>
    [...]
</dict>
</plist>
```

Hurra. Ohne das kaputte Interface kann mein Address Book problemlos mit meinem neuen DaviCal Server reden.
