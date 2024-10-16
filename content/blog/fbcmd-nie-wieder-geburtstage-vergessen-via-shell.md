---
aliases:
- /blog/2011/12/11/fbcmd-nie-wieder-geburtstage-vergessen-via-shell
- /archives/1827
comments:
- author: Lorlen
  content: "<p>Genial!! Damit nicht auff\xE4llt dass man allen immer dasselbe schreibt,
    m\xFCsste man das jetzt nur noch so ver\xE4ndern, dass aus zehn verschieden Geburtstagsnachrichten
    in Kombination mit verschiedenen Smileys am Ende, und dem Namen des Geburtstagskindes
    eine individuelle Nachricht gebastelt wird.</p>"
  date: '2011-12-12T02:12:46'
- author: dAnjou
  content: "<p>Ich hoffe, dass dieses Beispiel nur exemplarisch gew\xE4hlt wurde.
    Denn diese Sache ist schon ziemlich bezeichnend f\xFCr Facebook und seine Nutzer.</p>"
  date: '2011-12-12T04:04:11'
- author: Higgi
  content: "<p>@dAnjou, </p><p>das ist typisches Facebookverhalten. Meine Freundin
    und ich haben kurz hintereinander Geburtstag. Wir haben mal einen kleinen Test
    gemacht. Ich habe meinen Geburtstag drin stehen lassen, sie ihren gel\xF6scht.
    Ich hatte ca 40 Nachrichten von Leuten, die ich selten sehe/spreche und sie hatte
    3! Nachrichten von anderen. Text war trotzdem immer \"Alles gute zum Geburstag\"
    oder \xE4hnliches.</p>"
  date: '2011-12-13T01:06:58'
date: '2011-12-11T18:45:18'
tags:
- development
- web
- cronjob
- automatisch
- pinnwand
- facbeook
- geburtstag
- fbcmd
- automagisch
- wallpost
- fb
- birthday
- versenden
- reminder
- debian
- bash
title: 'fbcmd | Nie wieder Geburtstage vergessen via Shell. '
---

Die nachfolgende Beschreibung eines technischen Vorgangs würde die Mehrheit
der Gesellschaft wahrscheinlich als soziologisch fragwürdig abstempeln.
Jedoch beschreibe ich den Hergang trotzdem und gerade deswegen.

{{< figure src="/uploads/2011/12/112817814_facebook-birthday1.jpg" >}}

Ich weiß gar nicht mehr wie genau ich auf
[fbcmd](http://fbcmd.dtompkins.com/) gekommen bin.  Im Zweifel über einen
XML basierten Medienkanal. Jedenfalls ist fbcmd ein äußerst schönes Tool um
die gängigen Informationen zum eigenen Facebook Account auf der
Kommandozeile abzufragen. Dazu bietet es wahnsinnig trickreiche Queries und
"Vergruppungen" der Facebook Bekanntschaften. Alles dazu
[hier](http://fbcmd.dtompkins.com/syntax) und besonders
[hier](http://fbcmd.dtompkins.com/parameters/flist).

Somit lässt sich wunderbar ein automatisches "Alles Gute zum Geburtstag!"
bauen. Vorraussetzung ist hierfür eine funktionierende [Installation von fbcmd](http://fbcmd.dtompkins.com/installation).

Der von fbcmd vorgeschlagene Query um eine Geburstagsnachricht an die
Pinnwand von jenen zu senden, deren Geburtstag sich heute jährt lautet:

```
fbcmd WALLPOST =bday 'Alles Gute zum Geburstag!'
```

Das lässt sich natürlich wunderbar in einen Cronjob verbauen, der einmal
täglich um 15:00 eben diesen Query ausführt:

```
0 15 * * *   fbcmd WALLPOST =bday 'Alles Gute zum Geburstag!' > /dev/null
```

Weil ich aber wissen möchte, wem mein Rechner alles in meinem Namen zum
Geburtstag gratuliert hab ich das noch leicht modifiziert und lasse mich
via Mail darüber benachrichtigen:

```
0 15 * * *   fbcmd WALLPOST =bday 'Alles Gute zum Geburstag!'  | grep -v "^No Friends With Birthday Matches$" | mail -s "fbcmd Gratulation" user@domain.de
```

Eigentlich ist der Titel des Posts gar nicht richtig. Man vergisst Sie
trotzdem. Aber ein Device erledigt die Arbeit für einen :)
