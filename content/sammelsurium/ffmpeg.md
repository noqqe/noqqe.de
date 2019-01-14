---
title: ffmpeg
date: 2013-09-23T13:59:32
tags:
- Software
- ffmpeg
---

Remove sound from video file

    ffmpeg -i <input_file> -vcodec copy -an <output_file>

Video um 90 Grad rotieren

    ffmpeg -i <input_file>  -vf "transpose=1" <output_file>