---
categories:
- bsd
- openbsd
- osbn
date: 2016-09-15T11:23:11+02:00
tags:
- Paste
- pastecat
- pastebin
- selfhosted
- linx-server
title: PasteBins und Quotas
---
Ich hatte längere Zeit einen kleinen Paste Daemon benutzt der
[pastecat](https://github.com/mvdan/pastecat) heißt. Er war wunderbar.
Hatte ein einfaches Interface mit Commandline Options, in Go geschrieben,
globale Speicherplatzbegrenzungen, Max-Age Settings und ein einfaches
Webinterface.

Was er leider nicht kann (und das mit Absicht) sind MIMETypes. Dafür
interessiert sich der Daemon nicht und ist auch nicht im Scope. In Firefox
endet das beim Anzeigen eines `.png` in einfach nur ASCII Kram auf dem
Bildschirm des Users. In Safari und Chrome war das okay. In Firefox nicht.

{{< figure src="/uploads/2016/09/wpm65xmf.jpg" >}}

Letztens war es dann wieder soweit. Ich schaute mich nach etwas anderem um.
Alle Services auf der
[awesome-selfhosted](https://github.com/Kickball/awesome-selfhosted#pastebins)
Liste durchgewühlt wollte ich eigentlich schon wieder aufgeben, da fand ich
in der Sektion "File Sharing and Synchronization" den Link zu
[linx-server](https://github.com/andreimarcu/linx-server). Es sah genau so
aus wie ich mir das vorgestellt hab.

Nur lässt sie dem User etwas zu viel Raum zu bestimmen **wie lange** und **wie
viele** Files hochgeladen werden.

Erstmal dem **wie viel** widmen. Puh. Ich glaube das letzte Mal als ich irgendwo ein Filesystem Quota
gesetzt habe war 2010. Alles kommt ja irgendwie wieder.

Quota Option in die `fstab` hängen.
```
# vim /etc/fstab
fbc4391bb6b72c36.k /home ffs rw,nodev,nosuid,userquota=/var/quotas/home.user 1 2
```

Mit `edquota` die Größe (4GB) setzen (natürlich dem separaten User `paste`
vorausgesetzt).
```
# edquota paste
Quotas for user paste:
/home: KBytes in use: 28990, limits (soft = 400000, hard = 400000)
        inodes in use: 61, limits (soft = 0, hard = 0)
```

Quota überprüfen
```
# quota -u paste
Disk quotas for user paste (uid 1006):
  Filesystem  KBytes    quota   limit   grace    files   quota   limit   grace
       /home   28990   400000  400000              61        0       0
```

Was das `expiry` Setting angeht (sprich: **wie lange** wird das File
aufgehoben) kreisen meine Gedanken immernoch um verschiedene
Möglichkeiten.

Ich möchte im Zweifel nicht dass irgendwelche Spammer Content hochladen und
`expire` auf `never` setzen.

Anfangs hatte ich mit ein bisschen `nginx` Zauber experimentiert. Da die
API einfach nur das HTTP Header Field `Linx-Expiry` zur Identifikation
verwendet.

``` nginx
location / {
  proxy_pass http://127.0.0.1:5005/;
  proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
  proxy_set_header Host $host;
  proxy_set_header Linx-Expiry 172800;
}
```

Das hat auch wunderbar funktioniert bis ich statt der API mal das
Webinterface statt der API benutzt habe. Diese benutzt nämlich statt
die eigene API, einen Post Request. Vergleiche

    curl -H "Linx-Expiry: 172800" -H "Linx-Randomize: yes" -T $1 https://p.n0q.org/upload/

und

    curl -X POST --data "expiry=86400&content=foo" https://p.n0q.org/upload

funktioniert beides, aber danach stellte sich heraus, dass `HTTP POST` Parameter
überschreiben mit `nginx` editieren garnicht so einfach ist. Bzw nicht
geht. Da in der Dokumentation hierzu klargestellt wird

> The ngx_http_rewrite_module module is used to change request URI using
> PCRE regular expressions, return redirects, and conditionally select
> configurations.

Nun ist Linx ja OpenSource. Also
[Issues](https://github.com/andreimarcu/linx-server/issues/100)
[geöffnet](https://github.com/andreimarcu/linx-server/issues/99) und als
Workaround erstmal die `Expiry` hardcoded selbst reingepatched.

```diff
diff --git a/upload.go b/upload.go
index b0bbd9f..0d66326 100644
--- a/upload.go
+++ b/upload.go
@@ -67,7 +67,7 @@ func uploadPostHandler(c web.C, w http.ResponseWriter, r *http.Request) {
     if r.Form.Get("randomize") == "true" {
       upReq.randomBarename = true
     }
-    upReq.expiry = parseExpiry(r.Form.Get("expires"))
+    upReq.expiry = parseExpiry("172800")
     upReq.src = file
     upReq.filename = headers.Filename
   } else {
@@ -81,7 +81,7 @@ func uploadPostHandler(c web.C, w http.ResponseWriter, r *http.Request) {
     }

     upReq.src = strings.NewReader(r.FormValue("content"))
-    upReq.expiry = parseExpiry(r.FormValue("expires"))
+    upReq.expiry = parseExpiry("172800")
     upReq.filename = r.FormValue("filename") + "." + extension
   }

@@ -163,7 +163,7 @@ func uploadRemote(c web.C, w http.ResponseWriter, r *http.Request) {
   upReq.src = resp.Body
   upReq.deletionKey = r.FormValue("deletekey")
   upReq.randomBarename = r.FormValue("randomize") == "yes"
-  upReq.expiry = parseExpiry(r.FormValue("expiry"))
+  upReq.expiry = parseExpiry("172800")

   upload, err := processUpload(upReq)

@@ -194,7 +194,7 @@ func uploadHeaderProcess(r *http.Request, upReq *UploadRequest) {
   upReq.deletionKey = r.Header.Get("Linx-Delete-Key")

   // Get seconds until expiry. Non-integer responses never expire.
-  expStr := r.Header.Get("Linx-Expiry")
+  expStr := "172800"
   upReq.expiry = parseExpiry(expStr)

 }
```

Beide Probleme erstmal gelöst. Es soll durchaus Menschen geben
die ihren Paste Service ohne Expiry und Limits betreiben, aber irgendwie
hab ich mich dann etwas reingesteigert :P Trotzdem was gelernt.
