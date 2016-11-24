---
title: erb Templates
date: 2013-12-02T14:27:02.000000
tags: 
- Puppet
---


## If Else Entscheidungen

~~~
<%- if @role0 == "datanode" or @role0 == "namenode" -%>
command[check_load]=/usr/lib/nagios/plugins/check_load -w 25,22,20 -c 35,32,30
<%- else -%>
command[check_load]=/usr/lib/nagios/plugins/check_load -w 15,10,5 -c 30,25,20
<%- end -%>
~~~

## Template Limiter

~~~
<% i += 1 %> ## ruby code

<%= i += 1 %> ## mit print

<%- i += 1 %> ## suppress newline at front

<%- i += 1 -%> ## suppress newline at the end

<%## comment %>
~~~


## Looping durch eine Hiera Variable

braucht man ja

http://docs.puppetlabs.com/learning/templates.html

### Hiera

Im Endeffekt eine Liste daraus machen

~~~
hadoop::zookeeper:
  - "zk11.example.com"
  - "zk12.example.com"
  - "zk13.example.com"
~~~

### Manifest

~~~
class hadoop::zookeeper_base ($zookeepersrv = hiera(hadoop::zookeeper, '')) {

  file { "/etc/zookeeper/conf.noqqe/zoo.cfg":
    content => template("hadoop/zoo.cfg.erb"),
  }
}
~~~

### Template

Im template file \*.erb

~~~
<% @zookeepersrv.each do |x| -%>
server.<%= x.count %> <%= x %>:2888:3888
<% end -%>
~~~

### Testen des Templates

    erb -P -x -T '-' hadoop/templates/zoo.cfg.erb | ruby -c

## Loop mit Kommas in einer Zeile

~~~
<value><%- i = 0 ; @zookeepersrv.each do |x| -%><%= ',' if i > 0 %><%= x %>:2181<%- i += 1 -%><%- end -%></value>
~~~

In aufgesplittet mit Kommentaren

~~~
<value>
 <%- i= 0 ; @zookeepersrv.each do |x| -%> ## Loop
  <%= ',' if i > 0 %> ## Wenn erstes element, dann kein Komma
  <%= x %>:2181 ## Paste
  <%- i += 1 -%> ## Increase counter für Komma
 <%- end -%>
</value>
~~~

Ergebnis

~~~
-    <value>zk11.example.com:2181,zk12.example.com:2181,zk21.example.com:2181</value>
+    <value>zk11.example.com:2181</value>
~~~

oder

~~~
-    <value>zk11.example.com:2181,zk12.example.com:2181,zk21.example.com:2181</value>
+    <value>zk11.example.com:2181,zk12.example.com:2181,zk13.example.com:2181</value>
~~~
