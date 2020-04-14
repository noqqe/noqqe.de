---
title: exiftool
date: 2013-06-24T16:00:30
tags: 
- Software
- exiftool
- exif
---

## EXIF Data

Ausgeben:

    exiftool -a -u -g1 gasthauszumbernd.jpg

entfernen:

    exiftool -all= image.jpg

## Print EXIF GPS Data as decimal

YEAH

    $ exiftool  -gpslatitude -gpslongitude  -n /pfad
    GPS Latitude                    : 43.7407694444444
    GPS Longitude                   : 7.42793055555556

oder fuer ganze Ordner

    exiftool -gpslatitude -gpslongitude -n -T DIR > out.txt

Zum Beispiel so

    exiftool -gpslatitude -gpslongitude -n -T /Users/noqqe/Pictures/Photos\ Library.photoslibrary/Masters/2015/08/08/20150808-220735/ | uniq | wc -l

Complete Workflow

    exiftool -gpslongitude -gpslatitude -n -T /Users/noqqe/Pictures/Photos\ Library.photoslibrary/Masters/ -r  | uniq | awk '{print $2" "$1}' > /tmp/foo
    grep -v -- "- -" /tmp/foo | gsed -e 's/\s/, /' -e 's/^/["f", /' | gawk '{print $0 ", " FNR "],"}'']'

## Meta Data

``` bash
$ exiv2 gasthauszumbernd.jpg
File name       : gasthauszumbernd.jpg
File size       : 2793568 Bytes
MIME type       : image/jpeg
Image size      : 3264 x 2448
Camera make     : Apple
Camera model    : iPhone 4S
Image timestamp : 2013:06:24 11:25:20
Image number    :
Exposure time   : 1/657 s
Aperture        : F2.4
Exposure bias   :
Flash           : No, compulsory
Flash bias      :
Focal length    : 4.3 mm (35 mm equivalent: 35.0 mm)
Subject distance:
ISO speed       : 50
Exposure mode   : Auto
Metering mode   : Multi-segment
Macro mode      :
Image quality   :
Exif Resolution : 3264 x 2448
White balance   : Auto
Thumbnail       : image/jpeg, 9844 Bytes
Copyright       :
Exif comment    :
```

Mit entfernten Meta Daten

``` bash
$ exiv2 gasthauszumbernd.jpg
File name       : gasthauszumbernd.jpg
File size       : 2781280 Bytes
MIME type       : image/jpeg
Image size      : 3264 x 2448
gasthauszumbernd.jpg: No Exif data found in the file
```