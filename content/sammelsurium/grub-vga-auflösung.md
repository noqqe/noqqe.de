---
title: Grub VGA Auflösung
date: 2011-06-08T10:39:17
tags: 
- HowTos
---

~~~
title Debian GNU/Linux, kernel 2.6.24-1-686
root (hd0,0)
kernel /boot/vmlinuz-2.6.24-1-686 root=/dev/hda1 ro quiet vga=791
initrd /boot/initrd.img-2.6.24-1-686
~~~

The new parameter added is vga=791, which means a resolution of 1024x768
with a colour depth of 16 bits per pixel (65,536 colours). Some other VGA
modes would be:

~~~
791 - 1024x768, 16 bit
792 - 1024x768, 24 bit
794 - 1280x1024, 16 bit
795 - 1280x1024, 24 bit
~~~
