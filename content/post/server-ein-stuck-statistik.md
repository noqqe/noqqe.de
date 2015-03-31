---
date: 2009-04-26T12:20:30+02:00
type: post
slug: server-ein-stuck-statistik
status: publish
comments: true
title: Server | Ein Stück Statistik
aliases:
- /archives/595
categories:
- General
- Linux
tags:
- emote
- fun
- gif
- kommentare
- random
- spam
- Statistik
- teamspeak
- Teeworlds
---

Hab mir mal die Zeit genommen ein paar Logfiles zu analysieren :)

Teeworlds-Server

Der Teeworldsserver der auf Zwetschge läuft wird zur Zeit immer mehr bespielt :) Logfile ist mittlerweile 15 MB gross :)
Mit dem Befehl:

```
for x in `cat /var/log/screen.log | grep "joined the game" | awk '{print $3 }' |
  sort | uniq`; do grep $x /var/log/screen.log | wc -l | tr -d 'n'; echo " :
  $x"; done | sort -rn
```
hab ich ein bisschen durch die Gegend gegreppt und die Spieler Rausgegreppt die am öftesten auf den Server gekommen sind :) Hier die Liste bei einem der schönsten Paste-Services:
[http://paste.pocoo.org/show/114322/](http://paste.pocoo.org/show/114322/)

Random-Animations-Pool

Zum Teil auch aus Langeweile lass ich mir alle Heilige Zeit mal per CronJob die Statistik für die zufällig ausgewählten Bilder per Mail zuschicken. Das Skript dazu hab ich bei [CodeCocktail gepostet](http://codecocktail.wordpress.com/2009/02/01/zufallszahlen-mit-der-shell-bash/) Der Zufall ist dabei sehr ... parteiisch find ich:
Daily Emote Score for 26-04-2009-04:01:01

```
200 : http://zwetschge.org/emotes/emote32.gif
216 : http://zwetschge.org/emotes/emote20.gif
216 : http://zwetschge.org/emotes/emote5.gif
220 : http://zwetschge.org/emotes/emote58.gif
220 : http://zwetschge.org/emotes/emote73.gif
221 : http://zwetschge.org/emotes/emote11.gif
221 : http://zwetschge.org/emotes/emote49.gif
221 : http://zwetschge.org/emotes/emote64.gif
222 : http://zwetschge.org/emotes/emote8.gif
225 : http://zwetschge.org/emotes/emote63.gif
227 : http://zwetschge.org/emotes/emote74.gif
229 : http://zwetschge.org/emotes/emote55.gif
230 : http://zwetschge.org/emotes/emote2.gif
231 : http://zwetschge.org/emotes/emote56.gif
232 : http://zwetschge.org/emotes/emote22.gif
232 : http://zwetschge.org/emotes/emote26.gif
232 : http://zwetschge.org/emotes/emote4.gif
233 : http://zwetschge.org/emotes/emote29.gif
233 : http://zwetschge.org/emotes/emote50.gif
233 : http://zwetschge.org/emotes/emote54.gif
234 : http://zwetschge.org/emotes/emote57.gif
234 : http://zwetschge.org/emotes/emote66.gif
235 : http://zwetschge.org/emotes/emote13.gif
235 : http://zwetschge.org/emotes/emote28.gif
236 : http://zwetschge.org/emotes/emote0.gif
237 : http://zwetschge.org/emotes/emote46.gif
237 : http://zwetschge.org/emotes/emote65.gif
238 : http://zwetschge.org/emotes/emote71.gif
238 : http://zwetschge.org/emotes/emote79.gif
239 : http://zwetschge.org/emotes/emote12.gif
239 : http://zwetschge.org/emotes/emote14.gif
239 : http://zwetschge.org/emotes/emote1.gif
239 : http://zwetschge.org/emotes/emote37.gif
239 : http://zwetschge.org/emotes/emote75.gif
240 : http://zwetschge.org/emotes/emote19.gif
240 : http://zwetschge.org/emotes/emote61.gif
241 : http://zwetschge.org/emotes/emote24.gif
241 : http://zwetschge.org/emotes/emote59.gif
241 : http://zwetschge.org/emotes/emote70.gif
243 : http://zwetschge.org/emotes/emote77.gif
244 : http://zwetschge.org/emotes/emote18.gif
246 : http://zwetschge.org/emotes/emote16.gif
246 : http://zwetschge.org/emotes/emote40.gif
247 : http://zwetschge.org/emotes/emote43.gif
248 : http://zwetschge.org/emotes/emote9.gif
249 : http://zwetschge.org/emotes/emote51.gif
249 : http://zwetschge.org/emotes/emote62.gif
252 : http://zwetschge.org/emotes/emote21.gif
252 : http://zwetschge.org/emotes/emote42.gif
254 : http://zwetschge.org/emotes/emote33.gif
254 : http://zwetschge.org/emotes/emote34.gif
254 : http://zwetschge.org/emotes/emote35.gif
254 : http://zwetschge.org/emotes/emote72.gif
256 : http://zwetschge.org/emotes/emote41.gif
256 : http://zwetschge.org/emotes/emote44.gif
256 : http://zwetschge.org/emotes/emote45.gif
256 : http://zwetschge.org/emotes/emote69.gif
257 : http://zwetschge.org/emotes/emote27.gif
257 : http://zwetschge.org/emotes/emote60.gif
258 : http://zwetschge.org/emotes/emote30.gif
259 : http://zwetschge.org/emotes/emote39.gif
259 : http://zwetschge.org/emotes/emote3.gif
259 : http://zwetschge.org/emotes/emote53.gif
259 : http://zwetschge.org/emotes/emote76.gif
260 : http://zwetschge.org/emotes/emote23.gif
260 : http://zwetschge.org/emotes/emote36.gif
262 : http://zwetschge.org/emotes/emote25.gif
262 : http://zwetschge.org/emotes/emote80.gif
263 : http://zwetschge.org/emotes/emote52.gif
265 : http://zwetschge.org/emotes/emote67.gif
266 : http://zwetschge.org/emotes/emote31.gif
266 : http://zwetschge.org/emotes/emote6.gif
268 : http://zwetschge.org/emotes/emote10.gif
269 : http://zwetschge.org/emotes/emote15.gif
269 : http://zwetschge.org/emotes/emote17.gif
270 : http://zwetschge.org/emotes/emote68.gif
271 : http://zwetschge.org/emotes/emote38.gif
274 : http://zwetschge.org/emotes/emote78.gif
283 : http://zwetschge.org/emotes/emote7.gif
293 : http://zwetschge.org/emotes/emote48.gif
324 : http://zwetschge.org/emotes/emote47.gif
```

Spam-Kommentare

Zur Zeit explodiert (trotz Askimet) mein Spam-Kommentar Counter... und zwar immer nur bei dem Post zum[Thema Teeworlds-Server](/?p=459) Fragt mich bitte nicht an was das liegt... aber nun Gut... der Counter zählt mittlerweile rund 1200 Spams auf diesen Beitrag O_o. Wenn jemand vorschläge hat wie das zu vermeiden ist... immer her damit!

Blog

Auch der Blog selbst hat nach dem Relaunch (damals 15.000) schon wieder 3.500 Zugriffe und ganze 8 Feed-Reader :D. Find ich eigentlich ganz in Ordnung wenn man bedenkt das mir spontan nur 4 Leute einfallen die ich persöhnlich kenne die ihn lesen, und zum anderen ich hier nicht wirklich Intressantes Zeugs Blogge :)

Teamspeak-Server

Mittlerweile 15 GB Traffic sent by Server und ca 5 GB recieved by Server :) Auch viel mehr als ich anfangs dachte.

Greez, Flo

