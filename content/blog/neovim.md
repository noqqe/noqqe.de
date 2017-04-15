---
date: 2017-04-15T08:53:45+02:00
tags:
- vim
- neovim
title: neovim
---

Seit ca. 2 Monaten oder so benutze ich nun auf dem MacBook
[neovim](https://neovim.io). Und zwar einfach als Drop-In-Replacement.

``` bash
NVIM=$(which nvim 2>/dev/null)

if [[ -x $NVIM ]]; then
  alias vim=$NVIM
  export EDITOR=$NVIM
fi
```

Solchen quatsch muss ich eigentlich nur machen, weil ich [meine bashrc](https://github.com/noqqe/nano-bash/)
auf relativ vielen Systemen einsetze, aber das ist eine andere Geschichte.

Relativ angenervt hat mich der per default aktivierte Mouse Support. Ich
weiss nicht wer das nutzt aber eigentlich hindert es mir nur daran mal
etwas mit via Mouse Selektion aus dem `vim` herauszukopieren. Die `:help`
zu `mouse` fand ich auch etwas verwirrend da es nicht klar sagt wie man
einfach alle Varianten disabled, aber okay.

``` vim
if has('nvim')
  set mouse=r
  set noincsearch
endif
```

Wie man sieht ist noch eine zweite Option im Codeblock. `incsearch` ist ein
ebenfalls standardmäßig aktiviertes Features in `nvim` und bewirkt, dass
man bei der Suche via `/` so eine Art fuzzy search bekommt bei der einem
der Cursor wild über den ganzen Screen hüpft noch bevor man Enter drückt.
Das wäre nicht schlimm, würde ich nicht genau dieses letzte Enter immer
vergessen, da mein Cursor sich ja bereits an der richtigen stelle befindet
oder ich endlos versuche `n` zu tippen um zum nächsten Ergebnis zu kommen.

Ich denke mal da bin ich einfach Gewohnheitstier. Gefällt mir jedenfalls
nicht, deshalb mach ichs aus. Ich hab auch noch jede Menge anderen Unsinn
in meiner `.vimrc`. Der Trend geht allerdings zu
[KISS](https://en.wikipedia.org/wiki/KISS_principle). Ist ebenfalls auf
[Github](https://github.com/noqqe/vim)

Ansonsten muss ich sagen ist `neovim` eigentlich schick. Man merkt es nicht
und darauf kommt es doch an, oder?
