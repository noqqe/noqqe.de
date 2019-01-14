---
title: Ansible Cheatsheet
date: 2015-05-17T20:50:32
tags: 
- Software
- ansible
---

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