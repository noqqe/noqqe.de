---
title: "CSS"
date: 2020-03-12T10:21:05+01:00
tags:
- Software
---

Mit CSS tat ich mich schon immer sehr schwer. Ich verstehe das System oft
nicht und ich denke das geht auch Hand in Hand mit dem Fakt, das ich von
Design auch nichts anfangen kann.

Hier nun mein Cheatsheet

<!--more-->

## Class vs. ID

Ich denke ich werde es mir nie merken können. Deshalb schreib ich es hier
auf.

Eine Klasse

``` html
<div class="menu">
.menu {
  color: black;
}
```

Eine ID

``` html
<div id="menu">
#menu {
  color: black;
}
```

## Verschachtelte HTML Tags

> Select and style every <p> element that is inside <div> elements

``` css
div p {
  background-color: yellow;
}
```

## Attribute überschreiben

Manchmal schreiben 2 Definitionen in das gleiche Element

``` css
 a.color-projects {
     color: #f1fa8c !important;
 }
```

## Mobile Ansicht

Um responsive zu sei kann man mit `media` views auf die Screenauflösung gehen
und dann Elemente anders anzeigen bzw verstecken.

``` css
@media only screen and (max-width: 640px) {
  div#contentwrapper {
    max-width: 100%;
  }
}
```
