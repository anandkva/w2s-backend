# w2s Project API

This is a backend service built with **Go (Gin + MongoDB)** for user authentication and profile management.

## 🚀 Features
- User Registration with OTP verification (via Email)
- Secure Login with JWT
- Forgot / Reset Password (with OTP)
- User Profile Management (view, update profile, update email)

## 📂 API Documentation
- Importable collection: [API Collection](./API_Collections.yaml)

## 🛠️ Tech Stack
- Go (Gin Web Framework)
- MongoDB (User & OTP storage)
- Gomail (SMTP for OTP emails)
- Swagger (API Docs)

## ▶️ Running the Project
1. Start MongoDB
2. Configure environment variables:
   ```bash
   export SMTP_EMAIL="your-email@gmail.com"
   export SMTP_PASSWORD="your-app-password"
   export MONGO_URI="mongodb://localhost:27017/w2s"
