---
title: awk
date: 2011-07-26T12:38:13.000000
tags: 
- Programming
---


Eine kleine Sammlung von Schnippseln die hilfreich sein könnten.

## Range von Dingen ausgeben

    awk '{for(i=3;i<=NF;++i)print $i}' 

## Genereller Aufbau

    awk 'program' input-file1 input-file2 ... `
    awk -f program-file input-file1 input-file2 ...`

    `rule/pattern { action }`
    `rule/pattern { action }`

## Simple Examples

* Print the length of the longest input line:

    `awk '{ if (length($0) > max) max = length($0) } END { print max }' data`

* command-line variable assignment, like this:

    `$ awk -v sq="'" 'BEGIN { print "Here is a single quote <" sq ">" }'`
    `-| Here is a single quote <'>`

* print lines with "foo" in it

    $ awk '/foo/ { print $0 }' BBS-list
    -| fooey        555-1234     2400/1200/300     B
    -| foot         555-6699     1200/300          B
    -| macfoo       555-6480     1200/300          A
    -| sabafoo      555-2127     1200/300          C

## For Loop for each $1, $2 of line

~~~
awk -F, '{ for (i = 1; i <= NF; i++)
         print $i
}'
~~~

Example

~~~
$ echo "foo foo fo teo ewr " | awk ' { for (i = 1; i <= NF; i++) print $i}'
foo
foo
fo
teo
ewr
~~~

## Quotes in awk print

`\x27` heisst das Zauberwort

~~~
echo $x | awk -F, '{print "date.addColumn(\x27" $2 "\x27, \x27" $1 "\x27);" }'
~~~
