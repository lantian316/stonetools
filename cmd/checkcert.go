package cmd

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"time"
)

func checkCert(domain string) {
	conn, err := tls.Dial("tcp", domain+":443", &tls.Config{})
	if err != nil {
		log.Println(domain, ":", err)
		return
	}
	defer conn.Close()

	state := conn.ConnectionState()
	for _, v := range state.PeerCertificates {
		fmt.Println(domain, ":", v.NotAfter.Format("2006-01-02"))
		break
	}
}

func main() {
	file, err := os.Open("domains.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	limiter := make(chan struct{}, 5)
	timeout := make(chan bool, 1)

	for scanner.Scan() {
		domain := scanner.Text()
		limiter <- struct{}{}

		go func() {
			defer func() {
				<-limiter
			}()

			go func() {
				time.Sleep(3 * time.Second)
				timeout <- true
			}()

			checkCert(domain)
		}()

		select {
		case <-timeout:
			log.Println(domain, ": timeout")
		default:
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
