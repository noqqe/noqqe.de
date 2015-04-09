---
date: 2010-02-10T19:37:53+02:00
type: post
comments: true
title: Git | Repositories auf Server anlegen
aliases:
- /blog/2010/02/10/git-repositories-auf-server-anlegen
- /archives/867
categories:
- Development
- Linux
tags:
- git
- gitosis
- howto
- Linux
- progit
- repo
- repositories
- scm
- svn
- unix
---

![git-logo](/uploads/2010/02/git-logo.png)<del>I</del><del>c</del><del>h
bin ja zukunftsorientiert.</del> Mir wurde einbläut zukunftsorientierte
Software zu verwenden und sich nicht mit Relikten alter Generationen
rumzuprügeln. Nachdem die letzten Wochen mit SVN etwas holprig waren, mir
allerdings halfen das prinzipielle System einer Versionsverwaltung zu
verstehen, tat ich mir Git an. Git. Der Name ist ja erstmal unterirdisch
wenn mans so auf sich wirken lässt. Ganz im Gegensatz zum Banner der
Projekt-Homepage [www.git-scm.com](http://www.git-scm.com), welches ich
sehr nett finde. Aber Schluss mit EyeCandy.

Erstellte <del>gezwungenermaßen</del> freiwillig mit ein paar (Obacht, zwei
Links in einem Wort)
[How](http://progit.org)-[To's](http://scie.nti.st/2007/11/14/hosting-git-repositories-the-easy-and-secure-way)
ein Git-Repository auf zwetschge.org. Mithilfe der How-Tos,
[Gitosis](http://wiki.dreamhost.com/Gitosis), git-daemon-run und git-core
war das <del>relativ schnell</del> geschafft. Allerdings kann ich mir beim
besten Willen nicht merken wie ich ein  Repository für ein neues Projekt
erstelle. An der Stelle setzt der Blogpost an.

```
Serverside:
$ mkdir /home/git/repositories/project.git #Simpler Ordner
$ cd /home/git/repositories/project.git #Selbsterklärend
$ git --bare init #ServerGitRepo bauen

Clientside:
$ cd /home/Code/OrdermitProjekt
$ git init #Projekt einlesen
$ git add . #Alle Inhalte adden
$ git commit -a -m "Inital commit of Software XY" #LokalCommit
$ git remote add origin git@server.com:project.git #RepoServer in .git hinterlegen
$ git push origin master #Push zum Server
```

```
#Bei Verwendung  von Gitosis - zuerst:
Gitosis:
$ gitosis-init < /tmp/pubkeyofmember.pub
$ vim gitosis.conf
[group Projectteam]
members = user@host #Letzen Inhalte von Public-SSHKey
writable = project #Projektname abgleitet von project.git
$ git commit -a -m "Gitosis update for new Project" #LokalCommit für Rechte
$ git push #Auf RepoServer pushen
```

