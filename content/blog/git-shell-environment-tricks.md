---
comments:
- author: noqqe
  content: "<p>Das ist ja der absolute Wahnsinn wie schrecklich der Post im ubuntuusers
    Planeten aussieht.\_ </p>"
  date: '2012-03-10T13:50:00'
- author: Jakob
  content: "<p>git diff --staged ist dasselbe wie git diff --cached, also ist der
    Alias \u201Egsdiff\u201C \xFCberfl\xFCssig (Quelle: man 1 git-diff).</p>"
  date: '2012-03-10T21:43:28'
- author: Jakob
  content: "<p>Was ich \xFCbrigens auch empfehlen kann: vcprompt (<a href=\"https://bitbucket.org/gward/vcprompt)\"
    rel=\"nofollow\">https://bitbucket.org/gward/vcprompt</a>. Das kann man sehr sch\xF6n
    in seinen Prompt einbauen. Woher wei\xDF eigentlich dein Blog, dass ich das dunkle
    Solarized-Theme in meinem Terminal verwenden? :O Und ich habe \xFCbrigens auch
    ein `git diff`-Alias:</p><p>&gt; gdf ist ein Alias von `git diff HEAD | pygmentize
    -l diff | less -R'.</p>"
  date: '2012-03-10T21:53:49'
- author: noqqe
  content: "<p>\_Auch nett sich pygmentize f\xFCr sowas zu holen :)<br></p>"
  date: '2012-03-11T11:37:38'
- author: sebix
  content: <p>Nette Zusammenstellung, danke!</p>
  date: '2012-03-12T20:31:06'
- author: BO
  content: "<p>Was machst du eigentlich mit build-Hinterlassenschaften bei\_parse_git_dirty()?
    Auf ein build-Verzeichnis kann man sich ja nicht immer verlassen, das man evtl.
    als ignore eintragen k\xF6nnte.\_Ich benutze erst seit ein paar Tagen ernsthaft
    git(-svn) - keine Ahnung, wie da das \xFCbliche Vorgehen w\xE4re, wenn es eins
    gibt.</p>"
  date: '2012-03-14T02:12:22'
- author: noqqe
  content: "<p>Build hinterlassenschaften? Also Sachen die nicht Committet werden
    (sollen?) </p><p>Naja das ist ganz einfach abh\xE4ngig davon wie du deine git
    Repos pflegst. Eigentlich sollte dein Repo, wenn du nicht drin arbeitest, immer
    sauber sein. Gibts auch ne coole Sammlung von ignore Sets auf github die du vielleicht
    nutzen willst. </p>"
  date: '2012-03-14T06:55:39'
- author: noqqe
  content: "<p>Hast du nat\xFCrlich recht. Aber ehrlichgesagt nutze ich nur gsdiff
    - weswegen mir der andere nicht wirklich aufgefallen ist :/ </p>"
  date: '2012-03-14T06:56:39'
- author: BO
  content: "<p>Klar, wenn ich nichts darin mache, dann ist da auch nichts, das mich
    st\xF6ren kann, aber es w\xFCrde mich dann ja auch nicht st\xF6ren.\_Mit build-Hinterlassenschaften
    meine ich all den Kram, der produziert wird, wenn du im Projekt mal `make` machst.<br>Aber
    ich merk schon, die Frage an sich war schon bl\xF6d, genau daf\xFCr ist ja\_parse_git_dirty()
    da. Gehirnfurz, sorry. :)</p>"
  date: '2012-03-15T13:38:25'
- author: BO
  content: "<p>Konkret: `git status` gibt mir einen Haufen untracked files zur\xFCck,
    vieles kann man nat\xFCrlich ausklammern, *.so *.o und so, aber es gibt auch einfache
    C++ files darunter, die beim bauen zusammengebastelt werden. Z.B.\_<a href=\"http://queryyacc.cc\"
    rel=\"nofollow\">queryyacc.cc</a>, wie erreiche ich, dass mir die beim `git status`
    nicht mehr angezeigt wird? Ist da .gitignore das vern\xFCnftige Mittel zur Wahl?<br></p>"
  date: '2012-03-15T13:47:38'
date: '2012-03-10T10:00:00'
tags:
- tipps
- functions
- shell
- prompt
- tricks
- bashrc
- git
- linux
- statistics
- bash
- stats
title: Git Shell Environment Tricks
---

Heute mal ein Blick in meine `~/.bashrc` über viele Schnippsel und Funktionen
die ich im Laufe der Jahre für git zusammengesammelt habe. Wenns mir noch
einfällt mit Quelle :)

## Aliases

Teilweise selber gebastelt, teilweise aus [bash-it](https://github.com/revans/bash-it) geklaut.

``` bash
## git Aliases
alias gcl='git clone'
alias ga='git add'
alias gall='git add .'
alias g='git'
alias gs='git status'
alias gss='git status -s'
alias gl='git log --oneline'
alias gup='git fetch && git rebase'
alias gp='git push'
alias gpo='git push origin '
alias gb='git branch'
alias gcount='git shortlog -sn'
alias gdel='git branch -D'
alias gcm='git commit -a -m'
alias gll='git log --graph --pretty=oneline --abbrev-commit'
alias cdiff='git diff --cached'
alias gsdiff='git diff --staged'
```

Hervorheben will ich hier besonders `gll`:

{{< figure src="/uploads/2012/03/gitlog-graph.png" >}}

## git Informationen

Wer viel in verschiedenen git Repos unterwegs ist hilft vielleicht (wie mir)
diese Funktion. Selbst geschrieben.

``` bash
function git_info() {

    if [ -n "$(git symbolic-ref HEAD 2> /dev/null)" ]; then
        # print informations
        echo "git repo overview"
        echo "-----------------"
        echo

        # print all remotes and thier details
        for remote in $(git remote show); do
            echo $remote:
            git remote show $remote
            echo
        done

        # print status of working repo
        echo "status:"
        if [ -n "$(git status -s 2> /dev/null)" ]; then
            git status -s
        else
            echo "working directory is clean"
        fi

        # print at least 5 last log entries
        echo
        echo "log:"
        git log -5 --oneline
        echo

    else
        echo "you're currently not in a git repository"

    fi
}
```

``` bash
$ git_info
git repo overview
-----------------

origin:
* remote origin
  Fetch URL: git@github.com:revans/bash-it.git
  Push  URL: git@github.com:revans/bash-it.git
  HEAD branch: master
  Remote branch:
    master tracked
  Local branch configured for 'git pull':
    master merges with remote master
  Local ref configured for 'git push':
    master pushes to master (up to date)

status:
working directory is clean

log:
39f8ef9 add defaults autocompletion for OS X
67f642f Merge pull request #102 from jpschewe/master
d0ffb0d Merge remote-tracking branch 'berenm/master'
94a7b78 Revert "Revert new color framework"
87d7d7a Fixed issue #103 caused by "normal" colors not reseting bold/underline/... text attributes.
```

## git Bash Prompt

Eines der nützlichsten Dinge. Wenn das aktuellen Working Directory ein git Repository ist verändert sich der Prompt.
Es zeigt falls das der Fall ist den aktuell ausgecheckten Branch an und eine Asterisk (*) wenn das
PWD sich in einem uncommitteten Zustand befindet.

``` bash
GIT_THEME_PROMPT_DIRTY='*'

function git_prompt_info() {
  ref=$(git symbolic-ref HEAD 2> /dev/null) || return
  echo -e " (${ref#refs/heads/}$(parse_git_dirty))"
}

function parse_git_dirty {
  if [[ -n $(git status -s 2> /dev/null |grep -v ^# | grep -v "working directory clean" ) ]]; then
      echo -e "$GIT_THEME_PROMPT_DIRTY"
  else
    echo -e "$GIT_THEME_PROMPT_CLEAN"
  fi
}

PS1="\u@\h:\w\[\$(git_prompt_info)\]$ "
```

``` bash
noqqe@deathstar:~/Code$ cd octopress
noqqe@deathstar:~/Code/octopress (master)$ touch foobar
noqqe@deathstar:~/Code/octopress (master*)$
```

## git Remotes

Zwei Funktionen an die ich mich ziemlich gewöhnt habe, weil ich für meine git
repos meistens sowieso den selben Remote benutze. Ebenfalls veruntreut vom
bash-it Framework.

``` bash
function git_remote {
  echo "Running: git remote add origin git@n0q.org:$1"
  git remote add origin git@n0q.org:$1
}

function git_first_push {
  echo "Running: git push origin master:refs/heads/master"
  git push origin master:refs/heads/master
}
```

## git Statistiken

Hier noch ein Stück für die Statistik Liebhaber. Ich mags.

``` bash
function git_stats {
if [ -n "$(git symbolic-ref HEAD 2> /dev/null)" ]; then
    echo "Number of commits per author:"
    git --no-pager shortlog -sn --all
    AUTHORS=$( git shortlog -sn --all | cut -f2 | cut -f1 -d' ')
    LOGOPTS=""
    if [ "$1" == '-w' ]; then
        LOGOPTS="$LOGOPTS -w"
        shift
    fi
    if [ "$1" == '-M' ]; then
        LOGOPTS="$LOGOPTS -M"
        shift
    fi
    if [ "$1" == '-C' ]; then
        LOGOPTS="$LOGOPTS -C --find-copies-harder"
        shift
    fi
    for a in $AUTHORS
    do
        echo '-------------------'
        echo "Statistics for: $a"
        echo -n "Number of files changed: "
        git log $LOGOPTS --all --numstat --format="%n" --author=$a | cut -f3 | sort -iu | wc -l
        echo -n "Number of lines added: "
        git log $LOGOPTS --all --numstat --format="%n" --author=$a | cut -f1 | awk '{s+=$1} END {print s}'
        echo -n "Number of lines deleted: "
        git log $LOGOPTS --all --numstat --format="%n" --author=$a | cut -f2 | awk '{s+=$1} END {print s}'
        echo -n "Number of merges: "
        git log $LOGOPTS --all --merges --author=$a | grep -c '^commit'
    done
else
    echo "you're currently not in a git repository"
fi
}
```

``` bash git_stats Output
Number of commits per author:
   149  Mark Szymanski
    64  Robert R Evans
    53  Travis Swicegood
    22  Florian Baumann
    16  Jesus de Mula Cano
    14  John Schulz
    12  Ryan
    10  JFSIII
    10  Ryan Kanno
     9  David DeSandro
-------------------
Statistics for: Mark
Number of files changed: 52
Number of lines added: 1577
Number of lines deleted: 733
Number of merges: 19
-------------------
Statistics for: Robert
Number of files changed: 74
Number of lines added: 5817
Number of lines deleted: 3065
Number of merges: 16
-------------------
Statistics for: Travis
Number of files changed: 106
Number of lines added: 4416
Number of lines deleted: 3919
Number of merges: 19
-------------------
Statistics for: Florian
Number of files changed: 14
Number of lines added: 363
Number of lines deleted: 145
Number of merges: 3
-------------------
[...]
```

## git fehlende Files entfernen

``` bash
function git_remove_missing_files() {
  git ls-files -d -z | xargs -0 git update-index --remove
}
```

## git lokales ignore

``` bash
function local-ignore() {
  echo "$1" >> .git/info/exclude
}
```
