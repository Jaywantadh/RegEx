package main

import (
	"fmt"
	"log"
	"net"
	"net/smtp"
	"regexp"
	"strings"
)

func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func isDomainValid(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]
	_, err := net.LookupMX(domain)
	return err == nil
}

func sendVerificationEmail(email, token string) error {
	from := "Your_mail"
	password := "Your_app_password"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	to := []string{email}
	message := []byte(fmt.Sprintf("Subject: Email Verification\n\nPlease verify your email by clicking the following link: http://yourdomain.com/verify?token=%s", token))

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Print("Enter Your Email: ")  
    var email string 
    fmt.Scanln(&email)
	if isValidEmail(email) && isDomainValid(email) {
		token := "unique-verification-token"
		err := sendVerificationEmail(email, token)
		if err != nil {
			log.Fatalf("Failed to send verification email: %v", err)
		} else {
			fmt.Println("Verification email sent successfully!")
		}
	} else {
		fmt.Println("Invalid email address.")
	}
}
