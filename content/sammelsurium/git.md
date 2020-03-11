---
title: Git
date: 2020-03-05T12:00:00
tags:
- Software
- git
---

## Gelöschte Dateien

Files mit bestimmten Suffix aus Git History finden

    git log --all --summary -- '*.Rdata'

<!--more-->

Gelöschte Files aus History anzeigen

    git log --diff-filter=D --summary

File aus History löschen

    git filter-branch -f --tree-filter "rm -f foo/bar/baz.RData" HEAD

## Squash

Manchmal dauert es 1-18 Commits um ein Feature fertigzustellen. Da man das
aber nicht in `master` haben will, muss man ein bisschen was dazu tun.

```
[ on current messed up feature branch "TICK-1337-MESSY" with too many commits ]
git checkout master
git checkout -b TICK-1337-Feature
git merge --squash TICK-1337-MESSY
git add .
git commit -a -m "Add feature XYZ for TICK-1337"
git push
[ pull request ]
```

So bekommt man quasi alle Changes die seit `master` im Feature Branch
TICK-1337-MESSY passiert sind in den neuen Branch TICK-1337-Feature überführt
und kann dann eine schöne Commit Message formulieren.


## Ein einzelnes File aus einem Repo holen

Man kann ein einzelnes (bzw. auch mehrere Files) aus einem Repo holen, wenn
man `git archive` dafür verwendet

```
git archive -o out.tar --remote=ssh://github.com/noqqe/foo.git HEAD:master <file>
```

Super für Automatisierung oder ähnliches.

## git proxy

Manchmal kommt man in die Verl

    [http]
        sslVerify = false
        proxy = http://user:pass@proxy.example.com:3128
    [url "https://"]
        insteadOf = git://
        insteadOf = git+https://
        insteadOf = git+http://
        insteadOf = http://

## Search

search for commit content (i.e., actual lines of source, as opposed to
commit messages and the like), what you need to do is:

    git grep <regexp> $(git rev-list --all)

This will grep through all your commit text for regexp.
Here are some other useful ways of searching your source:
Search working tree for text matching regular expression regexp:

     git grep <regexp>

Search working tree for lines of text matching regular expression regexp1
or regexp2:

    git grep -e <regexp1> [--or] -e <regexp2>

Search working tree for lines of text matching regular expression regexp1
and regexp2, reporting file paths only:

    git grep -e <regexp1> --and -e <regexp2>

Search working tree for files that have lines of text matching regular
expression regexp1 and lines of text matching regular expression regexp2:

    git grep -l --all-match -e <regexp1> -e <regexp2>

## Submodules

Git Submodules sind ultimativ kaputt, meine Meinung. Ich versuche Software wo
es geht zu meiden die solche einsetzt.

    $ git submodule sync
    No submodule mapping found in .gitmodules for path 'bundle/Screen-vim---gnu-screentmux'
    No submodule mapping found in .gitmodules for path 'bundle/Vim-R-plugin'
    No submodule mapping found in .gitmodules for path 'bundle/screen.vim'
    No submodule mapping found in .gitmodules for path 'bundle/vim-powerline'
    No submodule mapping found in .gitmodules for path 'bundle/vundle'

Infos suchen:

    git config -l

Submodule löschen:

* Delete the relevant line from the .gitmodules file.
* Delete the relevant section from .git/config.
* Run git rm --cached path_to_submodule (no trailing slash).

```
$ git rm --cached bundle/R.vim
rm 'bundle/R.vim'
```

* Commit the project.
* Delete the now untracked submodule files.

## Inspect all versions for 1 specific line in one file

Das ist zum Beispiel einfach sau geil, wenn man im Repo vergessen hat
sauber die Versionen zu taggen und sucht, wann die version angehoben wurde.

    $ git log --pretty=short -u -L 16,16:setup.py

    commit 2ec513539add777f1dcc7a5e803538f5180fa2eb
    Author: Florian Baumann <wa1@noqqe.de>

        version bump

    diff --git a/setup.py b/setup.py
    --- a/setup.py
    +++ b/setup.py
    @@ -16,1 +16,1 @@
    -    version='18.0.0',
    +    version='19.0.0',

    commit 30c5f4ea436ca7b1c63c8ba647b971af024ee1e6
    Author: Florian Baumann <wa1@noqqe.de>

        version bump

    diff --git a/setup.py b/setup.py
    --- a/setup.py
    +++ b/setup.py
    @@ -16,1 +16,1 @@
    -    version='17.0.0',
    +    version='18.0.0',

## git gc

Wenn ich ein git repository wieder aufräumen möchte.

[http://git-scm.com/docs/git-gc](http://git-scm.com/docs/git-gc)
