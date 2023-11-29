---
title: Ansible
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

Often I want to test a ansible snippet locally and struggle with the
structure. So here it is.
```yaml
---
- name: "Apply locally"
  hosts: localhost
  connection: local
  tasks:
  - name: test command
    ansible.builtin.shell: touch /tmp/testkey
    args:
      executable: /bin/bash
      warn: no
      creates: /tmp/testkey

```

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

Create Task dependency

```yaml
- name: Ensure replicaset rs0 exists
  #[...]
  register: rsinit

- name: Create local admin
  #[...]
  when:
    - rsinit is changed

```

Fetch AWS Secret

```yaml

- name: Fetch AWS Secrets for Admin
  set_fact:
    password: "{{ lookup('amazon.aws.aws_secret', 'documentdb_root.password', nested=true) }}"
    username: "{{ lookup('amazon.aws.aws_secret', 'documentdb_root.username', nested=true) }}"

- name: Ensure replicaset rs0 exists
  community.mongodb.mongodb_replicaset:
    login_host: localhost
    login_user: {{username}}
    login_password: {{password}}
```

Check if OS Env var is set

```yaml
- name: Check if Variables are defined.
  fail:
    msg: "Environment variable AWS_SECRET_ACCESS_KEY is not defined or empty"
  when: lookup('env', "AWS_SECRET_ACCESS_KEY") | length < 0
```

Check if OS Env var contains a certain string

```yaml
- name: Check if correct environment is set in AWS
  fail:
    msg: "Environment variable AWS_ENV does not match {{ aws_env }}. Did you configure your environment correctly?"
  when: lookup('env', "AWS_ENV") != aws_env
```
