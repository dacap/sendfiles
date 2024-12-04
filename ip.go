// Copyright (C) 2017 David Capello
//
// This file is released under the terms of the MIT license.
// Read LICENSE.txt for more information.

package main

import (
	"net"
)

var DefaultPort = 8095

func getIpAddresses(ips []net.IP) []net.IP {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ips
	}
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			ips = append(ips, ip)
		}
	}
	return ips
}
