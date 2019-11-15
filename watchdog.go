package main

import (
	"fmt"
	"github.com/mksmsrgnk/serviceutils"
	"github.com/mksmsrgnk/smsutils"
	"log"
	"os"
)

var (
	hostName, url, workerService, from, to string
)

func main() {
	if err := serviceutils.CheckURL(url); err != nil {
		if err := serviceutils.RestartService(workerService); err != nil {
			if err := smsutils.Send(from, to, fmt.Sprintf("%v", err)); err != nil {
				log.Printf("sms error: %v", err)
			}
			log.Fatalf("%v", err)
		} else {
			message := fmt.Sprintf("%s on %s restarted successfully!", workerService, hostName)
			if err := smsutils.Send(from, to, message); err != nil {
				log.Printf("sms error: %v", err)
			}
			log.Printf("%s", message)
			return
		}
	}
	log.Printf("%s on %s is OK", workerService, hostName)
}

func init() {
	name, err := os.Hostname();
	if err != nil {
		log.Printf("can't get host name, error %v", err)
	}
	hostName = name
	url = os.Getenv("SERVICE_URL")
	workerService = os.Getenv("WORKER_SERVICE")
	smsutils.Address = os.Getenv("KANNEL_ADDRESS")
	smsutils.UserName = os.Getenv("KANNEL_USERNAME")
	smsutils.Password = os.Getenv("KANNEL_PASSWORD")
	from = os.Getenv("SMS_FROM")
	to = os.Getenv("SMS_TO")
}
