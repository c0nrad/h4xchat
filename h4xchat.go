package main

import (
	"bufio"
	"crypto/rc4"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var host, port, key string
	flag.StringVar(&host, "host", "localhost", "Host to connect to")
	flag.StringVar(&port, "port", "1337", "Port number of host to connect to")
	flag.StringVar(&key, "key", "mysup3rs3cr3tk3y123!", "The key used for encryption")
	flag.Parse()

	fmt.Printf("Connecting to " + host + ":" + port)
	for i := 0; i < 3; i++ {
		fmt.Printf(".")
		time.Sleep(200 * time.Millisecond)
	}
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	fmt.Println(" Connection secure.")
	keyBytes := []byte(key)

	go readMessages(conn, keyBytes)
	writeMessages(conn, keyBytes)

}

func readMessages(conn net.Conn, key []byte) {
	reader := bufio.NewReader(conn)
	for {
		line, _, err := reader.ReadLine()
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
