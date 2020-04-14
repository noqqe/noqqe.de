---
title: "sudo via fish-shell"
date: 2018-09-29T09:31:02+02:00
tags:
- fish
- sudo
---

Neben `ctrl`+`r` ist `sudo !!` eine der Funktionalitäten die ich bei der
`fish` Shell am meisten vermisst hab. Natürlich gehts [nicht nur mir so](https://github.com/fish-shell/fish-shell/issues/288)

Aus diesem Thread hab ich dann eine Lösung für mich gefunden.

``` fish
function bind_bang
  switch (commandline -t)
  case "!"
    commandline -t $history[1]; commandline -f repaint
  case "*"
    commandline -i !
  end
end
```

Besonders cool an dieser Lösung ist, dass beim eintippen des zweiten `!` auch
für den User sichtbar der letzte Befehl eingesetzt wird.

{{< figure src="/uploads/2018/09/sudo.gif" >}}

Was nicht so schön ist, ist wie das Keybinding. In Fish gibt es eine globale
Funktion in der alle Keybindings enthalten **muss**. Das ist
`~/.config/fish/functions/fish_user_key_bindings.fish`. Dort muss ich dann folgendes einfügen

```
### bind_bang ###
bind ! bind_bang
### bind_bang ###
```

... Ich mag es eigentlich meine `functions` [modular abzulegen](https://github.com/noqqe/fishstaebchen/tree/master/conf.d)
Warum man diesen `bind` nicht einfach an einer beliebigen stelle Aufrufen kann
wird aber bestimmt einen triftigen technischen Grund haben.

Außerdem ans Herz legen kann ich noch
[sudope](https://github.com/oh-my-fish/plugin-sudope). Das kleine Plugin mappt
`ctrl`+`s` (ja etwas Strange für `bash` Migranten) und prepended jede
Commandline mit `sudo`

{{< figure src="/uploads/2018/09/sudo2.gif" >}}

Und wie machen Macher von `sudope` das mit dem `fish_user_key_bindings`?
Richtig! Da wir auch einfach fröhlich am globalen File
`~/.config/fish/functions/fish_user_key_bindings.fish` herumgepatched.

``` fish
function fish_user_key_bindings
    ### sudope ###
    set -q sudope_sequence
    or set -l sudope_sequence \cs
    if not bind | string match -rq '[[:space:]]sudope$'
        if bind $sudope_sequence >/dev/null 2>/dev/null
            echo "sudope: The requested sequence is already in use:" (bind $sudope_sequence | cut -d' ' -f2-)
        else
            bind $sudope_sequence sudope
        end
    end
    ### sudope ###
end
```
