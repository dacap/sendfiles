// Copyright (C) 2017 David Capello

package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func fileSha1(fn string) string {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
