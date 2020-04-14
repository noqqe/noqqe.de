---
title: systemd
date: 2018-05-30T16:20:21
tags:
- systemd
- Software
---

### Ein einzelner Service

Als Beispiel mal der von Grafana, da er sehr schön ist.

``` ini
[Unit]
Description=Grafana instance
Documentation=http://docs.grafana.org
Wants=network-online.target
After=network-online.target
After=postgresql.service mariadb.service mysql.service

[Service]
EnvironmentFile=/etc/default/grafana-server
User=grafana
Group=grafana
Type=simple
Restart=on-failure
WorkingDirectory=/usr/share/grafana
RuntimeDirectory=grafana
RuntimeDirectoryMode=0750
ExecStart=/usr/sbin/grafana-server                                                  \
LimitNOFILE=10000
TimeoutStopSec=20
UMask=0027

[Install]
WantedBy=multi-user.target
```

Wie gewohnt... erstmal mit einem einfachen Service interagieren.

``` bash
systemctl restart/start/stop/status/cat <service>.service
```

### Verzeichnisse

``` bash
/etc/systemd/system/
/lib/systemd/system/
/usr/lib/systemd/system/
```

### Commands

Config des Systemd reloaden

``` bash
systemctl daemon-reload
```

Alle laufenden Services anzeigen

``` bash
systemctl status
```

### Timer

Timer sind sozusagen cron(1) auf Steroiden und führt periodisch einen Service
immer wieder aus.

Alle Timer anzeigen

``` bash
systemctl list-timers
```

Timer Config

``` bash
$ systemctl cat <servicename>.timer
[Unit]
Wants=network.target

[Timer]
OnCalendar=*:0/5
OnUnitActiveSec=5min
```

Der Service wird ueber Pattern Matching des Timer Names gemachted

``` ini
$ systemctl cat <servicename>.service
[Unit]
Description=<servicename>
Wants=network.target

[Service]
ExecStart=/usr/local/bin/<command> <parameter>
StandardOutput=journal
Type=oneshot
User=root
WorkingDirectory=/tmp/
```

### Tempfiles

Man kann auch via Systemd tmpfiles anlegen lassen. Die configs dazu liegen
hier:

``` bash
/etc/tmpfiles.d/*.conf
/run/tmpfiles.d/*.conf
/usr/lib/tmpfiles.d/*.conf
```

und eines davon sieht dann ungefähr so aus:

``` bash
# Type Path        Mode UID  GID  Age Argument
  d    /var/run/mongodb 0755 mongodb mongodb - -
```

`Type` hat unendlich viele
[Optionen](https://www.freedesktop.org/software/systemd/man/tmpfiles.d.html)
