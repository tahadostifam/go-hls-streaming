#!/usr/bin/env ruby

fname = ARGV[0]

if fname
    cmd_cb = 
    `
    ffmpeg -i #{cmd_cb} -c:a libmp3lame -b:a 128k -map 0:0 -f segment -segment_time 10 -segment_list outputlist.m3u8 -segment_format mpegts output%03d.ts
    `.strip
else
    puts "./convert_mp3 mp3_file"
end