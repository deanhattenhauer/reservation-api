package main

import (
	"log"
	"log/slog"

	"github.com/resend/resend-go/v3"
)

func SendEmail(apiKey string, to string, subject string, html string) error {

    client := resend.NewClient(apiKey)

    params := &resend.SendEmailRequest{
        From:    "onboarding@resend.dev",
        To:      []string{to},
        Subject: subject,
        Html:    html,
    }

    sent, err := client.Emails.Send(params)
    if err != nil {
		slog.Error("failed to send email", "error", err)
		return err
	}

    log.Printf("email sent: %s", sent.Id)
	return nil
}