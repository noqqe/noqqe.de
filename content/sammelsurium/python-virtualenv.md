---
title: Python Virtualenv
date: 2016-04-11T17:12:03.305000
tags: 
- Programming
- Python
- virtualenv
---


#### Usage

Ein neues virtualenv erstellen

    mkvirtualenv rvo

In ein virtualenv welchseln

    workon rvo

Virtualenv verlassen

    deactivate

Alle Venvs anzeigen lassen

    workon

Andere python version

    mkvirtualenv --python=python3.5 matomat

##### Installation

Es gibt einen einfacheren Wrapper dafür, der vieles einfacher macht.

    sudo pip install virtualenvwrapper

Configuration:

    if [[ -x /usr/local/bin/virtualenvwrapper.sh ]]; then
      export WORKON_HOME=~/.pyenvs
      export PROJECT_HOME=~/Code
      source /usr/local/bin/virtualenvwrapper.sh
    fi

Wenn man automatisch in das Verzeichnis wechseln will, kann man auch

    VIRTUALENVWRAPPER_WORKON_CD

setzen.
