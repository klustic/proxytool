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

## Use Case: `docker pull` through SOCKSv5 Tunnel

> These steps assume you have a SOCKS proxy listening on 127.0.0.1:1080, and Proxytool listening on 127.0.0.1:8080.

Docker does not provide native SOCKS proxy support, but it does provide HTTP/S proxy support. So you can use this tool to wrap your SOCKSv5 proxy with Proxytool as follows:

### Step 1: Add proxy settings for the docker daemon

Since images are pulled by the daemon and not the `docker` cli directly, you have to add daemon configurations for Proxytool and reload systemd:

```
# [root]# vim /etc/systemd/system/docker.service.d/http-proxy.conf
[Service]
Environment="HTTP_PROXY=http://127.0.0.1:8080"
Environment="HTTPS_PROXY=http://127.0.0.1:8080"
```

Then reload systemd:

```
[root]# systemctl daemon-reload
```

### [OPTIONAL] Step 2: Trust TLS certificates on the docker repository

This is required if the repo's TLS certificate doesn't have a valid trust chain:

```
[root]# vim /etc/docker/daemon.json
{
  "insecure-registries" : ["repo.internal.example.com"]
}
[root]# service docker restart
```

### Step 3: Pull container

```
docker pull repo.internal.example.com/core/container-with-all-the-secrets:latest
```
