---
comments:
- author: Anonymous
  content: "Moin erstmal ;D Bin eigentlich so ein Typ, der gerne mit Schubladen arbeitet.
    Die Probleme hier kenne ich zu gen\xFCge, meist kriegt man in den benannten Forum,
    lass die Finger davon und such dir was anderes. Aber der Punkt \"Der resignierende
    Senior\" trifft ein wenig auch auf mich zu. Bin fr\xFCher ein weniger aktiver
    gewesen, was probleml\xF6sungen angeht, jetzt mache ich eigentlich nur das was
    ich muss. \n\nAber sch\xF6ner Artikel ;D"
  date: '2014-04-19T21:09:20.874128'
- author: maltris
  content: "Sch\xF6ner Artikel, hat Spa\xDF gemacht zu lesen. Kann mich aber keiner
    Kategorie wirklich zuordnen. Mache von jedem ein bisschen. :D"
  date: '2014-04-20T07:31:04.566934'
- author: Spion
  content: Systemadministration, ist das ansteckend?!?
  date: '2014-04-20T10:18:35.255794'
- author: "Kr\xFCtschi"
  content: "Ich bin Dienstleister(Hotline) f\xFCr ein gr\xF6sseres Fach-Programm.
    Wer Wartung hat kann uns anrufen wir helfen. Bereits zwei mal hat sich heraus
    gestellt, das der Hausmeister zum Administrator wurde. \nBeide male hatte der
    Ur-Admin gek\xFCndigt. Beide Hausmeister wurden gefragt ob sie sich mit Computer
    auskennen und beide haben einfach Ja gesagt.\nBeide haben einfach nie etwas am
    System gemacht. Der eine hat sich beim Verkaufsgespr\xE4ch entlarft. Weil er nie
    was gesagt hat, sich extrem Vorsichtig gabe und alles von uns erledigt gehabt
    haben wollte. Er ist offenbar beim Nachfragen welches Server-Betriebssystem eigentlich
    derzeit l\xE4uft auf die Fressen gefallen.\n\nDer andere ist auf gefallen, weil
    er es selbst mit Anleitung nur umst\xE4ndlich geschaft hat eine Remotesession
    ein zu richten und dann nicht mal wusste wie man in Windows ein Netzlaufwerk verbindet
    und auch die Systemsteuerung nicht gefunden hatte."
  date: '2014-04-25T15:45:50.832633'
date: '2014-04-19T16:52:00'
tags:
- operations
- dienstleister
- sysops
- provider
- administration
- devops
- it
- unix
- systemadministration
- team
- systemadministrator
title: Administration von Systemen
---

Admins. In der System Administration treffen Einstellungen aufeinander.
Geschmacksrichtungen. Unterschiedliche Wissensstände. Glaubensrichtungen.
Interessen. Vorlieben.  Wer tut in einem Ops-Team was und warum. Was folgt ist
ein Gedankenmodell.

### Solutions

Lösungen. Eigentlich ist es das, worum es im Operations geht. Dem nachkommen
von Anforderungen an ein System. Ob neue Komponente, Sicherheitslücke oder
Setup-Aufbau, darauf lassen sich all diese Dinge herunterbrechen.

Die angewendeten Lösungen (Changes) lassen sich in maximal 2 der nachfolgenden
Kategorien einordnen. Ähnlich wie beim
[CAP-Theorem](https://en.wikipedia.org/wiki/CAP_theorem).

* Einfach: Der Change ist einfach, schnell umgesetzt und auch für Kollegen
  verständlich
* Korrekt: Die Lösung kann je nach Bedarf zukünftig erweitert werden, ist
  sicher und auf Korrektheit überprüft
* Risikofrei: Das entstehende Risiko für die Umgebung bei Anwendung der neuen
  Lösung ist möglichst gering.

### Practices

#### "Es ging nicht anders" Man hat keine Ahnung vom Thema. Die Lösung
entspringt Google, einem Unix-Forum Post von 1997 oder im besten Fall einem
Server Fault Thread mit 1,5 Upvotes.

{{< figure src="/uploads/2014/04/dirtyworkaround.gif" >}}

Diese Praktik erzeugt meistens wenig Aufwand, geht schnell. Hinterlässt aber
einen Gewissen Nachgeschmack. Fühlt sich dreckig an und sieht manchmal auch so
aus. Meist auch die Variante in der sich andere Teamkollegen beim
Drüberstolpern nachher die Augen auswaschen müssen.

#### "Best Practice"

Größtmögliche Korrektheit. Dieser Ansatz entspringt meistens entweder
persönlichem Interesse des Admins an der Software oder entsprechende
Schulung+Einarbeitung in das Thema.

{{< figure src="/uploads/2014/04/correct2.gif" >}}

Was in Erweiterbarkeit, leichter Administration und Stabilität endet muss mit
dem Eingehen von Risiken und Zeit bezahlt werden. Man ändert nicht einfach mal
so die Templates der Default Role im LDAP oder die Architektur des
Datenbankclusters.  Danach hat man aber allerdings erstmal Ruhe und das
angenehme Gefühl die Welt ein Stück besser gemacht zu haben.

#### "Hart overengineert"

Wenn auch das kleinste Ein-Server-Homepage-Setup von der Metzgerei um die Ecke
bis zum Anschlag puppetisiert, hoch-skalierbar und für alle Erweiterungen
bereit eingerichtet ist.

{{< figure src="/uploads/2014/04/witchcraft.gif" >}}

Der overengineerte Ansatz ist stabil. Allerdings schwierig im Team zu
maintainen und beansprucht zu viel Zeit. Was in hohen Einmalkosten, verlorener
Mühen und Nerven endet um für Eventualitäten gewappnet zu sein die wohl niemals
eintreten werden.

### Roles

Im Lauf eines Admin-Berufslebens begegnen einem auch allerhand Charaktere.  Ich
würde nicht von mir behaupten, ein Urgestein der Administration zu sein,
dennoch lernt man so einiges kennen.

#### Der Noob-Op

Es spielt überhaupt keine Rolle ob wegen neuer (unbekannter) Technologie oder
einfach nur Berufsanfänger. Beim Angehen eines Problems kommt die erstbeste
Lösung her, die gefunden wird.  Auch weil man oft garnicht in der Lage ist
beurteilen zu können ob die Lösung schön, wartbar oder auch nur im Ansatz
korrekt ist.

#### Der Realtitätsfremde

Es wird auch mal einfach ein C Programm geschrieben, gebaut und temporär
eingesetzt, nachdem im Programm der Wahl ein Flag, eine Option oder ein Feature
nicht so realisiert wie selbst erwartet.

Software wird selbst gepatched, statisch gebaut und auf Systeme verteilt.
Meister für Problemlösungen, finden aber auch Lösungen die
Betriebsführungstechnisch gesehen gar keine sind.

#### Der Engagierte Admin

Macht sich gerne länger Gedanken auch über temporäre Probleme oder niedriger
Priorität die einfach zu lösen sind. Der Anspruch an die Correctness steht
dabei an erster Stelle. Das Ergebnis will schliesslich hergezeigt werden
können.

Alles dauert recht lang, funktioniert aber. Schöne Systemkonfiguration mit
Kommentaren in den Configs bezahlt einem aber leider keiner. Funktionieren muss
es halt.

#### Der Hipster-Hacker

Nur Software einsetzen, die gerade irgendwo Trending auf HackerNews ist.
Heisser Scheiss solls sein. Enden tut das alles in einer Infrastruktur aus
pre-alpha Versionen die allesamt beim nächsten Update kaputtgehen. Ganz zu
schweigen von "production-ready" Software.

#### Der resignierende Senior

Je älter oder erfahrener der Admin dieser Kategorie desto stärker nähert er
sich dem Noob an.

Egal ob Faulheit, vergebene Bemühungen wegen der Vergänglichkeit des Systems,
baldigem Eintritt ins Rentenalter oder einfach nur, um nicht der Einzige zu
sein der sich mit dem System/ Konfiguration auskennt. Eingebaute Änderungen
sind pragmatisch, einfach, zielorientiert und mit möglichst wenig Aufwand.

### Abschliessend

Ich mag es Dinge in Schubladen zu sortieren und erwische mich oft an diesem
Gedankenmodell schrauben.  Ich weiss ehrlich gesagt nicht warum ich das
verblogge. Der Sketch dazu liegt schon ewig in meinen Drafts und über die
Feiertage kam ich mal dazu.

Um Hinweise bei vergessenen Solutions, Roles oder Practices in den Kommentaren
bitte ich sehr ;)
