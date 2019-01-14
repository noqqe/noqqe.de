---
aliases:
- /archives/496
date: '2009-02-18T20:12:43'
slug: mysql-in-php
tags:
- development
- mysql
- zufall
- linux
- abfrage
- apache2
- hardware
- databases
- php
title: MySQL in PHP
---

Zuerstmal, lief [mein Test über MySQL](http://seufz.wordpress.com/2009/02/12/a-byte-of-mysql-eine-kurzreferenz/)
in der Schule super :) Nachdem ich jetzt sowohl in PHP als auch MySQL
unterrichtet wurde, hat mich dann doch intressiert wie sich die beiden in
Kombination verhalten :) Die Datenbank die ich für
"[A-byte-of-MySQL](http://zwetschge.org/publications/A-byte-of-MySQL.pdf)"
erstellt habe diente als Test-Objekt:

``` php
mysql_connect('server','user','passwort');
# Verbinde zu mysql auf Server mit Benutzer und folgendem Passwort

# Wie wähle ich die zu verwendende Datenbank aus?
mysql_select_db("arbeit");

# Wie definiere ich den Befehl der auszführen ist?
$sql="SELECT * FROM kollegen order by rand();";

# Wie führe ich ihn aus?
$result = mysql_query($sql) OR die(mysql_error());

# Wie gebe ich das Ergebnis aus?
while ($row = mysql_fetch_assoc($result))
{
echo $row['ID']." ".$row['Name']." ".$row['Gehalt']." ".$row['Bereich']."
";
}
# Schleife: Für jeden Eintrag in der Datenbank eine Zeile erzeugen.
# Verbindung schliessen.
mysql_close();
?>
```

Komplett:

``` php
<html>
<?php
mysql_connect('server','user','passwort');
mysql_select_db("arbeit"); //Auswahl der DB
$sql="SELECT * FROM kollegen order by rand();";
$result = mysql_query($sql) OR die(mysql_error());
while ($row = mysql_fetch_assoc($result))
{
echo $row['ID']." ".$row['Name']." ".$row['Gehalt']." ".$row['Bereich']."<br>";
}
?>
</html>
```

Ergebnis:  [http://zwetschge.org/stuff/sql.php](http://zwetschge.org/stuff/sql.php)
