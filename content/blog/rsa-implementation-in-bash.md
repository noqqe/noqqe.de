---
comments:
- author: faheus
  content: wow, such read, very crypto, amaze
  date: '2014-01-03T21:29:11.496623'
- author: noqqe
  content: "Achja, wichtig ist noch, dass der gew\xE4hlte Publickey immer kleiner
    als p*q (also max) sein muss. "
  date: '2014-01-04T18:36:50.851435'
- author: onli
  content: "Vll interessiert dich das: Habe das vor Jahren auch mal in Bash (mit bc)
    gemacht. Nat\xFCrlich genauso wenig nutzbar, aber ein kleines bisschen kompletter,
    wenn ich das richtig sehe: \nhttps://gist.github.com/onli/8320277"
  date: '2014-01-08T18:10:55.285556'
- author: noqqe
  content: yeah! richtig nice und wesentlich kompletter! total gut!
  date: '2014-01-08T18:32:12.456048'
- author: onli
  content: "Danke :)\n\nMuss zugeben, ich finds selbst lustig, was mein damaliges
    Ich da gebaut hat - die richtige Mischung aus total f\xFCrchterlich (alleine die
    bescheuerte Grundidee, das f\xFCr ne ernsthafte Abgabe in Bash zu machen) und
    dann doch irgendwo clevere Umsetzung, die ich heute noch mag. Muss mich richtig
    gereizt haben."
  date: '2014-01-08T18:38:23.110133'
date: '2014-01-03T19:14:00'
tags:
- development
- shell
- reference
- cryptography
- rng
- privatekey
- crypto
- rsa
- pubkey
- bash
title: RSA Implementation in Bash
---

Teil 2 der Serie Crypto in Bash implementieren, heute:
[RSA](https://en.wikipedia.org/wiki/RSA_\(algorithm\)). Schon vor Längerem
hatte ich mal [ROT13 in purem Bash geschreiben](/blog/2011/08/13/rot13-verschlusselung-in-bash/), einfach um
zu sehen wies funktioniert.

{{< figure src="/uploads/2014/01/cryptography.jpg" >}}

Auf dem #30c3, den ich letzte Woche besuchen durfte, ging es natürlich auch
viel, viel um Crypto.  Die Do's and Dont's sozusagen. Ein
[Vortrag](http://events.ccc.de/congress/2013/Fahrplan/events/5502.html) war dabei besonders nett,
da er als eine Art Leitfaden zu verstehen war, wie man es nicht machen sollte.
Unter anderem RSA/AES selbst implementieren. Das hab ich dann aber einfach
mal gemacht.

### Basic RSA

Der Code ist so einfach wie nur irgendwie möglich gestaltet, um ihn verständlich
zu halten. Ich habe deshalb bewusst auf Funktionen und Sanity Checks verzichtet.

``` bash
p=37    # prime 1
q=89    # prime 2
e=25    # public key
d=2281  # private key
msg=$1  # the message
max=$((p * q)) # calc max for flattening

# show original message
echo "orig msg: $msg"

# encrypt message
i=$msg
for x in $(seq 2 $e); do
    msg=$((msg*i))      # multiply
    msg=$((msg%max))    # flatten
done
echo "encr msg: $msg"

# decrypt message
i=$msg
for x in $(seq 2 $d); do
    msg=$((msg*i))      # multiply
    msg=$((msg%max))    # flatten
done
echo "decr msg: $msg"
```

Manchem mag auffallen, wie im
[Debian-Stil](http://www.debian.org/security/2008/dsa-1571) ziemlich
statischer "Zufall" im Header kodiert ist. Für den Anfang ist das okay um
zu verstehen wie RSA so tickt.

``` bash
$ ./rsa.bash 999
orig msg: 999
encr msg: 2146
decr msg: 999
```

### RSA + Privkey Generierung

Was im obigen Beispiel passiert ist, ist natürlich sehr trivial, da die Keys
vordefiniert sind und hinten hinaus die allerwenigste Magie passiert.
Der interessante Teil ist eher die Schlüsselgenerierung.

``` bash
p=$1    # prime
q=$2    # prime
msg=$3  # the message
max=$((p * q)) # calc max for flattening
phi=$(((p-1)*(q-1))) # calc phi
pubprimes=( 1901 1907 1913 1931 1933 1949 1951 1973 1979 \
  7841 7853 7867 7873 7877 7879 7883 7901 7907 7919 \
  127  131  137  139  149  151  157  163  167  173 \
  6311 6317 6323 6329 6337 6343 6353 6359 6361 6367 \
  2221 2237 2239 2243 2251 2267 2269 2273 2281 2287 )

# choose random prime for pubkey e
e=${pubprimes[$RANDOM % ${#pubprimes[@]}]}

# extended euclid algorithm. borrowed and simplified from
# https://sites.google.com/site/algorithms2013/home/number-theory/extended-gcd-bash
function ext_euclid_algo() {
    agcd=$1
    b=$2
    x=0
    lastx=1
    y=1
    lasty=0
    while [ $b != 0 ]; do
        tempb=$b
        tempx=$x
        tempy=$y

        quotient=$(($agcd / $b))
        b=$(($agcd - $(($b * $(($agcd / $b))))))
        x=$(($lastx-$(($quotient*$x))))
        y=$(($lasty-$(($quotient*$y))))

        agcd=$tempb
        lastx=$tempx
        lasty=$tempy
    done

    echo $((lastx+$2))
}

# calculate private key
d=$(ext_euclid_algo $e $phi)

# print key environment
echo "pub: $e"
echo "priv: $d"
echo "phi: $phi"
[...]
letzter Teil wie oben
[...]
```

Worauf ich außerdem verzichtet habe ist die Auswahl einer
richtigen/anständigen Primzahl für den Publickey, da das Problem weder
logisch noch programmatisch besonders reizvoll ist. Zumindest was den RSA
Scope angeht.  Im RNG Scope natürlich eine ganz andere Geschichte.

Usage des Skripts hat sich nun auch von nur der Message hin zu Primzahl1,
Primzahl2, Message erweitert.

``` bash
$ ./rsa.bash 911 43 8234
pub: 1913
priv: 21977
phi: 38220
orig msg: 8234
encr msg: 33207
decr msg: 8234
```

Das gesamte Skript gibts nochmal auf
[Github](https://gist.github.com/noqqe/8245645). Bei Fehlern in der
Implementierung freue ich mich über Kommentare, aber bitte keine "da oben ist
eine Timing Attacke möglich"-Comments. Ich patche hier nicht OpenSSL.
Was dieses mal mehr gilt, als beim ROT13 Skript: Nicht für Dinge benutzen.
