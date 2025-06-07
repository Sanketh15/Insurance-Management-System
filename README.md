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


---

## ⚙️ Setup Instructions

### 📌 Prerequisites

Ensure the following are installed:

- [Go](https://go.dev/dl/) (v1.20 or later)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)
- [Live Server](https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer) (optional for frontend preview)

### 📥 Installation Steps


# 1️⃣ Clone the Repository
git clone https://github.com/your-username/insurance-management-system.git
cd insurance-management-system

# 2️⃣ Set Up PostgreSQL Database
# Open psql and create a new database
psql -U your_username
CREATE DATABASE insurance_db;

# 3️⃣ Configure Environment
# Either use a .env file or set values directly in code:
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=your_username
# DB_PASSWORD=your_password
# DB_NAME=insurance_db
# JWT_SECRET=your_secret_key

# 4️⃣ Install Dependencies
go mod tidy

# 5️⃣ Auto-Migrate DB Tables (inside main.go)
# db.AutoMigrate(&User{}, &Claim{})

# 6️⃣ Run the Backend Server
go run main.go

# 7️⃣ Launch the Frontend
# Option 1: Open frontend/index.html in browser
# Option 2: Serve using Python
python3 -m http.server 8080 --directory frontend
