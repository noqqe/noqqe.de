---
title: "Python requests hinter'm Proxy"
date: 2023-12-04T10:16:44+01:00
tags:
- python
- tenacity
- retry
---

Der Firmenproxy hinter dem ich arbeite ist oftmals in Verbindung mit VPN
ziemlich shaky. Das heisst Verbindungen werden nicht verlässlich aufgebaut.

<!--more-->

Gerade arbeite ich ~~gegen~~ mit der BitBucket Server API. Ich mache viele
Requests, lese User Permissions aus und reagiere darauf. Dabei schmiert mein kleines Script immer wieder ab. 

```
Traceback (most recent call last):
  File "/Users/noqqe/.asdf/installs/python/3.9.18/lib/python3.9/site-packages/urllib3/connection.py", line 203, in _new_conn
    sock = connection.create_connection(
  File "/Users/noqqe/.asdf/installs/python/3.9.18/lib/python3.9/site-packages/urllib3/util/connection.py", line 85, in create_connection
    raise err
  File "/Users/noqqe/.asdf/installs/python/3.9.18/lib/python3.9/site-packages/urllib3/util/connection.py", line 73, in create_connection
    sock.connect(sa)
ConnectionRefusedError: [Errno 61] Connection refused
```

Und das ganze Scrapen geht von vorne los. Jetzt habe ich eine ziemlich coole Lösung gefunden, wie man **einfach**
Retries einbauen kann. Mittels
[tenacity](https://tenacity.readthedocs.io/en/latest/)!


Anstatt irgend welche `try/except` Lösungen zu bauen, kann ich die Library
installieren und einen Decorator verwenden. In diesem definiert man wie lange
und mit welchen Backoffintervallen die Funktion retried werden soll.

```python
@retry(
    stop=stop_after_attempt(10),
    wait=wait_exponential(multiplier=1, min=1, max=60)
)
def _get_repo_users(project, repo):
  body = requests.get(baseurl+"/projects/{}/repos/{}/permissions/users/".format(project,repo))
  return json.loads(body.text)
```

Super einfach. 
