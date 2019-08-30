package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

var (
	errSuccess       = fmt.Errorf("success")
	errFailure       = fmt.Errorf("failure")
	handshakeSuccess = fmt.Sprintf("ssh: handshake failed: %s", errSuccess)
	handshakeFailure = fmt.Sprintf("ssh: handshake failed: %s", errFailure)
)

func main() {
	rawFlag := flag.Bool("raw", false, "raw output (can pipe to `ssh-keygen -L`)")
	flag.Parse()

	cc := ssh.ClientConfig{
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			if cert, ok := key.(*ssh.Certificate); ok {
				// TODO validate that the cert is signed by a trusted CA
				if *rawFlag {
					fmt.Printf("%s %s\n", cert.Type(), base64.StdEncoding.EncodeToString(cert.Marshal()))
				} else {
					encoder := json.NewEncoder(os.Stdout)
					err := encoder.Encode(cert)
					if err != nil {
						return err
					}
				}

				return errSuccess
			}

			fmt.Fprintf(os.Stderr, "No certificate returned by host, only: %T\n", key)
			return errFailure
		},
	}

	address := flag.Arg(0)
	if !strings.ContainsRune(address, ':') {
		address = fmt.Sprintf("%s:22", address)
	}

	c, err := ssh.Dial("tcp", address, &cc)
	if err == nil {
		c.Close()
		log.Fatalf("Unexpected condition: ssh.Dial returned no error")
	}

	switch err.Error() {
	case handshakeSuccess:
	case handshakeFailure:
		os.Exit(1)
	default:
		log.Fatalf("%T %v", err, err)
	}
}
