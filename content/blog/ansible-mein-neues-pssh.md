---
comments:
- author: marvin
  content: ich werde auch jeden tag mehr und mehr fan von ansible. gibt da soviel
    zu entdecken und soviele anwendungszwecke.
  date: '2015-07-13T11:42:32.819846'
date: '2015-07-12T18:47:14'
tags:
- development
- management
- adminstration
- python
- bsd
- ansible
- puppet
- ssh
- config
title: Ansible - Mein neues Pssh
---

### pssh

[pssh](https://code.google.com/p/parallel-ssh/) war für eine Menge Tätigkeiten
lange Tool der Wahl. Deployments, Services restarten, Files verteilen, User löschen,
Hotfixes usw.

    pssh -h ~/Code/mongoserver.txt -l root -t 350 -i "/etc/init.d/mongodb restart"

Der Nachteil ist aber allerdings, dass bei `pssh` Listen mit
Hostnamen in Form von Files gepflegt werden müssen. Das wird schnell
unübersichtlich und aufwändig.

<!--more-->

{{< figure src="/uploads/2015/07/ssh.jpg" >}}

### Ansible Hosts

Deshalb bin ich vor ca. einem Monat zu Ansible gewechselt. Vorrangig wegen des
Host-Matchings.

``` ini
[mongodb]
mongodb01.example.org
mongodb02.example.org
mongodb03.example.org

[li]
li-jmp-prod
li-jmp-dev
li-jmp-stage

[n0q]
noc.n0q.org ansible_python_interpreter=/usr/local/bin/python
aax.n0q.org ansible_python_interpreter=/usr/local/bin/python

[work:children]
li
mongodb
```

Hosts in diesem File können  mit Regex, Excludes und Gruppen von Ansible aus
direkt angesprochen werden. Auch Gruppen die andere Gruppen enthalten sind
möglich. Für Details [hier](http://docs.ansible.com/intro_patterns.html).

### Ansible als pssh Ersatz

Auch wenn ich `ansible` nicht so nutze wie die Meisten (nämlich als full-blown
Config Management) so ist es doch extrem handlich.

Zum Beispiel einen User auf allen Servern der Gruppe "work" löschen

    ansible work -m user -a 'name=alice state=absent remove=yes'

Wie oben im Beispiel alle MongoDB Prozesse neustarten

    ansible mongodb -m service -a 'name=puppet state=restarted'
    ansible mongodb -m service -a 'name=mongodb state=restarted'

Oder einen `nrpe`-Check mittels Shell Command bearbeiten:

    ansible n0q -m shell -a 'sed -i "s#plugins/check_mem#plugins/check_mem -C#" /etc/nrpe.cfg'

Oder ein bestimmtes File von allen angegebenen Hosts herunterladen

    ansible n0q -m fetch -a 'src=/etc/bashrc dest=/tmp/fetched'

Es fühlt sich ein bisschen an, wie die UNIX-API für alles was man so braucht.