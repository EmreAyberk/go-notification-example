package main

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

var yourDomain string = "sandboxa00df2f07efd4db1bb46391f42301722.mailgun.org"

var privateAPIKey string = "bc247506eff73365cfaea17c71d4fad7-20ebde82-86e2cb60"

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "send",
				Aliases: []string{
					"s",
				},
				Action: func(c *cli.Context) error {
					fmt.Println("girdimmm")
					mail, err := sendMail(c.String("sender"), c.String("subject"), c.String("body"), c.String("recipient"))
					if err != nil {
						return err
					}
					fmt.Println(mail)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "sender", Required: true},
					&cli.StringFlag{Name: "subject", Required: true},
					&cli.StringFlag{Name: "body", Required: true},
					&cli.StringFlag{Name: "recipient", Required: true},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func sendMail(sender, subject, body, recipient string) (string, error) {
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	return fmt.Sprintf("ID: %s Resp: %s\n", id, resp), err
}
