---
title: "Was NixOS kann - Ein Webserver Beispiel"
date: 2020-06-24T08:16:47+02:00
tags:
- NixOS
---

Mit [NixOS 20.03](https://nixos.org/nixos/manual/release-notes.html#sec-release-20.03) ist für den
Apache2 Webserver [ganz schön was passiert](https://github.com/NixOS/nixpkgs/commit/79215f0df1ddf4bf0db7dc4c5789f8dae9f9bb02).

Die Änderungen sind tiefgreifend, aber toll. Deswegen will ich kurz die alte
Konfiguration mit der Neuen vergleichen, denn ich denke das sich daraus
wunderbar die Vorteile von NixOS aufzeigen lassen.

<!--more-->

## Was brauche ich?

Als Beispiel schauen wir uns den `vHost` vom [chaostreff-nuernberg.de](https://www.chaostreff-nuernberg.de) an.
Die Anforderungen sind sehr simpel:

* Statisches Webseite Hosting
* HTTP -> Port 80
* HTTPS -> Port 443
* HTTP -> HTTPS Weiterleitung
* Let's Encrypt Zertifikat

Also wirklich **kein** Hexenwerk.

## Der alte Weg. (<=19.09)

Ich brauche also 2 vHosts. Einen, der die `ACME` Challenge aus dem
entsprechenden Directory ausliefert, und einen anderen, der dann HTTPS die
eigentliche Page ausliefert.

``` nix
services.httpd = {
 enable = true;
 logFormat = "combined";
 virtualHosts = [
  {
    documentRoot = "/var/www/chaostreff-nuernberg.de";
    listen = [{port = 80;}];
    hostName = "chaostreff-nuernberg.de";
    serverAliases = [ "www.chaostreff-nuernberg.de" ];
    globalRedirect = "https://k4cg.org";
    extraConfig = ''
    Alias /.well-known/acme-challenge/ /var/www/challenges/
    <Directory /var/www/challenges/>
       AllowOverride None
       Require all granted
       Satisfy Any
    </Directory>
    '';
  }
  {
    hostName = "chaostreff-nuernberg.de";
    serverAliases = [ "www.chaostreff-nuernberg.de" ];
    documentRoot = "/var/www/chaostreff-nuernberg.de/";
    listen = [{port = 443;}];
    enableSSL = true;
    sslServerCert = "/usr/local/acme-tiny/k4cg.org.crt";
    sslServerKey = "/usr/local/acme-tiny/k4cg.org.key";
    sslServerChain = "/usr/local/acme-tiny/intermediate.crt";
  }
  ];
};
```

Und die Let's Encrypt Zertifikate werden via eines Cronjobs mittels
[acme-tiny](https://github.com/diafygi/acme-tiny) regelmässig erneuert. Das
findet sich allerdings nicht in der NixOS Konfiguration wieder. Will ich auch
im Detail garnicht zeigen, weil es einfach ein hässliches Shell Script ist.

## Der neue Weg (>=20.03)

In der neuen Variante haben wir `security.acme`
[bekommen](https://nixos.org/nixos/options.html#security.acme). Dieses NixOS
Optionset stellt einen ACME Client ([Lego](https://github.com/go-acme/lego)) bereit. Das wäre an sich noch keine Zauberei.

Was hier geschafft wurde ist ist das Kunststück, diesen Client direkt in die
Apache2 Konfiguration zu integrieren! Durch das abgreifen der Infos und
Default Werte zwischen `security.acme` und `services.httpd` ergibt sich eine
super schlanke und elegante Konfiguration.

``` nix
services.httpd = {
 enable = true;
 logFormat = "combined";
 virtualHosts = {
  "chaostreff-nuernberg.de" = {
    hostName = "chaostreff-nuernberg.de";
    serverAliases = [ "www.chaostreff-nuernberg.de" ];
    documentRoot = "/var/www/chaostreff-nuernberg.de/";
    addSSL = true;
    enableACME = true;
  };
};

security.acme.certs = {
  "chaostreff-nuernberg.de" = {};
};
```

Was wegfällt ist ein ganzer vHost, manuelles Setzen der SSL Zerts, Challenges
Directory spezifieren und `listen` Ports auflisten. Ganz zu
schweigen von einen Cronjob der dann ein Shell Script ausführt um die
LetsEncrypt Zertifikate zu erneuern. All das ist jetzt "under the hood" der
Neuimplementierung des `httpd` Services.

Was man vielleicht garnicht so sieht: `security.acme` fragt nicht nur
"chaostreff-nuernberg.de" für die Zertifikate an, sondern fügt auch noch alle
weiteren DNS Namen aus dem
`services.httpd.virtualHosts."chaostreff-nuernberg.de".serverAliases` Array
als DNS alternative Names in das angefragte Zertifikat ein, obwohl diese
Informationen dem eigentlichen Tool (Lego) so direkt nicht zur Verfügung
stehen. Geil oder?

Besonders hervorheben will ich noch die neuen [SSL Optionen](https://nixos.org/nixos/options.html#httpd+ssl)

``` nix
services.httpd.virtualHosts.<name>.addSSL  # http + https
services.httpd.virtualHosts.<name>.forceSSL # http -> https
services.httpd.virtualHosts.<name>.onlySSL # https
```

und auch die `lego` Zertifikatserneuerungen über je einen Systemd-Timer pro
Zertifikat.

www.chaostreff-nuernberg.de war natürlich nicht der einzige `vHost` in unserem Setup. Durch die
Umstellung und dem nutzen der Vorteile von der deklarativen Settings haben
wir nun **~120 Zeilen Konfiguration weniger**. Die paar Bytes weniger sind
aber nicht der Punkt. Es ist leichter zu pflegen und weniger Moving Parts die
ich als Admin jonglieren, äh... orchestrieren muss.
