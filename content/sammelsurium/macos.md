---
title: "macOS"
date: 2021-08-20T08:41:51+02:00
tags:
- OS
---

<!--more-->

## Screenshots

Nach `jpg` ändern

    defaults write com.apple.screencapture type jpg

Default Location ändern

    defaults write com.apple.screencapture location /path/

## Updates

Wenn der Arbeitgeber Updates blockt, kann man sie auch über Commandline
einspielen.


    sudo softwareupdate --install --all --restart
