---
title: Puppet Client
date: 2013-06-30T16:44:25.000000
tags: 
- Puppet
---


## Default Ports um zum Master zu connecten

Port 8140 ist default

aber mit `--masterport 8888` kann das umspezifiziert werden.

## Installation der Pakete

    cd /root/
    wget http://apt.puppetlabs.com/puppetlabs-release-wheezy.deb
    dpkg -i puppetlabs-release-wheezy.deb apt-get update && apt-get install
    puppet

In der /etc/puppet/puppet.conf in der [main]-Sektion hinzufügen:

    server=vm30-puppet.example.com
    environment = production

Zertifikat anfordern:

    puppet agent --noop --test --waitforcert 30 sed -i -e
    's#START=no#START=yes#' /etc/default/puppet /etc/init.d/puppet start

Nun auf dem entsprechenden puppet master den Fingerprint vergleichen und
auf dem puppet master den Certificate Request signieren:

    ## puppet cert list
    > "server01.example.com" (E0:F9:54:6F:DF:F7:D7:58:78:68:F7:D1:0D:E6:57:E3)
    ## puppet cert sign server01.example.com

Auf dem Client in der /etc/default/puppet setzen:

    START=yes

und den Agent starten:

    /etc/init.d/puppet start

fertig.


## Deploying

Bei der Migration größerer Umgebungen ist dringend davon abzuraten, die
Puppet Manifeste "einfach so" zu schreiben und zu pushen. Ein
Entwicklungsvorgang kann wie folgt abgewickelt werden:

modules klonen:

    git clone git@gitlab.example.com:puppet-modules.git

manifests klonen

    git clone git@gitlab.example.com:manifests.git

Neues Manifest anlegen und Änderungen einpflegen

    vim manifests/customers/$kunde.pp

Die Änderungen vor dem Commit ins Repo zum Ausprobieren auf den Server
kopieren. Ich bevorzuge dafuer diesen "Shot":

    H="192.168.178.129" ; ssh $H "rm -r /tmp/puppet ; mkdir /tmp/puppet"  ;
    rsync -va manifests $H:/tmp/puppet/ ; rsync -va puppet-modules
    $H:/tmp/puppet/

Auf der Maschine die geänderten Manifeste ausprobieren:

    puppet apply --noop --show_diff
    --modulepath=/tmp/puppet/puppet-modules/ /tmp/puppet/manifests/site.pp

Wenn alles gut lief, kann die erste Initialisierung auch gleich lokal
angewendet werden.

    puppet apply --modulepath=/tmp/puppet/puppet-modules/
    /tmp/puppet/manifests/site.pp

Dieser Vorgang kann nun für alle Server wiederholt werden, die sich im
Setup/Umgebung befinden. So lassen sich selektiv Hosts testen und einzeln
"kaputt machen" anstatt Alles mit einem Riesen-Commit parallel zu
zerstören.
