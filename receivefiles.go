// Copyright (C) 2017 David Capello

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func writeFile(r io.Reader, fn string, size int64) {
	f, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := io.CopyN(f, r, size); err != nil {
		log.Fatal(err)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Receiving files from", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	// First line must contain the key
	var receivedKey string
	msg, _ := reader.ReadString('\n')
	fmt.Sscanf(msg, "key %s", &receivedKey)

	// Check that the key is valid
	if key != "" && receivedKey != key {
		fmt.Printf("Received key '%s' is not the expected key '%s'\n",
			receivedKey, key)
		fmt.Fprintf(conn, "invalid key\n")
		time.Sleep(5 * time.Second)
		return
	}

	// Send "ok" to the client (so it starts sending files)
	fmt.Fprintf(conn, "ok\n")

	for {
		var fn, sha1 string
		var size int64
		msg, _ := reader.ReadString('\n')

		if strings.Index(msg, "done") == 0 {
			break
		}

		fmt.Sscanf(msg, "file %s size %d sha1 %s", &fn, &size, &sha1)
		if msg == "" || fn == "" {
			break
		}

		fmt.Printf("Receiving file %s (%d bytes)...", fn, size)

		localFn := strings.Replace(fn, "/", "-", -1)
		localFn = "_received_" + localFn
		writeFile(reader, localFn, size)

		localSha1 := fileSha1(localFn)
		if sha1 == localSha1 {
			fmt.Printf(" OK\n")
		} else {
			fmt.Printf(" checksum FAILED\n")
			fmt.Printf("   local  SHA1: %s\n", localSha1)
			fmt.Printf("   remote SHA1: %s\n", sha1)
		}
	}

	fmt.Println("Closing connection to", conn.RemoteAddr())
}

func waitClients() {
	ln, err := net.Listen("tcp", ":" + strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func receiveFiles() {
	fmt.Println("Waiting for files...")
	if key != "" {
		fmt.Printf("(Use key '%s' to receive files)\n", key)
	}

	ips := getIpAddresses(make([]net.IP, 0))
	for _, ip := range ips {
		ip4 := ip.To4()
		if ip4 != nil {
			fmt.Println(" IP ", ip4.String())
		}
	}

	go waitClients()

	// Wait for Enter key
	bufio.NewReader(os.Stdin).ReadString('\n')
}
