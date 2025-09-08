package models

import "time"

type User struct {
	UserID   string  `bson:"userId" json:"userId"`
	Name     string  `bson:"name" json:"name"`
	Email    string  `bson:"email" json:"email"`
	Password string  `bson:"password" json:"-"`
	OTPData  OTPData `bson:"otpData" json:"-"`
}
type OTPData struct {
	OTP       string    `bson:"otp"`
	ExpiresAt time.Time `bson:"expiresAt"`
}
