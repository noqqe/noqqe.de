---
layout: post
title: "Administration von Systemen"
date: 2014-04-19T18:52:00+02:00
comments: true
categories:
- administraton
- unix
- devops
- sysops
- operations
- team
- systemadministrator
- systemadministration
- dienstleister
- it
- provider
- osbn
- ubuntuusers
- IT
---
Admins. In der Systemadministration treffen Einstellungen aufeinander.
Geschmacksrichtungen. Unterschiedliche Wissensstände. Glaubensrichtungen. Interessen. Vorlieben.
Wer tut in einem Ops-Team was und warum. Was folgt ist ein Gedankenmodell.

### Solutions

Lösungen. Eigentlich ist es das, worum es im Operations geht. Dem
nachkommen von Anforderungen an ein System. Ob neue Komponente, Sicherheitslücke oder Setup-Aufbau,
darauf lassen sich all diese Dinge herunterbrechen.

Die angewendeten Lösungen (Changes) lassen sich in maximal 2 der nachfolgenden
Kategorien einordnen. Ähnlich wie beim [CAP-Theorem](https://en.wikipedia.org/wiki/CAP_theorem).

* Einfach: Der Change ist einfach, schnell umgesetzt und auch
  für Kollegen verständlich
* Korrekt: Die Lösung kann je nach Bedarf zukünftig erweitert werden, ist sicher und auf Korrektheit überprüft
* Risikofrei: Das entstehende Riskio für die Umgebung bei Anwendung der neuen
  Lösung ist möglichst gering.

### Practices

#### "Es ging nicht anders"
Man hat keine Ahnung vom Thema. Die Lösung entspringt Google, einem Unix-Forum Post von 1997
oder im besten Fall einem Server Fault Thread mit 1,5 Upvotes.

{{< figure src="/uploads/2014/04/dirtyworkaround.gif" >}}

Diese Praktik erzeugt meistens wenig Aufwand, geht schnell. Hinterlässt aber
einen Gewissen Nachgeschmack. Fühlt sich dreckig an und sieht manchmal auch so
aus. Meist auch die Variante in der sich andere Teamkollegen beim Drüberstolpern nachher die Augen auswaschen müssen.

#### "Best Practice"
Größtmögliche Korrektheit. Dieser Ansatz entspringt meistens entweder
persönlichem Interesse des Admins an der Software oder entsprechende Schulung+Einarbeitung in
das Thema.

{{< figure src="/uploads/2014/04/correct2.gif" >}}

Was in Erweiterbarkeit, leichter Administration und Stabilität endet muss mit dem
Eingehen von Risken und Zeit bezahlt werden. Man ändert nicht einfach mal so die
Templates der Standardrolle im LDAP oder die Architektur des Datenbankclusters.
Danach hat man aber allerdings erstmal Ruhe und das angenehme Gefühl die Welt
ein Stück besser gemacht zu haben.

#### "Hart overengineert"
Wenn auch das kleinste Ein-Server-Homepage-Setup
von der Metzgerei um die Ecke bis zum Anschlag puppetisiert, hoch-skalierbar
und für alle Erweiterungen bereit eingerichtet ist.

{{< figure src="/uploads/2014/04/witchcraft.gif" >}}

Der overengineerte Ansatz ist stabil. Allerdings schwierig im Team zu
maintainen und beansprucht zu viel Zeit. Was in hohen Einmalkosten, verlorener
Mühen und Nerven endet um für Eventualitäten gewappnet zu sein die wohl niemals
eintreten werden.

### Roles
Im Lauf eines Admin-Berufslebens begegnen einem auch allerhand Charaktere.
Ich würde nicht von mir behaupten, ein Urgestein der Administration zu sein,
dennoch lernt man so einiges kennen.

#### Der Noob-Op
Es spielt überhaupt keine Rolle ob wegen neuer (unbekannter) Technologie oder einfach
nur Berufsanfänger. Beim Angehen eines Problems kommt die erstbeste Lösung her,
die gefunden wird.
Auch weil man oft garnicht in der Lage ist beurteilen zu können ob die Lösung schön,
wartbar oder auch nur im Ansatz korrekt ist.

#### Der Realtitätsfremde
Es wird auch mal einfach ein C Programm geschrieben, gebaut und temporär eingesetzt, nachdem
im Programm der Wahl ein Flag, eine Option oder ein Feature nicht so realisiert
wie selbst erwartet.

Software wird selbst gepatched, statisch gebaut und auf Systeme verteilt.
Meister für Problemlösungen, finden aber auch Lösungen die
Betriebsführungstechnisch gesehen garkeine sind.

#### Der Engagierte Admin
Macht sich gerne länger Gedanken auch über temporäre Probleme oder
niedriger Priorität die einfach zu lösen sind. Der Anspruch an die Correctness steht dabei an erster
Stelle. Das Ergebnis will schliesslich hergezeigt werden können.

Alles dauert recht lang, funktioniert aber. Schöne Systemkonfiguration mit
Kommentaren in den Configs bezahlt einem aber leider keinerr. Funktionieren muss es halt.

#### Der Hipster-Hacker
Nur Software einsetzen, die gerade irgendwo Trending auf Hackernews ist.
Heisser Scheiss solls sein. Enden tut das alles in einer Infrastruktur aus
pre-alpha Versionen die allesamt beim nächsten Update kaputtgehen. Ganz zu
schweigen von "production-ready" Software.

#### Der resignierende Senior
Je älter oder erfahrener der Admin dieser Kategorie desto stärker nähert er
sich dem Noob an.

Egal ob Faulheit, vergebene Bemühungen wegen der Vergänglichkeit des
Systems, baldigem Eintritt ins Rentenalter oder einfach nur, um nicht der Einzige zu sein der sich mit dem System/
Konfiguration auskennt. Eingebaute Änderungen sind pragmatisch, einfach,
zielorientiert und mit möglichst wenig Aufwand.

### Abschliessend

Ich mag es Dinge in Schubladen zu sortieren und erwische mich oft an diesem Gedankenmodell schrauben.
Ich weiss ehrlichgesagt nicht warum ich das verblogge. Der Sketch
dazu liegt schon ewig in meinem Draftordner und über die Feiertage kam ich mal
dazu.

Um Hinweise bei vergessenen Solutions, Roles oder Practices in den Kommentaren bitte ich sehr ;)
