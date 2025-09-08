# w2s Project API

This is a backend service built with **Go (Gin + MongoDB)** for user authentication and profile management.

## 🚀 Features
- User Registration with OTP verification (via Email)
- Secure Login with JWT
- Forgot / Reset Password (with OTP)
- User Profile Management (view, update profile, update email)

## 📂 API Documentation
- Importable collection: [API Collection](./API_Collections.yaml)
- Swagger UI: [https://w2s-backend.onrender.com/swagger/index.html](https://w2s-backend.onrender.com/swagger/index.html)  
- Swagger Spec: [swagger.json](./docs/swagger.json)

## 🛠️ Tech Stack
- Go (Gin Web Framework)
- MongoDB (User & OTP storage)
- Gomail (SMTP for OTP emails)
- Swagger (API Docs)

## ▶️ Running the Project
1. Start MongoDB
2. Configure environment variables (create a `.env` file in the project root):

   ```env
   MONGO_URI=mongodb://localhost:27017
   DB_NAME=w2s
   JWT_SECRET=your-secret-key
   PORT=8080

   SMTP_EMAIL=your-email@gmail.com
   SMTP_PASSWORD=your-app-password
