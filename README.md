# 🛡️ Insurance Management System

A web-based application that streamlines the insurance claim process for employees and administrators. Developed using **Golang** for backend logic, **GORM** for ORM, and **PostgreSQL** for data storage, the system provides a secure, efficient, and scalable platform for managing insurance claims.

## 🚀 Features

- 🧾 **Claim Submission**: Employees can submit insurance claims online with necessary details.
- 🔐 **Role-Based Authentication**: Secure login for both users and admins using JWT tokens.
- 📊 **Admin Dashboard**: Admins can view, approve, or reject claims and manage salary deductions.
- 📥 **Document Upload**: Supports uploading documents for verification.
- 🔄 **Real-Time Status Tracking**: Employees can track their claim status (Pending/Approved/Rejected).
- 📈 **Optimized Performance**: Uses GORM and efficient SQL queries for fast processing.

## 🛠️ Tech Stack

| Layer          | Technology           |
|----------------|----------------------|
| Backend        | Golang (Go)          |
| ORM            | GORM                 |
| Database       | PostgreSQL           |
| Frontend       | HTML, CSS, JavaScript|
| Auth           | JWT-based login      |
| Tools          | Postman, VS Code     |
| Version Control| Git & GitHub         |

## 📁 Folder Structure

/insurance-management-system
│
├── /backend
│   ├── /api             # Route handlers (e.g., /claim, /auth)
│   ├── /controllers     # Business logic for API routes
│   ├── /models          # GORM models for database tables
│   ├── /middlewares     # JWT auth, role-based access, logging
│   ├── /config          # DB connection and environment configs
│   ├── main.go          # Application entry point
│   └── go.mod           # Go module definition
│
├── /frontend
│   ├── index.html       # Main UI page (Login/Landing)
│   ├── /css             # Stylesheets
│   ├── /js              # Scripts for form handling, fetch API
│   └── /assets          # Logos, icons, and image files
│
├── /screenshots         # UI snapshots for README/docs
│
└── README.md            # Project documentation

## ⚙️ Setup Instructions

### Prerequisites

- Go (v1.20+)
- PostgreSQL
- Git

### Steps


# 1️⃣ PREREQUISITES:
# Make sure the following are installed:
# - Go (v1.20 or later): https://go.dev/dl/
# - PostgreSQL: https://www.postgresql.org/download/
# - Git: https://git-scm.com/
# - Live Server (for frontend, optional): https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer

# 2️⃣ CLONE THE REPOSITORY:
git clone https://github.com/your-username/insurance-management-system.git
cd insurance-management-system

# 3️⃣ SETUP POSTGRESQL DATABASE:
# Create a database named `insurance_db` (or your custom name).
# Example using psql:
# CREATE DATABASE insurance_db;

# 4️⃣ CONFIGURE ENVIRONMENT:
# If using a `.env` file or hardcoded config, set the following values:
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=your_username
# DB_PASSWORD=your_password
# DB_NAME=insurance_db
# JWT_SECRET=your_secret_key

# 5️⃣ INSTALL DEPENDENCIES (Golang modules):
go mod tidy

# 6️⃣ AUTO-MIGRATE DATABASE TABLES (inside main.go or setup script):
# db.AutoMigrate(&User{}, &Claim{})

# 7️⃣ RUN THE BACKEND SERVER:
go run main.go

# 8️⃣ START THE FRONTEND:
# Open frontend/index.html in browser
# OR use Live Server / Python server:
# python3 -m http.server 8080 --directory frontend

# ==========================================================
# 🧠 LEARNING OUTCOMES
# ==========================================================

# - Developed full-stack web app using Go and PostgreSQL
# - Implemented secure JWT-based authentication and role control
# - Built RESTful APIs with GORM ORM
# - Learned to optimize database queries and handle user sessions

# ==========================================================
# 📄 LICENSE
# ==========================================================

# This project is licensed under the MIT License.

# ==========================================================
# 🙌 ACKNOWLEDGEMENTS
# ==========================================================

# - https://go.dev/doc/
# - https://gorm.io/docs/
# - https://www.postgresql.org/docs/
# - https://github.com/dgrijalva/jwt-go
# ==========================================================
