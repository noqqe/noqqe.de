---
title: "Ansible Replace über mehrere Files"
date: 2019-07-16T13:45:35+02:00
tags:
- ansible
---

Ich komme in der Arbeit leider zu selten in den Genuss etwas `ansible` zu
machen, aber heute war es doch mal wieder so weit.

Gesucht habe ich ein Playbook welches Debian Jessie -> Stretch Updates macht.
Ja, Buster ist bereits released, ich fühle mich bereits schlecht, macht es
nicht noch schlimmer mit dem Hinweis darauf als Reaktion auf diesen Blogpost.

Gefunden habe ich dann
[dist-upgrade.yaml](https://github.com/debops/debops/blob/d05a7d949b104ea279e0779b6849ac2307ac518a/ansible/playbooks/tools/dist-upgrade.yml)
von [DebOps](https://github.com/debops/debops/). Leider sieht dieses
Playbook nur Repos in `/etc/apt/sources.list` vor.

```
- name: Change current release in APT sources
  replace:
    dest: '/etc/apt/sources.list'
    regexp: '{{ dist_upgrade_current_release }}'
    replace: '{{ dist_upgrade_new_release }}'
```

Da meine Server aber alle per Puppet mit Repos bestückt werden, befinden sich
dabei einzelne Files in `/etc/apt/sources.list.d/`. Was ich also tun musste
war, die RegEx von `replace` auf mehrere Files anwenden.

```
- name: Find all apt sources.list files
  find:
    paths: "/etc/apt/"
    patterns: "*.list"
    recurse: yes
  register: repos

- name: Change current release in APT sources
  replace:
    dest: "{{ item.path }}"
    regexp: '{{ dist_upgrade_current_release }}'
    replace: '{{ dist_upgrade_new_release }}'
  with_items:
    - "{{ repos.files }}"
```

Die Doku zum
[find](https://docs.ansible.com/ansible/latest/modules/find_module.html#find-module)
Ansible Modul war dabei sehr hilfreich, weil ich auch nicht wusste welche Art
von Datentyp das Resultset von `find` dann ist. Kurz noch etwas durch die
Doku geschaut wie man Variablen in einem Dict abgreift und schon wars fertig.

Mir gefällts, ob es den Menschen von DebOps auch taugt werden wir in [PR #898](https://github.com/debops/debops/pull/898) sehen
