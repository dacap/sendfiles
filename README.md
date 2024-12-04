# sendfiles

> Copyright (C) 2017-2024 David Capello
>
> This project is released under the terms of the MIT license.
> Read [LICENSE.txt](LICENSE.txt) for more information.

A program to send/receive files between two personal computers.

## Usage

We run the `sendfiles` in computer *A* to receive files:

    sendfiles

And in other computer *B* we specify the files to transfer:

    sendfiles file1.zip file2.zip

This will send `file1.zip` and `file2.zip` files from *B* to *A*. The
IP of *A* is located automatically by *B* with a scan of IP addresses
in the local network and the TCP port 8095 (we can change this port
with the `-p` argument).

## Key

`sendfiles` can receive a key parameter (`-k string`) to match
computers with the same key. The key is not used to encrypt data.

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
