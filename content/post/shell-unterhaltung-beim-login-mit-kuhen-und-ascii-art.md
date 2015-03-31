---
date: 2009-03-09T21:00:27+02:00
type: post
slug: shell-unterhaltung-beim-login-mit-kuhen-und-ascii-art
status: publish
comments: true
title: Shell | Unterhaltung beim Login mit Kühen und Ascii-ART
alias: /archives/545
categories:
- Coding
- Linux
tags:
- bash
- cow
- cowsay
- fortune
- fun
- login
- shell
- unterhaltung
- weisheit
---

![](http://zwetschge.org/pic/fortunecow.JPG)

Hübsch oder? Mithilfe von Fortune (das ich bei CodeCocktail ja schon einmal vorgestellt habe) und Cowsay wird bei jedem Login in eine Shell ein zufälliges Tier/Comicfigur in ASCII-ART abgebildet,  das mit Ratschlägen/Witzen/Weisheiten/und anderen lustigen Sachen zur Seite steht. Hier ein kleines HOW-TO:

```

apt-get install fortune fortune-de cowsay
```


vim .bashrc (des gewünschten Users bzw Root):

```

random_cow(){
cows=(`ls -1 /usr/share/cowsay/cows/*.cow`)
cow_num=${#cows[*]}
echo ${cows[RANDOM%$cow_num]}
}
if [[ $- == *i* ]]; then
fortune de | cowsay -f `random_cow` -W 75
fi
```


Weit erstmal zur Grundfunktion =) Das Ganze ist aber noch sehr weit personalisierbar sag ich mal =) Mit fortune -f  lässt sich eine Liste anzeigen mit Kategorien der Sprüchen. Diese wiederrum lassen sich beispielsweise statt "fortune de" in den .bashrc Eintrag einfügen (fortune stilblueten).

```

zwetschge:~# fortune -f
100,00% de
1,64% woerterbuch
0,63% sprichworte
0,23% elefanten
3,30% namen
0,12% bahnhof
0,23% anekdoten
1,65% ms
2,11% regeln
0,37% murphy
0,25% sicherheitshinweise
1,10% warmduscher
0,74% stilblueten
1,80% sprueche
1,43% letzteworte
7,70% unfug
0,14% loewe
0,61% mathematiker
0,69% lieberals
0,23% translations
65,30% zitate
6,06% witze
0,13% vornamen
0,19% huhn
1,06% computer
```


Genauso wie sich die ASCII-Comics noch anpassen lassen. In /usr/share/cowsay/cows/ liegen jede Menge Cow-Dateien aus denen man löschen/hinzufügen neue herunterladen oder selber basteln kann =) Meine Favorites sind natürlich diese beiden Star-Wars-Genossen hier:
![](http://zwetschge.org/pic/koalacow.JPG)

Greez, Flo
