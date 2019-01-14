---
date: '2016-02-14T20:19:42'
tags:
- pynacl
- code
- salsa20
- python
- crypto
- nacl
- blake2b
- opensource
title: Python Crypto - How do I Stromchiffre?
---

Ich schreibe schon seit Längerem eine Art Rewrite von cmddocs. Diesmal aber
mit MongoDB als Backend. Darin liegen meine Notizen, Wiki, Zitate,
Codeschnipsel und Journal. Wie genau das aussieht, wird ein anderer Post
beschreiben, wenn ich es mal OpenSource machen sollte...irgendwann(tm)

Auch wenn die MongoDB mit Authentication gesichert und der Transportweg mit
TLS verschlüsselt ist
([Data-in-Transit](https://en.wikipedia.org/wiki/Data_in_transit)),
verspürt man wie bei Full-Disk-Encryption den Drang die Daten auch auf der
Platte zu verschlüsseln.

Was noch übrig bleibt, ist
[Data-at-Rest](https://en.wikipedia.org/wiki/Data_at_rest) und dafür habe
ich mich für applikationsseitige Verschlüsselung entschieden, da ich
soetwas noch nicht gemacht habe.

{{< figure src="/uploads/2016/02/crypto.jpg" >}}

Den Content, den ich in "rvo" (Arbeitstitel..) einwerfe und absichern will,
möchte ich also mit einem Passwort verschlüsseln und dort ablegen. Dafür
sind aber mehrere Teilkomponenten nötig.

### Salsa20

Ich habe [Salsa20](https://de.wikipedia.org/wiki/Salsa20) in der
symmetrischen Variante als Algorithmus gewählt. Ich wollte irgendwas
modernes und nicht eines von 100.000 How Tos für AES in Python nachbauen.
In der deutschen Wikipedia taucht dafür das Wort "Stromverschlüsselung"
auf... *chchch*. Der Streamcipher bietet wohl Geschwindigkeit als auch
keine bekannten Angriffe die alle Phasen überwinden. Ich hab auch mal das
Paper von djb überflogen, aber natürlich kein Wort verstanden.

{{< figure src="/uploads/2016/02/salsa.png" >}}

Die Python Library [pyNaCL](http://pynacl.readthedocs.org/en/latest/),
nimmt einem dabei aber zum Glück einiges ab.

```python
import nacl.secret
import nacl.utils
key = nacl.utils.random(nacl.secret.SecretBox.KEY_SIZE)
box = nacl.secret.SecretBox(key)
msg = b"total laser"
nonce = nacl.utils.random(nacl.secret.SecretBox.NONCE_SIZE)
encrypted = box.encrypt(msg, nonce, encoder=nacl.encoding.HexEncoder)
```

Ich bin ja, was Crypto in der Entwicklung angeht, ziemlich unerfahren.
Deshalb wollte ich dem notwendigen Kram mir bekannter Ciphers, wie Content
zerstückeln für Block Ciphers und Padding, aus dem Weg gehen indem ich
eine "Stromverschlüsselung" *chchch* verwende.

Aber auch bei Salsa20 gibts für den Entwickler Eigenheiten, die man bei
der Implementierung berücksichtigen muss. Zum Beispiel die `nonce`. Nonce
ist ein String, der nur einmal pro Zugriff auf den Ciphertext verwendet
werden soll. Das heisst, auch wenn ich eine meiner Notizen lesend abrufe,
muss ich es neu verschlüsseln mit einem anderen 24 Byte String als `nonce`.

### Blake2b

Was mich ja nervt ist, dass Ciphers immer(?) eine fixe Größe eines Strings
als Secret Key erwarten. In den meisten Beispielen im Netz ist das Passwort
zur Verschlüsselung natürlich zufällig genau so lang, wie der
Algorithmus es erwartet.

{{< figure src="/uploads/2016/02/hash.jpg" >}}

Deshalb muss ein Passwort-Hash her, von denen es wieder n Varianten gibt.
Neben PBKDF2 (oder auch PDKDOMGWTFBBQ.. weil ich mir nie merken kann wie das
Acronym genau lautet) auch noch die "Stromchiffre" *chchch*
[Blake2](https://blake2.net). Eigentlich
ist dieser jetzt nicht so wirklich für Passwort Hashing geeignet, weil
Sachen wie Salts usw fehlen (korrigiert mich, sollte ich falsch liegen).
Zumindest wurde mir eben dieser ans Herz gelegt, um ihn in den Cipher
zu kippen, da er genau die 32 Byte auswirft, die ich für Salsa20 brauche.

```python
from pyblake2 import blake2b
key = blake2b(digest_size=32)
key.update(getpass.getpass("PASSWORD:"))
key = key.digest()
```

Einen normalen Workflow mit diesen beiden Komponenten findet ihr auf
[Github](https://gist.github.com/noqqe/cd9f8dc6477c7929f8b3)

### Indirektion

Nachdem ich den Blake2b Hash als Key für Salsa20 nutzen kann um meine
Dokumente zu verschlüsseln, stellen sich aber noch zwei weitere Probleme.

Erstens muss ich bei jedem Dokument das Passwort eingeben. Vertippe ich
mich, ist das Passwort für ein Dokument anders als für alle anderen.
Zweitens muss ich, sollte ich mich mal entscheiden mein Passwort zu ändern,
alle Dokumente entschlüsseln und mit dem neuen Passwort verschlüsseln.

Wie alle Probleme in der Informatik löst man sie wie? Genau, Indirektion!
Also generiere ich bei der ersten Benutzung einen Masterkey und
verschlüssle diesen mit einem Passwort, welches der User wählt.

{{< figure src="/uploads/2016/02/indirektion.png" >}}

So kann ich auch abfangen, ob das eingegebene Passwort (bei späterer
Benutzung) falsch oder richtig ist. Das Konzept ist natürlich jetzt auch
nicht auf meinem Mist gewachsen. Festplattenverschlüsselungssoftware macht
ähnliches. Wie geil wäre es auch, wenn du dein Disk Encryption Passwort
änderst und erstmal die ganze HDD neu gecrypted werden muss?

Meine Klasse "Crypto" in rvo enthält nun 4 Methoden: `init_master`,
`get_master`, `encrypt` und `decrypt`. Und funktionieren tuts auch noch.
Irre.
