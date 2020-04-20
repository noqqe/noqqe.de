---
title: erb Templates
date: 2013-12-02T14:27:02
tags:
- Puppet
- Software
---

## If Else Entscheidungen

``` ruby
<%- if @role0 == "datanode" or @role0 == "namenode" -%>
command[check_load]=/usr/lib/nagios/plugins/check_load -w 25,22,20 -c 35,32,30
<%- else -%>
command[check_load]=/usr/lib/nagios/plugins/check_load -w 15,10,5 -c 30,25,20
<%- end -%>
```

## Template Limiter

``` ruby
<% i += 1 %> ## ruby code

<%= i += 1 %> ## mit print

<%- i += 1 %> ## suppress newline at front

<%- i += 1 -%> ## suppress newline at the end

<%## comment %>
```

## Count Schleife

``` ruby
# reduce count by 1 because we start at 0
<% cpus = @processors['count'].to_i - 1 -%>

# loop through all cpus
<% for x in 0..cpus do -%>
w /sys/devices/system/cpu/cpu<%= x %>/cpufreq/scaling_governor - - - - performance
<% end -%>
```

## Looping durch eine Hiera Variable

braucht man ja

[Puppet Templates](http://docs.puppetlabs.com/learning/templates.html)

### Hiera

Im Endeffekt eine Liste daraus machen

``` yaml
hadoop::zookeeper:
  - "zk11.example.com"
  - "zk12.example.com"
  - "zk13.example.com"
```

### Manifest

``` puppet
class hadoop::zookeeper_base ($zookeepersrv = hiera(hadoop::zookeeper, '')) {

  file { "/etc/zookeeper/conf.noqqe/zoo.cfg":
    content => template("hadoop/zoo.cfg.erb"),
  }
}
```

### Template

Im template file \*.erb

``` ruby
<% @zookeepersrv.each do |x| -%>
server.<%= x.count %> <%= x %>:2888:3888
<% end -%>
```

### Testen des Templates

    erb -P -x -T '-' hadoop/templates/zoo.cfg.erb | ruby -c

## Loop mit Kommas in einer Zeile

``` ruby
<value>
 <%- i= 0 ; @zookeepersrv.each do |x| -%> ## Loop
  <%= ',' if i > 0 %> ## Wenn erstes element, dann kein Komma
  <%= x %>:2181 ## Paste
  <%- i += 1 -%> ## Increase counter für Komma
 <%- end -%>
</value>
```

Ergebnis

``` xml
+    <value>zk11.example.com:2181</value>
+    <value>zk11.example.com:2181,zk12.example.com:2181,zk13.example.com:2181</value>
```

## Hiera in Template Komplettbeispiel

Das hiera File (hieradata/tiers/production.yaml)

``` yaml
---
profiles::daemon::username: user
profiles::daemon::password: passw0rd
```

Das Manifest (site/profiles/manifests/daemon.pp)

``` puppet
class profiles::daemon (
  $username,
  $password,
) {

  file { '/etc/daemon/config':
    mode    => mode,
    owner   => root,
    group   => root,
    content => template("profiles/daemon/config.erb")
  }

}
```

Das Template (site/profiles/templates/daemon/config.erb)

``` yaml
memorylimit: 9001
user: <%= @username %>
password: <%= @password %>
```
