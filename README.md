# Snippetbox 📦

A web application for sharing snippets of text, built in **Go**.  
This project was created as a learning exercise to explore modern web development in Go.

---

## 🚀 Features
- User-friendly web interface for storing and viewing text snippets  
- Clean separation of concerns using Go’s `net/http` standard library  
- Secure cookie-based session management  
- Form validation and error handling  
- Middleware for logging, security headers, and panic recovery  
- Follows best practices for structuring Go web applications  

---

## 📂 Project Structure
```
.
├── cmd/web          # Application entry point
├── ui/              # HTML templates, static files (CSS, JS, images)
├── internal/        # Application-specific packages
│   ├── models       # Data models and database logic
│   ├── validator    # Form validation helpers
│   └── ...          
├── go.mod           
└── README.md
```

---

## ⚡ Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/<your-username>/snippetbox.git
cd snippetbox
```

### 2. Install dependencies
Go modules are used for dependency management:
```bash
go mod tidy
```

### 3. Set up the database
This app uses **Turso** (SQLite cloud). Create a database and update your DSN (Data Source Name) in the app configuration.

Example connection string:
```
libsql://<your-database-name>-<org>.turso.io?authToken=<your-auth-token>
```

### 4. Run the app
```bash
go run ./cmd/web
```

The app will start on **http://localhost:4000**.

---

## 🌍 Deployment
This app is deployed on **Railway**.  
You can easily deploy it by connecting your repository to [Railway](https://railway.app/) and configuring the required environment variables (like the Turso connection string).

---

## 🛠 Tech Stack
- **Language**: Go  
- **Database**: Turso (SQLite cloud)  
- **Templates**: HTML + Go’s `html/template` package  
- **Sessions**: Secure cookies  
- **Deployment**: Railway  

---

## 📖 References
- [Go Documentation](https://golang.org/doc/)  
- [Turso Documentation](https://docs.turso.tech/)  
- [Railway Documentation](https://docs.railway.app/)  

---

## 🤝 Contributing
This is primarily a learning project, but feel free to fork and experiment.

