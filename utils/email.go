package utils

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendOTPEmail(toEmail, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_EMAIL"))
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "🔐 Your OTP Code")

	htmlBody := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<style>
				.container {
					font-family: Arial, sans-serif;
					background-color: #f9f9f9;
					padding: 20px;
					border-radius: 8px;
					max-width: 400px;
					margin: auto;
					box-shadow: 0 2px 6px rgba(0,0,0,0.1);
				}
				.title {
					font-size: 18px;
					font-weight: bold;
					color: #333;
					margin-bottom: 10px;
				}
				.otp {
					font-size: 24px;
					font-weight: bold;
					color: #ffffff;
					background-color: #4CAF50;
					padding: 10px 20px;
					border-radius: 6px;
					display: inline-block;
					letter-spacing: 2px;
				}
				.footer {
					margin-top: 15px;
					font-size: 12px;
					color: #777;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<div class="title">Your OTP Code</div>
				<p>Please use the OTP below to complete your verification. It will expire in <b>5 minutes</b>.</p>
				<div class="otp">%s</div>
				<p class="footer">If you did not request this, please ignore this email.</p>
			</div>
		</body>
		</html>
	`, otp)

	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"),
	)

	return d.DialAndSend(m)
}
