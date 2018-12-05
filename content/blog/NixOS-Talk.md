---
title: "NixOS Talk @ K4CG"
date: 2018-12-05T12:25:55+01:00
tags:
- NixOS
- Talk
- Hackerspace
- K4CG
---

Vor kurzem habe ich einen kleinen Vortrag zu [NixOS](https://nixos.org) in der
[K4CG](https://k4cg.org) gemacht.

Dafür habe ich (mal wieder) ein neues Presentation Tool verwendet, in diesem
Fall [mdp](https://github.com/visit1985/mdp)

Der Source sieht dann ungefähr so aus und ist dabei immernoch compilierendes
Markdown

```
-----

-> # NixOS: Configuration Management! <-

Um das System zu verändern

~~~
vim /etc/nixos/configuration.nix
nixos-rebuild switch
~~~

Demo Time!

-----
```

Was dann compiled ungefähr so aussieht

{{< figure src="/uploads/2018/12/nix.png" >}}

Die Folien als PDF kann man hier sehen

[Folien](/uploads/2018/12/nixos.pdf)

Feedback immer gerne :)
