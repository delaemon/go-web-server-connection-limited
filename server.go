package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"golang.org/x/net/netutil"
)

func Default(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		time.Sleep(time.Millisecond * 500) // for watch netstat
		io.WriteString(w, "Default")
	})

	p, _ := strconv.Atoi(port)
	http_config := &http.Server{
		Addr: fmt.Sprintf("localhost:%d", p),
	}
	err := http_config.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func Limited(port, limit string) {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		time.Sleep(time.Millisecond * 500) // for watch netstat
		io.WriteString(w, "Limited")
	})

	p, _ := strconv.Atoi(port)
	addr := fmt.Sprintf("localhost:%d", p)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	l, _ := strconv.Atoi(limit)
	limit_listener := netutil.LimitListener(listener, l)
	http_config := &http.Server{}

	defer limit_listener.Close()
	err = http_config.Serve(limit_listener)
	if err != nil {
		log.Fatalln(err)
	}
}

func Server() {
	var port string
	var limit string
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&port, "port", "8080", "listen port")
	f.StringVar(&limit, "limit", "10", "concurrent limit")
	f.Parse(os.Args[1:])
	for 0 < f.NArg() {
		f.Parse(f.Args()[1:])
	}
	fmt.Println(port)
	fmt.Println(limit)
	//Default(port)
	Limited(port, limit)
}

func main() {
	Server()
}
