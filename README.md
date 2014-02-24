# H4xchat

A friend and I wanted a secure way to talk about some totally legal stuff. I wrote this so we could talk with resonable assurance that no one else would know our conversations. Well except for keyloggers/rootkits/and the NSA.

## How to use

First the server needs to be run somewhere where both clients have access to. I use a digital ocean droplet.

    > go run h4xServe.go

Now you can start talking!

    > go run h4xchat.go


## Code Overview 

Two components. The first is the server, and the client.

### Server

The server is a simple echo server. It works a lot like socket.io's broadcast. If it recieves text on one socket, is forwards it to all other sockets.

### Client

Most of the magic happens on the client 