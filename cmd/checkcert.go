package cmd

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)



// CheckCertCommand clean command struct
type CheckCertCommand struct {
	BaseCommand
}

// Init CheckCertCommand
func (ccc *CheckCertCommand) Init() {
	ccc.command = &cobra.Command{
		Use:   "checkcert",
		Short: "Clear lapsed context, cluster and user",
		Long:  "Clear lapsed context, cluster and user",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("checkcert start！")
			start()
			return nil
		},
	}

}


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

func start() {
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
