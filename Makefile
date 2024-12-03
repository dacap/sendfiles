# Copyright (C) 2017 David Capello

all: sendfiles

SENDFILES_SRC = sendfiles.go receivefiles.go main.go sha1.go ip.go

sendfiles: $(SENDFILES_SRC)
	go build -o $@ $^

cross:
	env GOOS=windows GOARCH=386 go build -v -o win/sendfiles.exe $(SENDFILES_SRC)
	env GOOS=darwin GOARCH=386 go build -v -o mac/sendfiles $(SENDFILES_SRC)
	env GOOS=linux GOARCH=386 go build -v -o lin/sendfiles $(SENDFILES_SRC)

package:
	cd win && zip ../sendfiles-windows.zip sendfiles.exe && cd ..
	cd mac && zip ../sendfiles-macosx.zip sendfiles && cd ..
	cd lin && zip ../sendfiles-linux.zip sendfiles && cd ..
