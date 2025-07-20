  PulseChat Backend - README body { font-family: Arial, sans-serif; line-height: 1.6; padding: 20px; max-width: 800px; margin: auto; background-color: #f9f9f9; } code, pre { background-color: #eee; padding: 4px 6px; border-radius: 4px; font-family: Consolas, monospace; } pre { background-color: #f0f0f0; padding: 10px; overflow-x: auto; } table { width: 100%; border-collapse: collapse; margin: 16px 0; } table, th, td { border: 1px solid #ccc; } th, td { padding: 8px; text-align: left; } h1, h2, h3 { color: #333; } hr { margin: 24px 0; }

🧠 PulseChat Backend
===================

This is the **backend** for a real-time chat application with support for user registration, login (with JWT), profile management, and messaging via WebSockets.

* * *

🚀 Features
-----------

*   ✅ User Registration & Login (with hashed passwords)
*   🔐 JWT-based Authentication
*   🧾 Profile retrieval
*   💬 Real-time Messaging using WebSockets
*   📦 RESTful API Endpoints (JSON)
*   🔌 CORS-enabled for frontend integration

* * *

🛠️ Tech Stack
--------------

*   **Language:** Go (Golang)
*   **Router:** Gorilla Mux
*   **Database:** PostgreSQL
*   **Authentication:** JWT
*   **Realtime:** WebSocket (Gorilla WebSocket)
*   **Environment:** `.env` for config

* * *

📁 Project Structure
--------------------

    .
    ├── main.go               # Entry point
    ├── handlers/             # All HTTP and WebSocket handlers
    ├── models/               # DB models and structs
    ├── middleware/           # JWT and CORS middleware
    ├── utils/                # Utility functions (token generation etc.)
    ├── db/                   # DB connection setup
    ├── go.mod / go.sum       # Go module files
    ├── .env                  # Env variables (DB URL, secret, etc.)
    

* * *

🔧 Setup Instructions
---------------------

### 1\. Clone the repo

    git clone https://github.com/your-username/chat-app-backend.git
    cd chat-app-backend

### 2\. Create `.env` file

    PORT=8080
    JWT_SECRET=your_jwt_secret
    DB_URL=postgres://username:password@localhost:5432/chatdb?sslmode=disable

### 3\. Install dependencies

    go mod tidy

### 4\. Run the server

    go run main.go

* * *

📬 API Endpoints
----------------

| Method | Endpoint    | Description             | Auth Required |
| ------ | ----------- | ----------------------- | ------------- |
| POST   | `/register` | Register new user       | ❌             |
| POST   | `/login`    | Login and get JWT token | ❌             |
| GET    | `/profile`  | Get current user info   | ✅             |
| GET    | `/ws`       | WebSocket connection    | ✅ (via token) |


**Note:** All protected routes require the `Authorization: Bearer <token>` header.

* * *

📡 WebSocket
------------

Connect to `/ws` with JWT token as query param:

    ws://localhost:8080/ws?token=your_jwt_token

Expected message structure:

    {
      "to": "receiver_username",
      "message": "Hello!"
    }

* * *

🧪 Testing
----------

Use [Postman](https://www.postman.com/) or [curl](https://curl.se/) to test the HTTP endpoints.

* * *

🛡️ Security
------------

*   Passwords are hashed using bcrypt.
*   JWT used for securing endpoints.
*   CORS enabled for frontend access.

* * *

👨‍💻 Author
------------

Built with ❤️ by **Your Name**

* * *

📄 License
----------

This project is licensed under the **MIT License**.