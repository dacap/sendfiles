// Copyright (C) 2017 David Capello

package main

import (
	"flag"
)

var key string
var port int = DefaultPort

func main() {
	flag.StringVar(&key, "k", "", "key to receive files")
	flag.IntVar(&port, "p", port, "TCP port")
	flag.Parse()

	if flag.NArg() < 1 {
		receiveFiles()
	} else {
		sendFiles(flag.Args())
	}

}
