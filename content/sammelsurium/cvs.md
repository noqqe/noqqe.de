---
title: cvs
date: 2015-11-05T14:29:45
tags: 
- Software
- cvs
---

Checkout (always over ssh)

    cvs -qd host:/src/cvsroot get .

Show commits

    cvs log -rVERSION_1_0:

Show tags

    cvs status -v

Nice other Cheatsheet
http://www.slac.stanford.edu/grp/cd/soft/cvs/cvs_cheatsheet.html
