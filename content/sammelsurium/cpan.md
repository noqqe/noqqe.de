---
title: "Cpan"
date: 2018-03-15T11:14:36+01:00
tags:
- Software
---

# Install

interactive

```
# cpan
Terminal does not support AddHistory.

cpan shell -- CPAN exploration and modules installation (v1.9205)
ReadLine support available (maybe install Bundle::CPAN or Bundle::CPANxxl?)

cpan[1]> install JMX::Jmx4Perl
```

shell

``` bash
perl -MCPAN -e "install JMX::Jmx4Perl"
```

shell ohne tests

``` bash
perl -MCPAN -e "CPAN::Shell->notest('install', 'JMX::Jmx4Perl')"
```