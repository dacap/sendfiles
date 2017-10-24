# sendfiles

> Copyright (C) 2017 David Capello

A program to send/receive files between two personal computers.

## Usage

We run the `receivefiles` in computer *A*:

    receivefiles

And in other computer *B* we specify the files to transfer:

    sendfiles package-windows.zip package-linux.zip

This will send `package-windows.zip` and `package-linux.zip` files
from *B* to *A*. The IP of *A* is located automatically by *B* with a
scan of IP addresses in the local network and the TCP port 8095
(we can change this port with the `-p` argument).

## Key

`receivefiles` and `sendfiles` can receive a key parameter (`-k string`)
to match computers with the same key. The key is not used to encrypt data.

## Plain data/key

All information will be transferred in plain data over the network
between computers. There is zero encryption.

## Protocol

Client scan IP addresses in the local network tries to connect to the
TCP port 8095, after it's connected it sends one line:

    key STRING\n

Server responds with:

    ok\n

when the key is accepted and the client can start sending files, or

    invalid key\n

if the connection will be terminate.

When the key is accepted, the client start sending files:

    file FILENAME size INT64 sha1 STRING\n
    BYTES[size]\n
    file FILENAME size INT64 sha1 STRING\n
    BYTES[size]\n
    ...
    done\n