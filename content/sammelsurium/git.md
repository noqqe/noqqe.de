---
title: Git
date: 2012-06-18T02:00:00.000000
tags: 
- Software
- git
---


Files mit bestimmten Suffix aus Git History finden

    git log --all --summary -- '*.Rdata'

Gelöschte Files aus History anzeigen

    git log --diff-filter=D --summary

File aus History löschen

    git filter-branch -f --tree-filter "rm -f foo/bar/baz.RData" HEAD

## git gc

Wenn ich ein git repository wieder aufräumen möchte.

    http://git-scm.com/docs/git-gc

## git proxy

    [http]
        sslVerify = false
        proxy = http://user:pass@proxy.example.com:3128
    [url "https://"]
        insteadOf = git://
        insteadOf = git+https://
        insteadOf = git+http://
        insteadOf = http://

## Performance Test

    git add test; time while [ $C -le 1004225 ]; do   date "+%Y%m%d %H%M%s%N" >> test ;   git commit -a -m "$C";   ((C++)); done

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

## Fucking Submodules

    $ git submodule sync
    No submodule mapping found in .gitmodules for path 'bundle/Screen-vim---gnu-screentmux'
    No submodule mapping found in .gitmodules for path 'bundle/Vim-R-plugin'
    No submodule mapping found in .gitmodules for path 'bundle/screen.vim'
    No submodule mapping found in .gitmodules for path 'bundle/vim-powerline'
    No submodule mapping found in .gitmodules for path 'bundle/vundle'

Infos suchen:

    git config -l

To remove a submodule you need to:

* Delete the relevant line from the .gitmodules file.
* Delete the relevant section from .git/config.
* Run git rm --cached path_to_submodule (no trailing slash).


    $ git rm --cached bundle/R.vim
    rm 'bundle/R.vim'

* Commit the superproject.
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
