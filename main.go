// Copyright (C) 2017-2024 David Capello
//
// This file is released under the terms of the MIT license.
// Read LICENSE.txt for more information.

package main

import (
	"flag"
)

var key string
var ip string
var port int = DefaultPort

func main() {
	flag.StringVar(&key, "k", "", "key to receive files")
	flag.StringVar(&ip, "ip", "", "IP of the receiver")
	flag.IntVar(&port, "p", port, "TCP port")
	flag.Parse()

	if flag.NArg() < 1 {
		receiveFiles()
	} else {
		sendFiles(flag.Args(), ip)
	}

}
