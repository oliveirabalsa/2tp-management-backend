# 2TP Management Backend

## 📌 Overview

2TP Management Backend is a simple Trello-like task management system built with **Golang**, **Gin Gonic**, and **GORM (SQLite)**. It supports user authentication, board creation, task management, and role-based access.

## 🚀 Features

- **User Authentication** (Signup, Login, Logout with JWT)
- **Board Management** (Create, View, Assign Users)
- **Column Management** (Create, View, Delete Columns within a Board)
- **Task Management** (Create, Assign, View, Delete Tasks within Columns)
- **Role-Based Access Control (RBAC)** (Admin/User)
- **SQLite Database for Local Persistence**

## 🛠️ Tech Stack

- **Backend:** Golang, Gin Gonic, GORM
- **Database:** SQLite
- **Authentication:** JWT

---

## 📌 Installation & Setup

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/yourusername/2tp-management-backend.git
cd 2tp-management-backend
```

### 2️⃣ Install Dependencies

```bash
go mod tidy
```

### 3️⃣ Setup Environment Variables

Create a `.env` file in the project root and define:

```env
JWT_SECRET=your-secret-key
```

### 4️⃣ Run the Server

```bash
go run main.go
```

---

## 📌 API Endpoints & Usage

### **🔐 User Authentication**

#### **1️⃣ Signup**

```http
POST /api/signup
```

**Request Body:**

```json
{
  "username": "testuser",
  "password": "password123"
}
```

**Response:**

```json
{
  "message": "User registered successfully"
}
```

#### **2️⃣ Login**

```http
POST /api/login
```

**Request Body:**

```json
{
  "username": "testuser",
  "password": "password123"
}
```

**Response:**

```json
{
  "token": "your-jwt-token"
}
```

#### **3️⃣ Logout** (Protected)

```http
POST /api/logout
```

**Headers:**

```http
Authorization: Bearer your-jwt-token
```

**Response:**

```json
{
  "message": "Successfully logged out"
}
```

---

### **📌 Board Management**

#### **1️⃣ Create a Board** (Protected)

```http
POST /api/boards
```

**Request Body:**

```json
{
  "title": "Project Board"
}
```

**Response:**

```json
{
  "id": 1,
  "title": "Project Board"
}
```

#### **2️⃣ Get All Boards** (Protected)

```http
GET /api/boards
```

**Response:**

```json
[
  {
    "id": 1,
    "title": "Project Board"
  }
]
```

---

### **📌 Column Management**

#### **1️⃣ Create a Column** (Protected)

```http
POST /api/columns
```

**Request Body:**

```json
{
  "title": "To Do",
  "board_id": 1
}
```

**Response:**

```json
{
  "id": 1,
  "title": "To Do"
}
```

#### **2️⃣ Get Columns for a Board** (Protected)

```http
GET /api/boards/1/columns
```

**Response:**

```json
[
  {
    "id": 1,
    "title": "To Do"
  }
]
```

---

### **📌 Task Management**

#### **1️⃣ Create a Task** (Protected)

```http
POST /api/tasks
```

**Request Body:**

```json
{
  "title": "Write Documentation",
  "description": "Complete API docs",
  "column_id": 1
}
```

**Response:**

```json
{
  "id": 1,
  "title": "Write Documentation",
  "description": "Complete API docs"
}
```

#### **2️⃣ Get Tasks by Column** (Protected)

```http
GET /api/columns/1/tasks
```

**Response:**

```json
[
  {
    "id": 1,
    "title": "Write Documentation",
    "description": "Complete API docs"
  }
]
```

---

## 🔒 **Basic Rules & Notes**

1. **User Signup** is only allowed with role `user` by default.
2. **JWT Token** must be included in all protected routes as `Authorization: Bearer <token>`.
3. **Only Admins** can create Boards and assign Users.
4. **Columns and Tasks** must be created within an existing Board.
5. **Tasks are assigned to the currently authenticated user.**

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature-name`)
3. Commit your changes (`git commit -m "Add new feature"`)
4. Push to your branch (`git push origin feature-name`)
5. Open a Pull Request

---

## 📝 License

This project is open-source and available under the MIT License.
