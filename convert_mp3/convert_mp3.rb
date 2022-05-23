#!/usr/bin/env ruby

require "securerandom"

inp_filename = ARGV[0]
segment_time
save_filename = 
    File.basename(
        inp_filename,
        File.extname(inp_filename)
    )

segment_time = 17

if inp_filename
    if File.exist?(inp_filename)
        `mkdir ./songs/#{save_filename}`
    
        puts `
        ffmpeg -i #{inp_filename.strip} -c:a libmp3lame -b:a 128k -map 0:0 -f segment -segment_time #{segment_time} -segment_list ./songs/#{save_filename}/#{save_filename}.m3u8 -segment_format mpegts ./songs/#{save_filename}/chunk_%03d.ts
        `.strip
    
        puts "\nConvert finished!"
    else
        puts "File not found"
    end
else
    puts "./convert_mp3 mp3_file"
end