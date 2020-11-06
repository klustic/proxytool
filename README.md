# Proxytool

Lighweight utility for wrapping a SOCKSv5 proxy with an HTTP(S) proxy.

## Usage

```
## Download and build
go get github.com/klustic/proxytool
go build github.com/klustic/proxytool

## Usage
$ ./proxytool -h
Usage of ./proxytool:
  -http string
    	Where to set up the HTTP/HTTPS proxy (default "127.0.0.1:8080")
  -socks string
    	Location of SOCKS proxy (default "127.0.0.1:1080")
  -verbose
    	Toggle verbose logging to STDOUT (default true)


$ ./proxytool 
Listening for HTTP/HTTPS proxy connections and forwarding through SOCKS proxy.
- HTTP/HTTPS proxy ..: 127.0.0.1:8080
- SOCKSv5 proxy .....: 127.0.0.1:1080
- DNS resolutions ...: remote (not currently configurable)

2020/11/06 15:30:16 [001] INFO: Got request / icanhazip.com GET http://icanhazip.com/
2020/11/06 15:30:16 [001] INFO: Sending request GET http://icanhazip.com/
2020/11/06 15:30:17 [001] INFO: Received response 200 OK
2020/11/06 15:30:17 [001] INFO: Copying response to client 200 OK [200]
2020/11/06 15:30:17 [001] INFO: Copied 14 bytes to client error=<nil>
...
```