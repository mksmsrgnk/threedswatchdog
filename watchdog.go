package main

import (
	"fmt"
	"github.com/mksmsrgnk/serviceutils"
	"github.com/mksmsrgnk/smsutils"
	"log"
	"os"
	"strings"
)

var (
	hostName, urlToCheck, workerService,
	from, to, userName, password, kannelURL string
)

func sendSMS(text string) {
	for _, t := range strings.Split(to, ",") {
		err := smsutils.NewKannel(userName, password, kannelURL).
			NewTextMessage(from, t, text).
			Send()
		if err != nil {
			log.Printf("error: %v", err)
		}
	}
}

func main() {
	if err := serviceutils.CheckURL(urlToCheck); err != nil {
		if err := serviceutils.RestartService(workerService); err != nil {
			sendSMS(fmt.Sprintf("%v", err))
			log.Fatalf("%v", err)
		}
		text := fmt.Sprintf("%s on %s restarted successfully!",
			workerService, hostName)
		sendSMS(text)
		log.Printf("%s", text)
		return
	}
	log.Printf("%s on %s is OK", workerService, hostName)
}

func init() {
	name, err := os.Hostname()
	if err != nil {
		log.Printf("can't get host name, error %v", err)
	}
	hostName = name
	urlToCheck = os.Getenv("SERVICE_URL")
	workerService = os.Getenv("WORKER_SERVICE")
	userName = os.Getenv("KANNEL_USERNAME")
	password = os.Getenv("KANNEL_PASSWORD")
	kannelURL = os.Getenv("KANNEL_URL")
	from = os.Getenv("SMS_FROM")
	to = os.Getenv("SMS_TO")
}
