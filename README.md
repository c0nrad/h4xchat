# h4xchat

A friend and I wanted a secure way to talk about some totally legal stuff. I wrote this so we could talk with resonable assurance that no one else would know our conversations. Well except for keyloggers/rootkits/and the NSA.

## How to use

First the server needs to be run somewhere where both clients have access to. I use a digital ocean droplet.

    > ./h4xServe -h
    Usage of ./h4xServe:
      -port="1337": Port to host server on
    > ./h4xServer

Now you can start talking!

    > ./h4xchat -h
    Usage of ./h4xchat:
      -host="localhost": Host to connect to
      -key="mysup3rs3cr3tk3y123!": The key used for encryption
      -port="1337": Port number of host to connect to
    > ./h4xchat

## Disclaimer

As of Feb 24th, it uses RC4 for encryption. But since everyone has to be in the same place of PRNG I reset it on every message. Essentially turning my encryption into a N-time pad. Which is trvially breakable.

## Contributing

Ideas? Send an email to sclarsen@mtu.edu.