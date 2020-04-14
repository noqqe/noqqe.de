---
title: NFS Debugging
date: 2013-06-12T10:25:14
tags:
- Filesystems
---

Show all hosts that can connect to NFS

    showmount -a host
    showmount -e host

To troubleshoot an NFS mounting problem (NOT in order!) :

If you're automounting, try static mounting of the same filesystem,
to a different mount point, like /mnt/nfs or /mnt.  It'll probably give
more useful error messages.

Try a sniffer, like ethereal, tethereal, snoop or tcpdump -v.  Look
for NFS or RPC errors in the sniffer output.

Try truss/strace/par/trace
against rpc.mountd.  You probably don't want to do this with nfsd - it
tends to just sit in kernel space all the time.

Check your logs

Make sure you're exporting to and mounting from an FQDN.  Sometimes
weird things happen when you use short hostnames.

Try exporting "insecure", in case you have a host checking for a
specific port range.  Or alternatively, see if you can persuade the host
that's not using reserved ports, to use reserved ports - EG, on
AIX, this can be done with:

    echo "/usr/sbin/nfso -o nfs_use_reserved_ports=1" &amp;&amp; /etc/rc.net

Make sure the user doing the NFS mount isn't in too many groups.  If
you're in a large number of groups, NFS mounts can fail, seemingly
inexplicably.  You can usually check this with the "id" command.
If it's above some OS-specific
threshold (most likely 8, 16 or 32), then NFS may refuse to give a mount
due to the large number of groups.

Try unexporting everything, and reexporting.

Try completely shutting down NFS and restarting it

Make sure there isn't a firewall blocking some important traffic.
Sometimes even NFS clients will require accepting some incoming traffic,
initiated by the server.  This command can be very useful for this:

    nmap -sR -I RPC dcs.nac.uci.edu

It may or may not help to add -p1-65535 to the options.

I suggest running this on the server against the server, on the server
against the client, on the client against the client, and on the client
against the server - then compare the results.  The runs against the
client should be the same, and the runs against the server should be the
same.  If something is getting blocked over the network that isn't
blocked via localhost, then you can be pretty assured that there's a
firewall or something (network problem?) blocking some traffic.

You can expect the server to have greater RPC service requirements than the
client.  The client, if it is also an NFS server, may have the same RPC
services registered, but usually NFS will actually use a proper
subset of the RPC services on an NFS server (may even be a set of size 0
:).

If you're automounting, and you have static mounting working, there
are two scenarios to consider:

> On systems that have both automount and automountd programs,
> automountd is the daemon, and automount is a program that is supposed
> to make automountd notice changes in its maps.
> On systems that only have an automount program, automount is the
> daemon, and you need to kill and restart it (without</i> using the
> -9 signal!) to make it see changes.

Are all of the relevant daemons running?  You probably want something
like the following in rpcinfo -p:

     program vers proto   port
      100000    2   tcp    111  portmapper
      100000    2   udp    111  portmapper
      100021    1   udp  32775  nlockmgr
      100021    3   udp  32775  nlockmgr
      100021    4   udp  32775  nlockmgr
      100021    1   tcp  32768  nlockmgr
      100021    3   tcp  32768  nlockmgr
      100021    4   tcp  32768  nlockmgr
      100024    1   udp  32776  status
      100024    1   tcp  32769  status
      100011    1   udp    671  rquotad
      100011    2   udp    671  rquotad
      100011    1   tcp    690  rquotad
      100011    2   tcp    690  rquotad
      100003    2   udp   2049  nfs
      100003    3   udp   2049  nfs
      100003    2   tcp   2049  nfs
      100003    3   tcp   2049  nfs
      100005    1   udp    693  mountd
      100005    1   tcp    708  mountd
      100005    2   udp    693  mountd
      100005    2   tcp    708  mountd
      100005    3   udp    693  mountd
      100005    3   tcp    708  mountd

(the numbers in the left column are more significant than the names in
the right column)

From there, you can get probably to the daemon names using netstat -ap
and/or lsof.

Make sure that the actual daemon names sound NFS-related; sometimes a
non-RPC program will steal a port that rpcbind/portmap thought it could
allocate - but couldn't.

Alternatively, you can just run my [RPC Health](http://stromberg.dnsalias.org/~dstromberg/rpc-health.html)
script - but note that it won't detect missing services, only services
that are registered but not responding to a minimal test.

Try the mount with TCP or UDP, whichever you haven't tried already.
TCP should be better on long hauls or flakey networks, and UDP should be
better on close, reliable networks.  But if one isn't working, go ahead
and try the other anyway.

Are you using a flakey version of NFS?  EG, are both of the systems that
cannot communicate via NFS using the still-rough NFSv4 (Wed Feb 23
14:16:34 PST 2005)?  IIRC, idmap is indicative of NFSv4 on a Fedora Core
3 system.  NFSv4 reportedly worked better in FC2 than it does in FC3,
though yum -y update may have changed that by now.  It's probably worth
it to try at least NFS v2 and v3, and maybe v4 as well.

Try a different blocksize for read and/or write.  8192 is a good
number to try, if you haven't yet (most systems default to this).  8192
is -not- always optimal though.  Some sun systems used to crash if you
used a blocksize of 32768.  Also, some linux systems default to 1024,
which is a good choice on particularly flakey networks, or when you're
stuck with a poor network card.

Can you mount a different filesystem from the NFS server, but not the
one you want?

Are there permissions on the -mount-point-, underneath a mounted
filesystem, that are confusing matters?  I once saw an NFS problem that
turned out to be due to this on a SunOS 4.1.x system.

Do you have a firewall that is blocking ICMP packets inappropriately?
Some ICMP's are hazardous, but others can be essential to non-flakey
network communication.

Check showmount -e nfs.server.com

Check your netgroups, and NIS in general, if you're exporting to
netgroups.  Also try removing the netgroup export temporarily, and just
exporting to the host you need to have access from but isn't working.

If you have a large number of mounts, and suddenly subsequent mounts
start failing, and the same thing happens after a reboot, you may be
running out of privileged ports.

If you run man for each NFS daemon in turn, do they have an option for
cranking up verbosity?  If so, and you've gotten this far, you may as
well try it.  :)

Linux: Try enabling debugging facilities and checking for errors:

RPC debugging:

    echo 2048 > /proc/sys/sunrpc/rpc_debug
    grep . /proc/net/rpc/*/content
    ls -l /proc/fs/nfsd
    cat /proc/fs/nfs/exports

NFS debugging:

    ## turn on linux nfs debug
    echo 1 > /proc/sys/sunrpc/nfs_debug
    ## turn off linux nfs debug
    echo 0 > /proc/sys/sunrpc/nfs_debug

Facilities in perspective:

Actually, there is a whole bitmask of values you can use here in order
o selectively turn on or off parts of the debugging code.
See the NFSDBG_* defines in include/linux/nfs_fs.h.

There are similar bitmasks for the RPC, NLM (i.e. lockd) and nfsd
subsystems in include/linux/sunrpc/debug.h, include/linux/lockd/debug.h
and include/linux/nfsd/debug.h. These bitmasks act
on /proc/sys/sunrpc/rpc_debug, /proc/sys/sunrpc/nlm_debug
and /proc/sys/sunrpc/nfsd_debug respectively.

Note as I said earlier, though, this is really designed for debugging
purposes. There are no plans to convert it into an administrative tool.

Linux: Run this once while the NFS server is working, and then again when
the NFS server is having problems:

``` bash
cat /etc/exports
cat /proc/fs/nfsd/exports
grep . /proc/net/rpc/*/content
```

Post to any and all relevant mailing lists and newsgroups :)  Do
this sequentially, not in parallel - to keep the people you want help
from, from getting annoyed by reading and rereading the same message
over and over again unnecessarily.  Do not cross post.

Call the relevant vendors :)
