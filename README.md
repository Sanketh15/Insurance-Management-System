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


---

## ⚙️ Setup Instructions

### 📌 Prerequisites

Ensure the following are installed:

- [Go](https://go.dev/dl/) (v1.20 or later)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)
- [Live Server](https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer) (optional for frontend)

### 📥 Installation Steps

```bash
# 1️⃣ Clone the Repository
git clone https://github.com/your-username/insurance-management-system.git
cd insurance-management-system

# 2️⃣ Set Up PostgreSQL Database
# Open psql and create a new database:
psql -U your_username
CREATE DATABASE insurance_db;

# 3️⃣ Configure Environment
# Either use .env file or hardcode the following:
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=your_username
# DB_PASSWORD=your_password
# DB_NAME=insurance_db
# JWT_SECRET=your_secret_key

# 4️⃣ Install Go Dependencies
go mod tidy

# 5️⃣ Auto-Migrate DB Tables (inside main.go or setup script)
# Example:
# db.AutoMigrate(&User{}, &Claim{})

# 6️⃣ Run the Backend Server
go run main.go

# 7️⃣ Launch the Frontend
# Option 1: Open frontend/index.html directly in browser
# Option 2: Serve via Python
python3 -m http.server 8080 --directory frontend

# ==========================================================
# 📡 API ENDPOINTS
# ==========================================================

# POST   /login          - Authenticate user/admin
# POST   /register       - Register a new user
# GET    /claims         - View all claims (Admin only)
# POST   /claim          - Submit a new claim
# PUT    /claim/:id      - Approve/Reject claim (Admin)

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
