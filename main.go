package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	alerts "github.com/opsgenie/opsgenie-go-sdk/alertsv2"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
)

func triggerAlert(domain string) {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(os.Getenv("OPSGENIE_API_KEY"))

	alertCli, _ := cli.AlertV2()

	request := alerts.CreateAlertRequest{
		Message:     "SSL certificate expired",
		Description: fmt.Sprintf("Certificate of %s expired", domain),
		Priority:    alerts.P1,
	}

	_, err := alertCli.Create(request)
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	domains := strings.Split(os.Getenv("DOMAINS"), ",")

	for {
		for i := range domains {
			site := fmt.Sprintf("%s:443", domains[i])
			log.Println("Checking", site)

			conn, err := tls.Dial("tcp", site, nil)
			if err != nil {
				log.Println(err)
				triggerAlert(domains[i])
				continue
			}

			err = conn.VerifyHostname(domains[i])
			if err != nil {
				log.Println(err)
				continue
			}
		}
		log.Println("Sleeping for next 24 hours...")
		time.Sleep(24 * time.Hour)
	}
}
