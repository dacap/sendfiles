# Copyright (C) 2017 David Capello

all: sendfiles

SENDFILES_SRC = sendfiles.go receivefiles.go main.go sha1.go ip.go

sendfiles: $(SENDFILES_SRC)
	go build -o $@ $^

clean:
	-rm -f sendfiles.exe sendfiles windows/sendfiles.exe macos/sendfiles linux/sendfiles
	-rmdir windows macos linux

cross:
	env GOOS=windows GOARCH=amd64 go build -v -o windows/sendfiles.exe $(SENDFILES_SRC)
	env GOOS=darwin GOARCH=amd64 go build -v -o macos/sendfiles $(SENDFILES_SRC)
	env GOOS=linux GOARCH=amd64 go build -v -o linux/sendfiles $(SENDFILES_SRC)

package:
	cd windows && zip ../sendfiles-windows.zip sendfiles.exe && cd ..
	cd macos && zip ../sendfiles-macosx.zip sendfiles && cd ..
	cd linux && zip ../sendfiles-linux.zip sendfiles && cd ..
