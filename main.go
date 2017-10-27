// Copyright (C) 2017 David Capello

package main

import (
	"flag"
	"fmt"
)

var key string
var port int = DefaultPort

func main() {
	flag.StringVar(&key, "k", "", "key to receive files")
	flag.IntVar(&port, "p", port, "TCP port")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("You must specify the command send|receive")
		return
	}

	switch flag.Arg(0) {
	case "send":
		sendFiles(flag.Args()[1:])
	case "receive":
		receiveFiles()
	}

}
