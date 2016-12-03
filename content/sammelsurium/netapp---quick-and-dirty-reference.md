---
title: NetApp - Quick and Dirty Reference
date: 2012-04-10T14:56:18
tags: 
- OS
---

* `sysconfig -a` : shows hardware configuration with more verbose information
* `sysconfig -d` : shows information of the disk attached to the filer
* `version` : shows the netapp Ontap OS version.
* `uptime` : shows the filer uptime
* `dns info` : this shows the dns resolvers, the no of hits and misses and other info
* `nis info` : this shows the nis domain name, yp servers etc.
* `rdfile` : Like "cat" in Linux, used to read contents of text files/
* `wrfile` : Creates/Overwrites a file. Similar to "cat > filename" in Linux
* `aggr status` : Shows the aggregate status
* `aggr status -r` : Shows the raid configuration, reconstruction information of the disks in filer
* `aggr show_space` : Shows the disk usage of the aggreate, WAFL reserve, overheads etc.
* `vol status` : Shows the volume information
* `vol status -s` : Displays the spare disks on the filer
* `vol status -f` : Displays the failed disks on the filer
* `vol status -r` : Shows the raid configuration, reconstruction information of the disks
* `df -h` : Displays volume disk usage
* `df -i` : Shows the inode counts of all the volumes
* `df -Ah` : Shows "df" information of the aggregate
* `license` : Displays/add/removes license on a netapp filer
* `maxfiles` : Displays and adds more inodes to a volume
* `aggr create` : Creates aggregate
* `vol create <volname> <aggrname> <size>` : Creates volume in an aggregate
* `vol offline <volname>` : Offlines a volume
* `vol online <volname>` : Onlines a volume
* `vol destroy <volname>` : Destroys and removes an volume
* `vol size <volname> [+|-]<size>` : Resize a volume in netapp filer
* `vol options` : Displays/Changes volume options in a netapp filer
* `qtree create <qtree-path>` : Creates qtree
* `qtree status` : Displays the status of qtrees
* `quota on` : Enables quota on a netapp filer
* `quota off` : Disables quota
* `quota resize` : Resizes quota
* `quota report` : Reports the quota and usage
* `snap list` : Displays all snapshots on a volume
* `snap create <volname> <snapname>` : Create snapshot
* `snap sched <volname> <schedule>` : Schedule snapshot creation
* `snap reserve <volname> <percentage>` : Display/set snapshot reserve space in volume
* `/etc/exports` : File that manages the NFS exports
* `rdfile /etc/exports` : Read the NFS exports file
* `wrfile /etc/exports` : Write to NFS exports file
* `exportfs -a` : Exports all the filesystems listed in /etc/exports
* `cifs setup` : Setup cifs
* `cifs shares` : Create/displays cifs shares
* `cifs access` : Changes access of cifs shares
* `lun create` : Creates iscsi or fcp luns on a netapp filer
* `lun map` : Maps lun to an igroup
* `lun show` : Show all the luns on a filer
* `igroup create` : Creates netapp igroup
* `lun stats` : Show lun I/O statistics
* `disk show` : Shows all the disk on the filer
* `disk zero spares` : Zeros the spare disks
* `disk_fw_update` : Upgrades the disk firmware on all disks
* `options` : Display/Set options on netapp filer
* `options nfs` : Display/Set NFS options
* `options timed` : Display/Set NTP options on netapp.
* `options autosupport` : Display/Set autosupport options
* `options cifs` : Display/Set cifs options
* `options tcp` : Display/Set TCP options
* `options net` : Display/Set network options
* `ndmpcopy <src-path> <dst-path>` : Initiates ndmpcopy
* `ndmpd status` : Displays status of ndmpd
* `ndmpd killall` : Terminates all the ndmpd processes.
* `ifconfig` : Displays/Sets IP address on a network/vif interface
* `vif create` : Creates a VIF (bonding/trunking/teaming)
* `vif status` : Displays status of a vif
* `netstat` : Displays network statistics
* `sysstat -us 1` : begins a 1 second sample of the filer's current utilization (crtl - c to end)
* `nfsstat` : Shows nfs statistics
* `nfsstat -l` : Displays nfs stats per client
* `nfs_hist` : Displays nfs historgram
* `statit` : beings/ends a performance workload sampling [-b starts / -e ends]
* `stats` : Displays stats for every counter on netapp. Read stats man page for more info
* `ifstat` : Displays Network interface stats
* `qtree stats` : displays I/O stats of qtree
* `environment` : display environment status on shelves and chassis of the filer
* `storage show <disk|shelf|adapter>` : Shows storage component details
* `snapmirror intialize` : Initialize a snapmirror relation
* `snapmirror update` : Manually Update snapmirror relation
* `snapmirror resync` : Resyns a broken snapmirror
* `snapmirror quiesce` : Quiesces a snapmirror bond
* `snapmirror break` : Breakes a snapmirror relation
* `snapmirror abort` : Abort a running snapmirror
* `snapmirror status` : Shows snapmirror status
* `lock status -h` : Displays locks held by filer
* `sm_mon` : Manage the locks
* `storage download shelf` : Installs the shelf firmware
* `software get` : Download the Netapp OS software
* `software install` : Installs OS
* `download` : Updates the installed OS
* `cf status` : Displays cluster status
* `cf takeover` : Takes over the cluster partner
* `cf giveback` : Gives back control to the cluster partner
* `reboot` : Reboots a filer
