package mailing

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func SendMail(to []string,subject,message string) error{
	from := os.Getenv("mail_sender")
	body := "From : "+from+"\n"+
		"To: "+ strings.Join(to,",")+"\n"+
		"Subject: "+subject+"\n\n"+
		message

	fmt.Println(from)
	fmt.Println(os.Getenv("smtp_host"))

	auth := smtp.PlainAuth("",from,os.Getenv("mail_password"),os.Getenv("smtp_host"))
	smtpAddress := fmt.Sprintf("%s:%s",os.Getenv("smtp_host"),os.Getenv("smtp_port"))
	fmt.Println(smtpAddress)
	err := smtp.SendMail(smtpAddress,auth,from,append(to),[]byte(body))
	if err != nil {
		return err
	}

	return nil
}
