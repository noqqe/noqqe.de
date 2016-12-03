---
aliases:
- /archives/881
- /blog/2010/02/11/wordpress-augenzucker-fur-code-tag
date: '2010-02-11T20:09:42'
tags:
- development
- style
- code
- sytnax
- theme
- wordpress
- block
- wp
- php
- css
- codeblock
title: "Wordpress | Augenzucker f\xFCr Code-Tag."
---

EyeCandy mein ich. Auf
[CodeX](http://codex.wordpress.org/Writing_Code_in_Your_Posts) von
Wordpress gibts Dokumentationen des Sources. Wenn man
<del>stundenlang</del> kurz sucht findet man auch nach dem was man will.
Jedenfalls gefiel mir die standardmäßige Ausführung von dem code-Tag in
Wordpress nicht. Mit dem Artikel auf CodeX und etwas Spielerei hab ich
meine style.css jetzt so ummodifiziert, das das ganze wie folgt aussieht:

``` css
/* Code pre */
code { max-width: 500px;
color: white;
margin: 1px;
display: block;
overflow: auto;
padding: 1px;
background: #000000 }
```

Was übrigens gleichzeitig den Zusatz in meiner style.css darstellt.