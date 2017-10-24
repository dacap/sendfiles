# Copyright (C) 2017 David Capello

all: bin/sendfiles bin/receivefiles

bin/sendfiles: sendfiles.go sha1.go ip.go
	go build -o $@ $^

bin/receivefiles: receivefiles.go sha1.go ip.go
	go build -o $@ $^
