// Copyright (C) 2017 David Capello
//
// This file is released under the terms of the MIT license.
// Read LICENSE.txt for more information.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func sendFile(conn net.Conn, fn string) {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := io.Copy(conn, f); err != nil {
		log.Fatal(err)
	}
}

func connectToServer(ip net.IP, files []string) {
	conn, err := net.Dial("tcp", ip.String() + ":" + strconv.Itoa(port))
	if err != nil {
		return
	}

	if conn != nil {
		fmt.Println("Connected to " + ip.String())

		// Send key (it can be an empty string)
		fmt.Fprintf(conn, "key %s\n", key)

		// Check if we've received "ok" from the server
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')

		if strings.Index(msg, "ok") != 0 {
			if key != "" {
				fmt.Printf("Invalid key '%s' for %s\n",
					key, ip.String())
			} else {
				fmt.Println("Server requires a key")
			}
			return
		}

		for _, fn := range files {
			fmt.Printf("Sending %s to %s...", fn, ip.String())

			fi, err := os.Stat(fn)
			if err != nil {
				log.Fatal("Error getting file size from", fn)
				continue;
			}

			_, fnOnly := filepath.Split(fn)
			fmt.Fprintf(conn,
				"file %s size %d sha1 %s\n",
				fnOnly, fi.Size(), fileSha1(fn))

			sendFile(conn, fn)
			fmt.Printf(" OK\n")
		}

		fmt.Fprintf(conn, "done\n")
		time.Sleep(5 * time.Second)
		conn.Close()
	}
}

func scanIps(files []string) {
	ips := getIpAddresses(make([]net.IP, 0))
	for _, ip := range ips {
		ip4 := ip.To4()
		// Check if the IP4 address is from a private network
		if ip4 != nil &&
			((ip4[0] == 10) ||
			 (ip4[0] == 172 && (ip4[1] >= 16 && ip4[1] <= 31)) ||
			 (ip4[0] == 192 && ip4[1] == 168)) {
			for i := 0; i <= 255; i++ {
				serverIp4 := net.IPv4(ip4[0], ip4[1], ip4[2], byte(i))
				go connectToServer(serverIp4, files)
			}
		}
	}
}

func sendFiles(files []string) {
	if len(files) < 1 {
		fmt.Println("You must specify at least one file to send")
		return
	}

	fmt.Println("Searching for receiver...")

	go scanIps(files)

	// Wait for Enter key
	bufio.NewReader(os.Stdin).ReadString('\n')
}
