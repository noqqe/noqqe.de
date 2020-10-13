---
title: "Direnv"
date: 2020-10-13T10:50:34+02:00
tags:
- direnv
- terraform
---

Zwischenzeitlich habe ich ca. 10-15 verschiedene Terraform Codebases auf
meinem MacBook. Jedes davon wird verwendet, ich entwickle daran.

Damit kommen auch verschiedene AWS Accounts und/oder User zum Tragen. Die
Anmeldung an diese funktioniert über Umgebungsvariablen. Credentials in der
Codebase gehen nicht.

<!--more-->

Jedes mal diese Variablen ändern ist sehr mühsam. `vim` auf die `fish` Config
auf, ändern, `source`n und los.

Deshalb bin ich vor einiger Zeit zu [direnv](https://direnv.net) gewechselt.
`direnv` verspricht Umgebungsvariablen zu setzen, sobald man in ein
bestimmtes Directory wechselt.

```
> cat .envrc
export TF_ENV=xx
export AWS_ACCESS_KEY=xxx
export AWS_SECRET_ACCESS_KEY=xxx
export TF_VAR_access_key=$AWS_ACCESS_KEY
export TF_VAR_secret_key=$AWS_SECRET_ACCESS_KEY
```

In der minimalsten Form sieht das `.envrc` File wie ein Shell Skript aus. Am
Anfang fühlt sich das sogar etwas komisch an, wenn man die `fish` Shell
verwendet. Diese Syntax wäre dort garnicht unterstützt.

Wenn ich dann aber in das entsprechende Verzeichnis wechsle:

{{< figure src="/uploads/2020/10/direnv.png" >}}

Zieht mir `direnv` meine für dieses Projekt benötigten Environment Variables
an und ich kann gleich losarbeiten.

`direnv` kann auch noch einiges mehr. So sind die eingebauten
"[Plugins](https://direnv.net/man/direnv-stdlib.1.html)" für
bestimmte Programmiersprachen oder z.B. `PATH_add` anschaut, merkt man das `direnv`
eben genau nicht "nur" Shell ist, das gesourced wird.

Coole Sache, spart mit viel Mühe.
