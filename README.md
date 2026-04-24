# 📁 Go Drive Explorer (Gin + Google Drive API)

A lightweight Google Drive explorer built using **Golang (Gin)** and the **Google Drive API v3**.
This project demonstrates OAuth 2.0 authentication, persistent token handling, and file navigation similar to Google Drive.

---

## 🚀 Features

- 🔐 OAuth 2.0 authentication (Google Sign-In)
- ♻️ Persistent login using refresh tokens
- 📂 Browse Google Drive folders (recursive navigation)
- 📄 View files and folders (like Drive UI)
- ⬅️ Back navigation using stack
- ⚡ REST API built with Gin
- 🌐 Simple frontend (HTML + JS)

---

## 🧱 Tech Stack

### Backend

- Go (Golang)
- Gin Web Framework
- Google OAuth2
- Google Drive API v3

### Frontend

- HTML
- Vanilla JavaScript (Fetch API)

---

## 📁 Project Structure

```
go-drive-manager/
│
├── main.go                 # Entry point, routes, server setup
│
├── config/
│   └── oauth.go            # OAuth configuration (client ID, scopes)
│
├── handlers/
│   ├── auth.go             # Login + OAuth callback handlers
│   └── drive.go            # Drive API endpoints (list files)
│
├── services/
│   └── drive.go            # Google Drive service creation
│
├── templates/
│   └── index.html          # Frontend UI
│
├── static/
│   └── script.js           # Frontend logic (navigation)
│
├── tokens/
│   └── token.json          # Stored OAuth tokens (ignored in Git)
│
├── .env                    # Environment variables (ignored in Git)
├── .gitignore              # Git ignore rules
├── go.mod                  # Go module definition (dependencies)
├── go.sum                  # Dependency checksums (auto-generated)
└── README.md               # Project documentation
```

---

## ⚙️ Setup Instructions

### 1️⃣ Clone Repository

```bash
git clone https://github.com/latesh-munde/Google-Drive-Manager.git
cd go-drive-manager
```

---

### 2️⃣ Install Dependencies

```bash
go mod tidy
```

---

### 3️⃣ Setup Google OAuth

Go to Google Cloud Console:

1. Create a project
2. Enable **Google Drive API**
3. Create OAuth 2.0 credentials
4. Authorized JavaScript origins

```
http://localhost:8080
```

5. Add redirect URI:

```
http://localhost:8080/auth/callback
```

---

### 4️⃣ Configure Environment Variables

Create `.env` file:

```
GOOGLE_CLIENT_ID=your_client_id
GOOGLE_CLIENT_SECRET=your_client_secret
GOOGLE_REDIRECT_URI=http://localhost:8080/auth/callback
```

---

### 5️⃣ Run the Server

```bash
go run main.go
```

Server will start at:

```
http://localhost:8080
```

---

## 🔐 Authentication Flow

1. Open:

   ```
   http://localhost:8080/auth/login
   ```

2. Login with Google
3. After success:
   - Token stored in `tokens/token.json`
   - No need to login again (auto refresh)

---

## 🌐 Using the App

### Open UI

```
http://localhost:8080/
```

### Features

- 📂 Click folder → Navigate inside
- ⬅️ Back button → Go to previous folder
- 📄 Files displayed with icons

---

## 🔌 API Endpoints

### 🔐 Auth

| Method | Endpoint       | Description            |
| ------ | -------------- | ---------------------- |
| GET    | /auth/login    | Start OAuth flow       |
| GET    | /auth/callback | Handle Google callback |

---

### 📂 Drive

| Method | Endpoint                   | Description          |
| ------ | -------------------------- | -------------------- |
| GET    | /drive/files?folderId=<id> | List files in folder |

---

## 🔄 How It Works

1. User logs in via Google OAuth
2. Backend exchanges code → tokens
3. Tokens stored locally
4. Frontend requests `/drive/files`
5. Backend calls Google Drive API
6. Data returned and rendered

---

## 🔐 Token Handling

- Access token expires (~1 hour)
- Refresh token is stored
- Go OAuth client auto-refreshes tokens
- No repeated login required

---

## ⚠️ Security Notes

- Do NOT commit `.env`
- Do NOT expose `CLIENT_SECRET`
- Rotate credentials if leaked
- Use secure storage (DB) in production

---

## 🚀 Future Improvements

- ✏️ Rename files
- 📦 Move files between folders
- 📑 Copy files
- 🗑 Delete files
- 🔍 Search functionality
- 👥 Multi-user authentication
- 🧠 Clean architecture refactor
- 🎨 Better UI (React / Tailwind)

---

## 🎯 Learning Outcomes

This project demonstrates:

- OAuth 2.0 flow in Go
- REST API design with Gin
- External API integration
- Token persistence strategy
- Basic frontend-backend integration
- Folder navigation logic

---

## 📌 Author

Built as a backend-focused project using Go + Gin to simulate a simplified Google Drive experience.

---
