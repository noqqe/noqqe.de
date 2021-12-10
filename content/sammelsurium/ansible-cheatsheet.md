---
title: Ansible Cheatsheet
date: 2015-05-17T20:50:32
tags:
- Software
- ansible
---

## General

Apply a playbook

    ansible-playbook nichtparasoup.yml

Limit a playbook on tags

    ansible-playbook userenv.yml --tag git

Limit a playbook on a host

    ansible-playbook userenv.yaml --limit noc.example.com

## Common Task Oneliners

Ansible can be very handy for those kind of things.

Execute a generic command.

    ansible mongo-test -m command -a "ls /tmp"

Restart a service

    ansible app-prod -m service -a 'name=nrpe state=restarted'

Fetch a file from remote hosts

    ansible mongo-prod -m fetch -a 'src=/var/log/mongodb/mongodb-shardsrv.log dest=/tmp/'

Copy files on a remote location

    ansible mongo-prod -m copy -a 'src=~/foo.txt dest=/root/foo.txt owner=root'

Execute a shell command

    ansible app-prod -m shell -a 'sed -i "s#plugins/check_mem#plugins/check_mem -C#" /etc/nrpe.cfg'

Remove a user

    ansible app-test -m user -a 'name=johnd state=absent remove=yes'

Delete a file

    ansible -i hosts app-test -m file -a 'path=/tmp/trollololo state=absent'

## Playbook Cheatsheet

Install a package

```yaml
- name: Install java 11
  ansible.builtin.package:
    name: java-11-amazon-corretto
    state: installed
```

Create a directory

```yaml
- name: Create directory
  ansible.builtin.file:
    path: "/tmp/foo"
    state: directory
    mode: '0755'
```

Create a symlink

```yaml
- name: Create a Java Home Symlink
  ansible.builtin.file:
    src: /usr/lib/jvm/java-11-amazon-corretto.x86_64
    dest: /home/java
    state: link
```

Create a file with content

```yaml
- name: Configure .profile for user
  copy:
    dest: "/home/user/.profile"
    mode: u=rwx,g=rx,o=rx
    owner: user
    content: |
      export PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin
      alias select=_select
      alias l='ls -lah'
      source /etc/profile.d/*.sh
```

Loop. Mehrere Items erstellen mit einem Task

```yaml
- name: Create mount point directories for shares
  ansible.builtin.file:
    path: "{{ item }}"
    state: directory
    mode: '0755'
  with_items:
    - /moo
    - /foo
    - /bar
```

Mount a volume (AWS EFS Volume, in this case)

```yaml
- name: Mount /move
  ansible.posix.mount:
    path: /move
    src: i8234u234:/
    fstype: efs
    opts: defaults,_netdev
    state: present
```

