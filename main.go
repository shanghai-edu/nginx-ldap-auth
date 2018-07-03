package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shanghai-edu/nginx-ldap-auth/g"
	"github.com/shanghai-edu/nginx-ldap-auth/http"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	go http.Start()

	select {}
}
