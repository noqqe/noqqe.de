---
title: "Verschiedene vim Versionen & etwas micro"
date: 2017-10-03T09:15:23+02:00
draft: false
tags:
- vim
- micro
- editor
---

Wow. Gerade realisiert das ich hier ganze 5(!) Monate nicht mehr gebloggt
habe. Quasi den ganzen Sommer über. Es fühlt sich auch ein bisschen fremd
an wieder mit [hugo](https://gohuho.io) zu arbeiten, was wahrscheinlich
hauptsächlich daran liegt das die Befehle dazu aus meiner `bash` History
rausrotiert sind und ich sie nicht wie sonst mit `ctrl+r` suchen konnte.

Wie auch immer. In letzter Zeit hab ich wieder etwas an meinem `vim` herumgebastelt.
Meistens entstand dieses Gebastel (mit dem ich auch noch nicht wirklich fertig bin)
wegen verschiedenen `vim` Versionen auf Servern und Clients. Genau gesagt `vim 7.4`,
`vim 8` und `neovim` den ich auf meinen Clients [benutze](https://noqqe.de/blog/2017/04/15/neovim/).

Mein Lieblingskandidat ist diese Zeile.

```
set clipboard+=unnamedplus
```

Sie führt dazu, dass `dd`/`yy` Befehle ihren Inhalt in ein persistentes
Register schreiben und man eben diesen Inhalt wieder via `p` einfuegen
kann. Der wichtige Teil: Auch zwischen 2 verschiedenen Files hin und her.
Wie man die Maus deaktiviert hatte ich zwar schon im letzten Post geklärt,
aber um meinen Hass auszudrücken hier nochmal. Warum zur Hölle vim 8?!

```
set mouse=r
```

Aber auch nur für mich nützliche Sachen wie einen Zeitstempel einfügen,
haben den Weg in meine private Vim Config gefunden.

```
nmap <silent> <leader>d :print=strftime('%F %H:%M')<CR>A
```

oder automatische Fehlerkorrektur via `spellcheck`, die ich nun schon
länger für zum Beispiel [jrnl](http://jrnl.sh) nutze.

```
nmap <silent> <leader>f z=1<CR><CR>
```

Wer mir auf [Twitter](https://twitter.com/noqqe) folgt hat das vielleicht
schon gesehen. Via `+/foobar` kann man beim Befehlsaufruf den Suchstring
gleich mitgeben. Mit der nachfolgenden `bash` Funktion kann ich rekursiv
nach einem String in files "greppen" und alle Files mit Treffern öffnen und
den gesuchten String gleichzeitig in `vim` highlighten. Nice!

```
function agvim () {
  local files=$(ag -l "$@")
  vim +/"$@" $files
}
```

Viele dieser Sachen sind zwar nun vielleicht auf meinem lokalen Rechner
schoen konfiguriert, aber ich kann in der Arbeit niemandem zumuten eben
diese Config auch zu bedienen. Wir haben uns deshalb fuer den `root`
Account auf allen Servern auf eine minimale Konfiguration geeinigt die wir
über Puppet ausrollen können. Das Modul
[zanloy-vim](https://forge.puppet.com/zanloy/vim/readme) kann ich daher
sehr empfehlen. Diese Config funktioniert für die meisten, ohne viel
Tamtam, egal ob Schwarzes oder Weisses Terminal. Wichtig ist doch, dass man
sich auf dem `vim` im `root` Account verlassen kann. Wenn nichts mehr geht,
sollte der noch tun.

``` puppet
class { 'vim':
  ensure           => present,
  set_as_default   => true,
  opt_bg_shading   => 'light',
  opt_indent       => true, # enables 2 space tabs
  opt_lastposition => false,
  opt_powersave    => false,
  opt_syntax       => true,
}
```

In other news: Ich hab mir mal ein bisschen die Zeit genommen den Editor
[Micro](https://micro-editor.github.io) anzusehen und zu benutzen. Es
wirkt alles sehr nett und vor allem einheitlich. Die Konfiguration ist
übersichtlich, nicht so wie bei `vim`, wo man jede Option erst bei `:help`
nachschlagen muss. Tatsächlich schrieb ich diesen Post auch mit `micro`,
einfach um mal wieder etwas neues zu versuchen.

Was aber witzig ist, `ctrl-z` ist der Shortcut für undo. Ich wusste garnicht
das man das abfangen kann.
