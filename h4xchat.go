package main

import (
	"bufio"
	"crypto/rc4"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1337")
	if err != nil {
		panic(err)
	}

	key := []byte("a very very very very secret key")

	go readMessages(conn, key)
	writeMessages(conn, key)

}

func readMessages(conn net.Conn, key []byte) {
	reader := bufio.NewReader(conn)
	for {
		line, _, err := reader.ReadLine()
		fmt.Println(string(line))
		clear := rc4XOR(key, line)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(clear))
	}
}

func writeMessages(conn net.Conn, key []byte) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()

		cipher := rc4XOR(key, ([]byte)(line))
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(conn, string(cipher)+"\n")
	}
}

// Super bad I know. Essentially a OTP. Maybe switch to TDES? RSA?
func rc4XOR(key, text []byte) []byte {
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		panic(err)
	}
	cipher.XORKeyStream(text, text)

	return text
}
