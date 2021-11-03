package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mailgun/mailgun-go/v4"
	"log"
	"time"
)

type Mail struct {
	Sender    string `json:"sender"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Recipient string `json:"recipient"`
}

var yourDomain string = "sandboxa00df2f07efd4db1bb46391f42301722.mailgun.org"

var privateAPIKey string = "bc247506eff73365cfaea17c71d4fad7-20ebde82-86e2cb60"

func handler(ctx context.Context, snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		mail := Mail{}
		err := json.Unmarshal([]byte(snsRecord.Message),&mail)
		if err != nil {
			log.Println(err)
		}

		_,err = sendMail(mail)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	lambda.Start(handler)
}

func sendMail(mail Mail) (string, error) {
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(mail.Sender, mail.Subject, mail.Body, mail.Recipient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	return fmt.Sprintf("ID: %s Resp: %s\n", id, resp), err
}
