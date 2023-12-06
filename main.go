package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main(){

    proxy := goproxy.NewProxyHttpServer()
    proxy.Verbose = true

    proxy.OnRequest(goproxy.DstHostIs("twitter.com")).DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
        fmt.Println("Got request from "+req.Host)
        return req, nil
    })


    log.Fatal(http.ListenAndServe(":3003", proxy))
}

