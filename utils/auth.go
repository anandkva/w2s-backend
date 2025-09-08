package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
	"w2s-backend/database"
	"w2s-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func GenerateOTP() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return fmt.Sprintf("%06d", n.Int64())
}

func SaveOTP(email, otp string) error {
	otpData := models.OTPData{
		OTP:       otp,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	update := bson.M{
		"$set": bson.M{
			"otpData": otpData,
		},
	}

	_, err := database.UpdateOne("users", bson.M{"email": email}, update)
	return err
}

func VerifyOTP(email, otp string) bool {
	res := database.FindOne("users", bson.M{"email": email}, nil)
	if res.Err() != nil {
		return false
	}

	var user struct {
		OTPData models.OTPData `bson:"otpData"`
	}
	if err := res.Decode(&user); err != nil {
		return false
	}

	if user.OTPData.OTP == otp && time.Now().Before(user.OTPData.ExpiresAt) {
		_, _ = database.UpdateOne("users", bson.M{"email": email}, bson.M{
			"$unset": bson.M{"otpData": ""},
		})
		return true
	}

	return false
}
