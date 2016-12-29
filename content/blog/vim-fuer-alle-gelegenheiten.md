---
comments:
- author: tux.
  content: "Hab' ich eigentlich schon mal die prima Voreinstellungen von Emacs gew\xFCrdigt?
    ;-)"
  date: '2015-10-09T21:02:47.290991'
- author: Jochen
  content: "Bin seit einiger Zeit auch etwas tiefer in die vim-markdown-Sache \neingetaucht
    und stelle immer wieder erstaunt fest was vim doch f\xFCr ein \nVieh von Editor
    ist ;-)\n\nWeitere Anregungen und weiterf\xFChrende Links finden sich auf der
    [Vim for Writers](http://naperwrimo.org/wiki/index.php?title=Vim_for_Writers)
    Seite. Und falls die markdown-Files mal Tabellen enthalten sollten kommt man mit
    dem Plugin 'dhruvasagar/vim-table-mode' voll auf seine Kosten ;-)"
  date: '2015-10-10T01:44:25.414993'
- author: noqqe
  content: nein, aber mach mal!
  date: '2015-10-12T09:19:29.279169'
- author: noqqe
  content: sehr sehr geiler Tipp! Viel raus gezogen!
  date: '2015-10-16T10:45:28.716977'
- author: Jochen
  content: Gern geschehen
  date: '2015-10-17T13:47:17.250640'
- author: marvin
  content: 'mal wieder eine riesen inspiration... wie so oft in dem blog. gerade bei
    staticsitegeneretators macht es sinn vim maximal aufzubohren. was ich bis jetzt
    nicht geschafft habe: bei markdown automatisch goyo zu aktivieren.'
  date: '2015-10-22T09:04:08.808078'
- author: noqqe
  content: 'oh danke :) Freut mich!


    Ich habe meine vim config mittlerweile auf github


    https://github.com/noqqe/vim'
  date: '2015-10-22T09:38:00.688204'
date: '2015-10-09T14:29:14'
tags:
- development
- devops
- vim
- blog
- config
title: "vim f\xFCr alle Gelegenheiten"
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