---
title: pip
date: 2016-12-08T14:55:04
tags:
- Python
- Programming
- pip
---

Register a user and create config

    [distutils]
    index-servers = pypi

    [pypi]
    repository: https://pypi.python.org/pypi
    username: user
    password: xxx


## Python2

Register a new package

    python setup.py register -r pypi

Upload a new version of a package

    python setup.py sdist upload -r pypi

## Python3

First build the sdist

    python setup.py sdist

Then install `twine` and upload it using this new tool

    pip install twine
    twine upload dist/*

