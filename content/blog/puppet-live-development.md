---
title: "Puppet Live Development"
date: 2018-02-02T14:28:05+01:00
tags:
- Puppet
- puppetserver
---

Beruflich habe ich 2017 so gut wie ausschliesslich
[Puppet](https://puppetlabs.com) gemacht. Das hat mich vielerlei Dinge gelehrt
und schon fast ein bisschen was von Development gehabt. Dazu aber evtl. in
anderen Posts mehr. Worum es in diesem Post gehen soll ist eine Methode zum
angenehmen Entwickeln von Puppet Code.

### Das Problem mit Config Management

Gerade im Configuration Management läuft es etwas anders als
beim "traditionellen" Entwickeln.

Ich kann meine Codebase leider nicht einfach "bauen" und den Build dann lokal
ausführen. Ich kann leider genausowenig meine automatisierte Testsuite laufen
lassen um sicherzustellen das das von mir geschriebene Feature auch wirklich
läuft (von Syntax Errors mal abgesehen).

{{< figure src="/uploads/2018/01/cfgmgmt.jpg" >}}

Was ich brauche ist ein Client im Environment auf dem meine Changes ausgerollt
werden können und ein Server der den `Catalog` compiled in dem alle `Resources`
für den anfragenden Host enthalten sind. Das trifft auf meine Client Maschine
leider nicht zu, sofern ich nicht plötzlich einen HDFS Node mit DNS/Mail/NTP
Settings unseres Datacenters auf meinem MacBook betreiben will.

Was sind also die Alternativen? Ich könnte anfangen einfach das Puppet Repo zu
klonen, meine Changes machen, nach `master` pushen und hoffen das ich keine
Fehler gemacht habe. Der Master holt sich das Repo und die Change sind live.
Ich könnte sehen was auf den Nodes passiert. Das ist natürlich totaler Quatsch
und denke ich muss das erklären warum dem so ist. Dinge werden schief gehen und
1000 messy-Commits später gehts dann irgendwann wenn nicht vorher schon das
Datacenter in Flammen steht.

Ich könnte mich auch via SSH auf dem Puppet Master einloggen und meine Changes
im Working Directory editieren. Die Operation am offenen Herzen würde zwar die
Commit-Mess lösen, aber was hierbei alles schief gehen kann muss ich wohl auch
niemandem erklären.

### Was also tun?

`Puppetserver` (die Java Implementation) besitzt die wunderbare Fähigkeit
mehrere `Environments` zu unterstützen.

```
/etc/puppetlabs/code/enviromments/production   <-- Working Dir
/etc/puppetlabs/code/environments/JIRA-1337
/etc/puppetlabs/code/environments/JIRA-2342
```

Standardmäßig benutzt jeder Agent das Environment `production`, bedeutet aber
das ich das auch umkonfigurieren kann. Um mein neues Puppet Feature zu
entwickeln, klone ich das Repo mit anderem Namen in das Directory und
installiere meine Libraries nach.

```
git clone git@server:puppet.git JIRA-1337
cd JIRA-1337
git checkout -b JIRA-1337
librarian-puppet install
```

Hier kann ich meine Changes machen und auf Testing/Staging Maschinen
ausprobieren ob es funktioniert. Die Nodes im Puppet Environment sprechen so
schliesslich mit dem exakt selben Master, von dem sie auch sonst ihre Config
bekommen. Das hat nochmal extra Charme wenn man die SSL Client Zertifikate, die
jeder Host benötigt, im Hinterkopf hat. Alles was ich in der Comanndline beim
`One-Time Run` ändern muss ist das Environment.

```
$ ssh <staging-datanode>
$ puppet agent -t --noop --environment JIRA1337
[Changes get applied....]
```

--noop ist dabei der Dry-Run.

Sobald ich zufrieden mit meinem Feature bin, kann ich in einen neuen Remote
Branch pushen, Kollegen reviewen in Stash meinen Pull-Request und schliesslich
kann ich ihn (hoffentlich) mergen. Durch Branch Namen und evtl Commit Messages
habe ich dann auch sowohl in Stash als auch in Jira die Links zum PR
hergestellt und der Kreis schliesst sich dann. Nachvollziehbarkeit, Yeah!

### More convenience please!

Das alles ist natürlich nicht manuell nötig. Ich schrieb also ein script
`puppet-create-env` welches als Parameter die Jira Ticketnummer an der man
gerade arbeitet entgegen nimmt und tue all das was ich gerade eben beschrieben
habe. Als extra Feature geht das Script auch durch alle alten `Environments`
und überprüft mittels `git reflog` wann das letzte mal Changes vorgenommen
wurden und löscht Environments die älter als 5 Tage sind.

Das hat es uns im Team extrem erleichtert zusammenzuarbeiten und gleichzeitig
unabhängig voneinander zu sein. Nachteil ist natürlich das alle auf dem Puppet
Master ihren `vim` laufen haben, aber das ließe sich theoretisch über ssh FUSE
mount auch beheben.

## Andere Konstrukte

Es gibt natürlich andere Methoden um mit Puppet zu arbeiten. Früher hat man
hauptsächlich die ganze Codebase auf den Puppet Node kopiert und dort dann
lokal mittels `puppet apply` ausführt. Beschrieben
[hier](/sammelsurium/puppet-client/), was mit
[Hiera](https://docs.puppet.com/hiera/) und
[PuppetDB](https://docs.puppet.com/puppetdb/) aber ziemlich schwierig geworden
ist.

Eine andere Variante ist [r10k](https://github.com/puppetlabs/r10k), was uns
aber viel zu aufgeblasen erschien. Im Grunde benutzt aber auch r10k den Multi
Environment Ansatz.

Ebenfalls denkbar sind Puppet-Dev Kisten in denen statt eines `--environment
XXX` eben `--server <puppet-dev>` auf Agent-Seite mitgegeben wird.

Ich kenne auch Teams die ein Set aus Virtualbox VMs mittels
[Vagrant](https://www.vagrantup.com) auf ihrem Laptop herumtragen und damit ihr
Produktions Env nachsimulieren, aber das scheint sehr träge und fehleranfällig
zu sein.

Oder falls jemand noch andere Ideen hat, schreibt mir eine Mail!
