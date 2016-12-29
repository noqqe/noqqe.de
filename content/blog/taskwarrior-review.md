---
comments:
- author: Anonymous
  content: "Die Idee klingt eigentlich gut. Mal schauen ob ich f\xFCr meine Todo-Listen
    so etwas ebenfalls hinbekomme."
  date: '2016-05-08T15:49:50.517502'
- author: Anonymous
  content: "Eigentlich bin ich ja ein grosser fan von CLI tools, aber als ich das
    letzte mal (vor ca. 2 Jahren) geschaut habe, gab es keine gescheite M\xF6glichkeit
    l, die tasks \xFCber mehrere Devices zu syncen. Weisst du was, wie es da aktuell
    aussieht?"
  date: '2016-05-08T16:36:02.730065'
- author: pbeckingham
  content: Nicely done.
  date: '2016-05-08T22:31:05.585023'
- author: Anonymous
  content: 'Taskwarrior hat sog. "Taskserver", die gibts auch als Service: https://freecinc.com/about'
  date: '2016-05-09T02:52:11.827145'
- author: noqqe
  content: thanks!
  date: '2016-05-09T08:36:42.924364'
- author: noqqe
  content: Was nutzt du denn?
  date: '2016-05-09T08:37:20.011658'
- author: noqqe
  content: genau das!
  date: '2016-05-09T08:37:32.115206'
- author: senden9
  content: "Aktuell trello.com. Hab schon nachgesehen. Die bieten eine API an. Eventuell
    schaue ich mir aber mal das Taskwarrior-\xD6kosystem an."
  date: '2016-05-09T17:54:38.428788'
date: '2016-05-08T09:52:27'
tags:
- openbsd
- python
- accomplished
- opensource
title: Taskwarrior Review
---

Wie schon oft erwähnt, nutze ich [Taskwarrior](https://taskwarrior.org) zum
Verwalten meiner Todos. Bei allen Todo Listen ist immer das Problem - zu
viel zu erledigen. Man fühlt sich als würde man gegen eine unendliche Menge
an Todos ankämpfen die niemals weniger wird.

{{< figure src="/uploads/2016/05/done.jpg" >}}

Vor ein paar Wochen las ich dann diesen Lifehacker
[Post](http://lifehacker.com/5960794/keep-your-completed-to-do-lists-on-a-done-wall-to-stay-motivated).

Im Grunde steht da nur "Sieh dir an was du geschafft hast". Deshalb hab ich
mir mit ein wenig Python auch soetwas gebaut.

``` python
import commands
import json
import simplemail

cmd = '/usr/local/bin/task end.after:today-1wk export'
mailaddr = 'user@example.com'

def parse(output):
    tasks = json.loads(output)
    message = u''
    for task in tasks:
        try:
            message += u'* ' + task["project"] + ": " + task["description"] + u'\n'
        except:
            message += u'* ' + task["description"] + u'\n'

    mail(message)
    return True

def mail(message):
    simplemail.Email(
        smtp_server = "localhost",
        from_address = mailaddr,
        to_address = mailaddr,
        subject = u'Accomplished tasks this week',
        message = message
    ).send()

def main():
    (status, output) = commands.getstatusoutput(cmd)
    output = unicode(output, 'utf8', errors='replace')

    if status:
        sys.stderr.write('Error running task command')
        return False

    parse(output)

if __name__ == "__main__":
    main()
```

Eigentlich mailt es mir nur die Tasks, die ich innerhalb der letzten
7 Tage erledigt habe.

{{< figure src="/uploads/2016/05/done2.jpg" >}}

Per Cronjob schicke ich mir diese Mail am Freitag um 19:00. Gutes Gefühl
die Mail zu lesen.