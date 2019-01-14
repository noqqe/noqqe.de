---
title: Bash Email Validation
date: 2011-06-07T15:38:56
tags:
- Programming
- RegEx
---

Validieren von Email in Bash

~~~ { .bash }
regex="^[a-z0-9!#\$%&'*+/=?^_\`{|}~-]+(\.[a-z0-9!#$%&'*+/=?^_\`{|}~-]+)*@([a-z0-9]([a-z0-9-]*[a-z0-9])?\.)+[a-z0-9]([a-z0-9-]*[a-z0-9])?\$"

i="test@terra.es"
if [[ $i =~ $regex ]] ; then
   echo "OK"
else
   echo "not OK"
fi
~~~