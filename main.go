package main

import (
	"flag"
	"net/http"

	_ "expvar"
	_ "net/http/pprof"
)

var (
	addr = flag.String("http", ":80", "listen address")
)

func init() {
	flag.Parse()
}

func main() {
	http.ListenAndServe(*addr, nil)
}
