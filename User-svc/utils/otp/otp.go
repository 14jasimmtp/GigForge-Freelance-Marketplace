package otp

import (
	"fmt"
	"math/rand"
	"net/smtp"

	"github.com/spf13/viper"
)


func SendVerificationOtp(email string) (int,error) {
	from := viper.GetString("OTP_Email")
	fmt.Println(from)
	password := viper.GetString("App_Password")
	fmt.Println(password)

	to := []string{email}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	otp := GenerateRandomOTP()
	fmt.Println("otp : ",otp)
	message := []byte(fmt.Sprintf(
		"From: %s \r\n"+
		"To: %s \r\n"+
		"Subject: OTP verification For GigForge\r\n"+
		"\r\n"+
		"Your verification OTP for GigForge is %v\r\n", from, email, otp,
	))

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err, "hi")
		return 0,err
	}
	fmt.Println("Email Sent Successfully!")
	return otp,nil
}

func SendForgotPasswordOTP(email string) {

}

func GenerateRandomOTP() int {
	otp := rand.Intn(9000) + 1000
	return otp
}
