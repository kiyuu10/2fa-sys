package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/spf13/viper"
)

// Generate OTP 6 digits
func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// HashOTP generate from OTP
func HashOTP(otp string) string {
	hash := sha256.Sum256([]byte(otp))
	return hex.EncodeToString(hash[:])
}

// SendEmail sends through email
func SendEmail(email string, otp string) error {
	api := viper.GetString("email.api_key")
	secret := viper.GetString("email.pri_key")
	mailjetClient := mailjet.NewMailjetClient(api, secret)
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "stuser1298@gmail.com",
				Name:  "Cookie",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: email,
					Name:  "user",
				},
			},
			Subject:  "Your OTP code",
			TextPart: "Dear passenger 1, welcome to our system! this is your otp code: " + otp,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
	return nil
}
