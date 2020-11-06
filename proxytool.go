package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/klustic/proxytool/socks"
)

func main() {
	socksAddr := flag.String("socks", "127.0.0.1:1080", "Location of SOCKS proxy")
	httpAddr := flag.String("http", "127.0.0.1:8080", "Where to set up the HTTP/HTTPS proxy")
	verbose := flag.Bool("verbose", true, "Toggle verbose logging to STDOUT")
	flag.Parse()

	fmt.Println("Listening for HTTP/HTTPS proxy connections and forwarding through SOCKS proxy.")
	fmt.Printf("- HTTP/HTTPS proxy ..: %s\n", *httpAddr)
	fmt.Printf("- SOCKSv5 proxy .....: %s\n", *socksAddr)
	fmt.Println("- DNS resolutions ...: remote (not currently configurable)\n")

	socksDialer := socks.NewDialer("tcp4", *socksAddr)
	proxy := goproxy.NewProxyHttpServer()
	proxy.Tr.Dial = socksDialer.Dial
	proxy.ConnectDial = socksDialer.Dial
	proxy.Verbose = *verbose
	log.Fatal(http.ListenAndServe(*httpAddr, proxy))
}
