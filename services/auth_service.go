package services

import (
	"errors"
	"time"
	"w2s-backend/database"
	"w2s-backend/models"
	"w2s-backend/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func RegisterUser(name, email string) error {
	if err := database.FindOne("users", bson.M{"email": email}, nil).Err(); err == nil {
		return errors.New("email already exists")
	}
	otp := utils.GenerateOTP()
	// if err := utils.SendOTPEmail(email, otp); err != nil {
	// 	return errors.New("failed to send OTP")
	// }
	user := models.User{
		UserID:   utils.NewUUID(),
		Name:     name,
		Email:    email,
		Password: "",
		OTPData: models.OTPData{
			OTP:       otp,
			ExpiresAt: time.Now().Add(5 * time.Minute),
		},
	}
	if _, err := database.InsertOne("users", user); err != nil {
		return err
	}
	return nil
}

func VerifyUserOTP(email, otp, password string) error {
	if !utils.VerifyOTP(email, otp) {
		return errors.New("invalid or expired OTP")
	}
	hashed, _ := utils.HashPassword(password)
	update := bson.M{"$set": bson.M{"password": hashed}}
	if _, err := database.UpdateOne("users", bson.M{"email": email}, update); err != nil {
		return errors.New("failed to update password")
	}
	return nil
}

func LoginUser(email, password string) (string, error) {
	res := database.FindOne("users", bson.M{"email": email}, nil)
	if err := res.Err(); err != nil {
		return "", errors.New("invalid credentials")
	}
	var user models.User
	if err := res.Decode(&user); err != nil {
		return "", errors.New("failed to decode user")
	}
	if user.Password == "" {
		return "", errors.New("account not verified")
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}
	token, err := utils.GenerateToken(user.UserID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return token, nil
}

func ResetPassword(email, newPassword string) error {
	hashed, _ := utils.HashPassword(newPassword)
	update := bson.M{"$set": bson.M{"password": hashed}}
	if _, err := database.UpdateOne("users", bson.M{"email": email}, update); err != nil {
		return errors.New("failed to reset password")
	}
	return nil
}
