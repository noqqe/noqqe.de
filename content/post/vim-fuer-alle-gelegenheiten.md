---
categories:
- blog
- development
- devops
- osbn
- planetenblogger
- ubuntuusers
date: 2015-10-09T16:29:14+02:00
tags:
- vim
- config
- blog
title: vim für alle Gelegenheiten
---

Ich benutze vim für so gut wie alles. Blogposts, Mails, Notizen, Dokumentation, Skripte und Code.
Doch egal wie sehr ich diesen wunderbaren Editor auch konfiguriere, es kommen immer wieder
Situationen auf, in denen das ein oder andere Setting keinen Sinn ergibt.

{{< figure src="/uploads/2015/10/vim.jpg" >}}

Spelling, Textwidth und Wrapping sind zum Beispiel solche Settings.
Über einen [Blogpost](http://www.swamphogg.com/2015/vim-setup/) über Vim mit Markdown stiess ich dann auf ein
[Plugin](https://github.com/reedes/vim-pencil), welches sich
um Soft-/Hard-Wrapping kümmert.

Im README befand sich ein Beispiel, in dem für verschiedene FileTypes unterschiedliche Optionen
gesetzt werden.

``` vim
augroup pencil
  autocmd!
  autocmd Filetype *      call pencil#init({'wrap': 'soft', 'textwidth': 75})
                            \ | setl textwidth=0 wrapmargin=0 wrap
  autocmd Filetype mail   call pencil#init({'wrap': 'soft', 'textwidth': 75})
                            \ | setl sw=2 ts=2 noai nonu nornu
                            \ | setl tw=100
                            \ | setl fdo+=search
  autocmd FileType markdown call pencil#init({'wrap': 'soft', 'textwidth': 80})
                            \ | setl spell spelllang=de,en fdl=4 noru nonu nornu
                            \ | setl tw=100
                            \ | setl fdo+=search
  autocmd FileType text   call pencil#init({'wrap': 'soft', 'textwidth': 75})
                            \ | setl spell spelllang=de,en fdl=4 noru nonu nornu
                            \ | setl tw=75
                            \ | setl fdo+=search
augroup END
```

Für die meisten vim User dürfte das nichts neues sein. Gerade Textwidth und Wrapping
haben bei mir auch erstmal Versteh-Bedarf gehabt. Ich kann jetzt auch die Tagbar (für Variablen und
Funktionen) nur für bestimmte Filetypes öffnen. Außerdem hab ich noch ein paar weitere Plugins gefunden.

``` vim
Plugin 'tpope/vim-fugitive' " Git Wrapper
Plugin 'tpope/vim-commentary' " auto commenting with keybinding gc
Plugin 'airblade/vim-gitgutter' " git diff line next to line numbers
Plugin 'junegunn/goyo.vim'  " writer fullscreen mode
Plugin 'reedes/vim-pencil'  " Soft-, Hard-Wrapping
Plugin 'ntpeters/vim-better-whitespace' " highlightning for whitespace
```

So sieht mein vim wenn ich ein .markdown File öffne, komplett anders aus als bei Puppet oder Python.
Kontextabhängiger Editor. So. gut. Ich werde da noch weiter rumbasteln, denke ich mal sehr.
