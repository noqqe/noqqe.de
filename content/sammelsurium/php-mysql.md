---
title: PHP MySQL
date: 2012-01-21T10:56:40
tags:
- Databases
- MySQL
---

Performance-wise it doesn't matter what you use. The difference is that mysql_fetch_object returns object:

``` php
while ($row = mysql_fetch_object($result)) {
    echo $row->user_id;
    echo $row->fullname;
}
```

mysql_fetch_assoc() returns associative array:

``` php
while ($row = mysql_fetch_assoc($result)) {
    echo $row["userid"];
    echo $row["fullname"];
}
```

and mysql_fetch_array() returns array:

``` php
while ($row = mysql_fetch_array($result)) {
    echo $row[0];
    echo $row[1] ;
}
```
