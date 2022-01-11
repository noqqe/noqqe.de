---
title: "AWS ECS Webshell via Tomcat"
date: 2022-01-11T11:13:39+01:00
tags:
- AWS
---

In AWS [ECS](https://aws.amazon.com/ecs/) kann man Docker Container
hochfahren, ohne dass man sich um das Host OS kümmern muss. Das ist meistens
super cool und einfach.

<!--more-->

Aber wenn man doch mal in die Verlegenheit kommt, während des Aufsetzens in
die Maschine blicken zu wollen (`netstat`, Logfiles, Connectivity
überprüfen...) hat man keine Chance.

{{< figure src="/uploads/2022/01/sandbox.jpg" >}}

Naja, nicht ganz. Wo man arbitrary Code ausführen kann gibt es immer eine
Möglichkeit.

Ich hatte das oben beschriebene Bedürfnis, in den Container reinschauen zu
wollen. Logs prüfen. Mounts und File Permissions checken.  Was also tun?

Webshell! Ich rolle einfach eine zusätzliche kleine Tomcat Anwendung in
meinen Container um dann via Browser Befehle im Container ausführen zu
können.

Dazu wird eine kleine Datei `index.jsp` erstellt, welche die Shell zur
Verfügung stellen wird.

```html
<FORM METHOD=GET ACTION='index.jsp'>
<INPUT name='cmd' type=text>
<INPUT type=submit value='Run'>
</FORM>
<%@ page import="java.io.*" %>
<%
   String cmd = request.getParameter("cmd");
   String output = "";
   if(cmd != null) {
      String s = null;
      try {
         Process p = Runtime.getRuntime().exec(cmd,null,null);
         BufferedReader sI = new BufferedReader(new
InputStreamReader(p.getInputStream()));
         while((s = sI.readLine()) != null) { output += s+"</br>"; }
      }  catch(IOException e) {   e.printStackTrace();   }
   }
%>
<pre><%=output %></pre>
```

Um aus diesem `index.jsp` File ein Tomcat WAR zu bauen, ist nichts nötig,
außer ein kleiner `zip` Command.

```
zip webshell.war index.jsp
```

Dieses WAR Archiv kann dann (zusammen mit der normalen Applikation) in den
Docker Container eingebaut werden.

```Dockerfile
FROM tomcat:9.0.55-jre11-temurin-focal
ENV TZ="Europe/Berlin"

# Install Debuggin Tools
RUN apt-get update ; apt-get install -y telnet netcat bash

# Deploy Tomcat Applications
COPY webshell.war /usr/local/tomcat/webapps/webshell.war
```

Ist der Container gebaut und in die private Registry hochgeladen, kann er in die ECS Task Spezifikation (wie üblich)
referenziert werden. Wenn der Container hochgefahren ist, sollte das so
aussehen.

{{< figure src="/uploads/2022/01/webshell.png" >}}

Man kann in aller Ruhe seinen Debugging Tätigkeiten nachgehen.

Und bitte, baut die Scheisse wieder aus wenn ihr den Internet Access enabled.
