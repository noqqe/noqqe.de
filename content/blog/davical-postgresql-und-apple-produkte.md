---
comments:
- author: Pacer
  content: "<p>Hi,</p><p>thanks for your description. I followd it, but needed to
    enter the complete URL into the field Principal-URL (http://[host]/davical/pacer/contats/)</p><p>When
    I left http://[host] away, I always got an error in the debug log, saying that
    this URL is not supported. With the full URL I don't get any errors.\_</p><p>BUT:
    I don't see any contacts from the CardDAV server in my Contacts-App. Don't know
    why.</p><p>To make it even more strange, I see the apache request in the access_log:</p><p>212.91.225.xx
    - pacer [27/Apr/2012:20:12:29 +0200] \"PROPFIND /davical/caldav.php/pacer/contacts/
    HTTP/1.1\" 207 1396 \"-\" \"iOS/5.1 (9B176) dataaccessd/1.0\"</p><p>... which
    looks ok to me.</p><p>I looked for another CardDAV Client to test if the server
    works. I tested Evolution in openSuse, and it quickly and successfully read my
    two test contacts. But evolutions access_log entries look a little different:</p><p>212.91.225.xx
    - pacer [27/Apr/2012:20:17:35 +0200] \"PROPFIND /davical/caldav.php/pacer/contacts/
    HTTP/1.1\" 207 878 \"-\" \"Evolution/3.2.1\"212.91.225.xx - pacer [27/Apr/2012:20:17:35
    +0200] \"PROPFIND /davical/caldav.php/pacer/contacts/ HTTP/1.1\" 207 338 \"-\"
    \"Evolution/3.2.1\"212.91.225.xx - pacer [27/Apr/2012:20:17:35 +0200] \"PROPFIND
    /davical/caldav.php/pacer/contacts/ HTTP/1.1\" 207 878 \"-\" \"Evolution/3.2.1\"212.91.225.xx
    - pacer [27/Apr/2012:20:17:36 +0200] \"GET /davical/caldav.php/pacer/contacts/000000006ECBCBE9F133AD46925453EB1E78012B64002000.vcf
    HTTP/1.1\" 200 233 \"-\" \"Evolution/3.2.1\"212.91.225.xx - pacer [27/Apr/2012:20:17:36
    +0200] \"GET /davical/caldav.php/pacer/contacts/000000006ECBCBE9F133AD46925453EB1E78012B24002000.vcf
    HTTP/1.1\" 200 319 \"-\" \"Evolution/3.2.1\"</p><p>Do you have an idea?</p>"
  date: '2012-04-27T18:25:11'
date: '2012-04-09T18:05:00'
tags:
- web
- apple
- contacts
- addressbook
- app
- caldav
- carddav
- ical
- linux
- davical
- postgres
title: DaviCal, PostgreSQL und Apple Produkte
---

Vor ewigen Zeiten hab ich mir mal ein eGroupware installiert. Es war schrecklich
und deshalb hab ich es durch etwas weniger Schreckliches ersetzt.
[DaviCal](http://davical.org).

Nicht das ich spontan Lust gehabt hätte sowas wie Kontakte und so mal zu
Ordnen.  Vielmehr hat mich die PostgreSQL Datenbank intressiert die dazu
nötig ist. Installation und Konfiguration ist so schön dokumentiert das
ich das hier nicht widerholen brauche. Die Integration in Thunderbird
verlief ohne Probleme. Dann kamen meine beiden Geräte aus dem Hause Crapple
mit der Synchronisation der Kontakte dran.

## iPhone

Mein iPhone wollte den hingeworfenen Remote für Contacts nicht einfach so
hinnehmen. Ewiges "Verifying" durch den Installationsassistenten und jedesmal
gefühlte 5-10 Minuten später hat er mich dann abgewiesen. Im Apache Log ist
davon nichts aufgetaucht. Nach bisschen wälzen in der Dokumentation hier
und da:

> CardDav-Accounts only could set up via iPhone configuration utility. It had to
> set up a principal url (e.g. https://domain.tld/caldav.php/username/contacts)
> for a working account.

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

``` xml
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

Hurra. Ohne das kaputte Interface kann mein Address Book problemlos mit
meinem neuen DaviCal Server reden.
