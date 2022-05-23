# GoLang HLS Streaming

### Mp3 Convert to TS-Chunks
```bash
ruby ./convert_mp3/convert_mp3.rb sweet_dreams.mp3
```

### HLS Stream Server
```bash
go run ./hls_stream_server/main.go
```

You can get stream by using this HLS-Client -> [hlsplayer.net](https://www.hlsplayer.net)