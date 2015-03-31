---
layout: post
title: "Gitolite migrieren"
date: 2012-11-30T20:25:00+02:00
comments: true
categories:
- osbn
- git
- gitolite
- Debian
- OpenBSD
---

Eine der "Challenges" beim Umziehen auf die OpenBSD Kiste war auch, meine
gitolite Installation samt aller `git` Repositories zu migrieren.

{% img center /uploads/2012/11/ghostbusters.jpg %}

Das Einrichten auf dem neuen Server ($newServer) kann eigentlich genau so ablaufen,
wie in der [Anleitung](https://github.com/sitaramc/gitolite/blob/master/README.txt)
beschrieben.

Als nächstes habe ich in .git/config des alten `gitolite-admin` Repositories
die Konfiguration des alten Servers ($oldServer) in $newServer geändert

{% codeblock %}
[remote "origin"]
    url = git@$newServer:gitolite-admin
{% endcodeblock %}

Wenn man diese dann pushed wird der `push` rejected, weil kein auto-merging
stattfinden kann.

{% codeblock %}
$ gitolite-admin (master)$ git push origin master
To git@$newServer:gitolite-admin
 ! [rejected]        master -> master (non-fast-forward)
error: failed to push some refs to 'git@$newServer:gitolite-admin'
To prevent you from losing history, non-fast-forward updates were rejected
Merge the remote changes (e.g. 'git pull') before pushing again.  See the
'Note about fast-forwards' section of 'git push --help' for details.
{% endcodeblock %}

Also erst wie schon im Hint von git steht `pull`en

{% codeblock %}
$ gitolite-admin (master)$ git pull
warning: no common commits
remote: Counting objects: 6, done.
remote: Compressing objects: 100% (4/4), done.
remote: Total 6 (delta 0), reused 0 (delta 0)
Unpacking objects: 100% (6/6), done.
From $newServer:gitolite-admin
 + 66155ef...17e8ee7 master     -> origin/master  (forced update)
Auto-merging conf/gitolite.conf
CONFLICT (add/add): Merge conflict in conf/gitolite.conf
Automatic merge failed; fix conflicts and then commit the result.
{% endcodeblock %}

Der ausgezeichnete Conflict sollte dann ungefähr so aussehen:

{% codeblock %}
gitolite-admin (master*)$ vim conf/gitolite.conf
<<<<<<< HEAD
<YOUR CONFIGURATION HERE>
=======
repo gitolite-admin
    RW+     =  YourKey.pub

repo testing
    RW+     =   @all
>>>>>>> 17e8ee7c81ee76c5ec1882a4d65d8dd7a3d468d0
{% endcodeblock %}

selbst wenn es nicht so aussieht, sollte es nicht allzukompliziert sein,
den Merge durchzuziehen. Kernel Hacker schaffen sowas ja auch irgendwie :)

{% codeblock %}
$ gitolite-admin (master*)$ git commit -a -m "Merged old config into blank new gitolite-admin repo"
[master e8ac423] Merged old config into blank new gitolite-admin repo
{% endcodeblock %}

Wenn das geschafft ist wirds interessanter. Beim `push` des gemergten
Commits erstellt das gitolite auf dem neuen Server alle Repositories,
die es bis dato noch nicht kannte.

{% codeblock %}
gitolite-admin (master)$ git push origin master
Counting objects: 319, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (170/170), done.
Writing objects: 100% (315/315), 29.96 KiB, done.
Total 315 (delta 82), reused 306 (delta 79)
remote: Initialized empty Git repository in /home/git/repositories/bash-statistic-utils.git/
remote: Initialized empty Git repository in /home/git/repositories/vim.git/
[...]
remote: Initialized empty Git repository in /home/git/repositories/wikidata.git/
To git@$newServer:gitolite-admin
   17e8ee7..e8ac423  master -> master
{% endcodeblock %}

Auf dem neuen Server kann man dann einfach via rsync die Repository Inhalte
nachziehen.

{% codeblock %}
rsync -av root@$oldServer:/home/git/repositories/* repositories/
{% endcodeblock %}


## Testen

Um zu testen ob die Migration erfolgreich wahr, hab ich ein Repo in `tmp/`
ausgecheckt und es mit dem letzten Stand auf dem alten Server verglichen.

{% codeblock %}
$ cd /tmp
$ git clone git@$newServer:repo repo
$ cd repo
$ git log -1
{% endcodeblock %}

## Repos umarbeiten

Da meine Repos alle im selben Verzeichnis liegen, fiel es mir sehr leicht alle .git/config Files zu finden
und diese umzuschreiben

Evaluieren:

{% codeblock %}
for x in $(locate --regex '/home/noqqe/Code/.*/.git/config$'); do
  fgrep '$oldServer' $x
done
{% endcodeblock %}

Renaming:

{% codeblock %}
for x in $(locate --regex '/home/noqqe/Code/.*/.git/config$') ; do
  sed -i -e 's#$oldServer#$newServer#' $x
done
{% endcodeblock %}

Merkt man eigentlich an der Frequenz der Blogposts, dass ich Urlaub habe? :)
