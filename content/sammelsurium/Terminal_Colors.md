---
title: Terminal Colors
date: 2016-01-22T12:54:50.000000
tags: 
- terminal
- bash
- Programming
- ansii
- ascii
- colors
- color
---


16 Color ANSI

    for code in {30..37}; do \
    echo -en "\e[${code}m"'\\e['"$code"'m'"\e[0m"; \
    echo -en "  \e[$code;1m"'\\e['"$code"';1m'"\e[0m"; \
    echo -en "  \e[$code;3m"'\\e['"$code"';3m'"\e[0m"; \
    echo -en "  \e[$code;4m"'\\e['"$code"';4m'"\e[0m"; \
    echo -e "  \e[$((code+60))m"'\\e['"$((code+60))"'m'"\e[0m"; \
    done

256 Color ANSI

    for x in {0..255} ; do echo -ne "\e[38;5;${x}m" $x "\e[0" ; done

