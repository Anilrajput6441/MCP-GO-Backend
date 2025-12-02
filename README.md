# MCP Go Backend

A robust RESTful API backend built with Go (Gin framework) and MongoDB for task management with user authentication and MCP (Model Context Protocol) integration for AI-powered task operations.

##  Features

- **User Authentication**: JWT-based authentication with registration, login, and token refresh
- **Task Management**: Full CRUD operations for tasks with user-specific access
- **Partial Updates**: Update only the fields you provide (title, description, or status)
- **MCP Integration**: Special endpoints for AI agents to interact with tasks
- **User Management**: Profile management and password change functionality
- **Security**: Password hashing, JWT tokens, and CORS middleware
- **MongoDB**: NoSQL database for flexible data storage

##  Tech Stack

- **Language**: Go 1.24+
- **Web Framework**: Gin
- **Database**: MongoDB
- **Authentication**: JWT (golang-jwt/jwt/v5)
- **Password Hashing**: bcrypt (via golang.org/x/crypto)
- **Environment Variables**: godotenv

##  Prerequisites

- Go 1.24 or higher
- MongoDB (local or cloud instance)
- Git

##  Installation & Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd MCP-GO-Backend
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   
   Create a `.env` file in the root directory:
   ```env
   MONGO_URI=mongodb://localhost:27017
   JWT_SECRET=your-secret-key-here
   JWT_REFRESH_SECRET=your-refresh-secret-key-here
   ```

4. **Start MongoDB**
   
   Make sure MongoDB is running on your system. If using a cloud instance, update the `MONGO_URI` in your `.env` file.

5. **Run the server**
   ```bash
   go run cmd/server/main.go
   ```

   The server will start on `http://localhost:8080`

##  Project Structure

```
MCP-GO-Backend/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── env.go              # Environment configuration
│   ├── db/
│   │   └── mongo.go            # MongoDB connection
│   ├── handlers/
│   │   ├── authHandler.go      # Authentication handlers
│   │   ├── taskHandler.go      # Task CRUD handlers
│   │   ├── userHandler.go      # User management handlers
│   │   ├── userSecurityHandler.go # Password change handler
│   │   └── mcp.go              # MCP-specific handlers
│   ├── middleware/
│   │   ├── authMiddleware.go   # JWT authentication middleware
│   │   └── cors.go             # CORS configuration
│   ├── models/
│   │   ├── user.go             # User data model
│   │   └── task.go             # Task data model
│   ├── routes/
│   │   ├── routes.go           # Main route registration
│   │   ├── authRoute.go        # Auth routes
│   │   ├── taskRoute.go        # Task routes
│   │   ├── userRoute.go        # User routes
│   │   └── mcp.go              # MCP routes
│   ├── services/
│   │   ├── authService.go      # Authentication business logic
│   │   └── taskService.go      # Task business logic
│   └── utils/
│       └── jwt.go              # JWT utility functions
├── docs/                       # Documentation
├── go.mod                      # Go module dependencies
└── go.sum                      # Dependency checksums
```

##  Authentication

All protected routes require a JWT token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

### Authentication Flow

1. **Register**: Create a new user account
2. **Login**: Get access token and refresh token
3. **Refresh**: Get a new access token using refresh token
4. **Protected Routes**: Use access token for authenticated requests

##  API Endpoints

### Health Check

- `GET /health` - Server health check
- `GET /ping` - Simple ping endpoint

### Authentication (`/auth`)

- `POST /auth/register` - Register a new user
  ```json
  {
    "email": "user@example.com",
    "password": "password123",
    "full_name": "John Doe"
  }
  ```

- `POST /auth/login` - Login and get tokens
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```
  Response:
  ```json
  {
    "access_token": "...",
    "refresh_token": "..."
  }
  ```

- `POST /auth/refresh` - Refresh access token
  ```json
  {
    "refresh_token": "..."
  }
  ```

### Tasks (`/tasks`) - Protected

- `POST /tasks` - Create a new task
  ```json
  {
    "title": "Task title",
    "description": "Task description"
  }
  ```
  Response:
  ```json
  {
    "id": "task_id"
  }
  ```

- `GET /tasks` - Get all tasks for the authenticated user
  Response:
  ```json
  [
    {
      "id": "...",
      "title": "...",
      "description": "...",
      "status": "todo",
      "user_email": "...",
      "created_at": "...",
      "updated_at": "..."
    }
  ]
  ```

- `PUT /tasks/:id` - Update a task (partial updates supported)
  ```json
  {
    "title": "Updated title",      // Optional
    "description": "Updated desc",  // Optional
    "status": "in-progress"         // Optional
  }
  ```
  **Note**: Only include the fields you want to update. Empty strings will be ignored.

- `DELETE /tasks/:id` - Delete a task

### Users (`/users`) - Protected

- `GET /users` - Get user profile
- `PUT /users` - Update user profile
- `DELETE /users` - Delete user account
- `PUT /users/change-password` - Change password

### MCP Endpoints (`/mcp/task`) - Protected

Special endpoints designed for AI agents using the Model Context Protocol:

- `POST /mcp/task/list` - List all tasks for the authenticated user
  ```json
  {
    "tasks": [...]
  }
  ```

- `POST /mcp/task/create` - Create a task via AI
  ```json
  {
    "title": "Task title",
    "description": "Task description"
  }
  ```

- `POST /mcp/task/update/:id` - Update a task via AI
  ```json
  {
    "id": "task_id",
    "title": "Updated title",
    "description": "Updated description",
    "status": "done"
  }
  ```

- `DELETE /mcp/task/delete/:id` - Delete a task via AI
  ```json
  {
    "id": "task_id"
  }
  ```

## 🎯 Task Status

Tasks can have the following statuses:
- `todo` - Task is pending (default)
- `in-progress` - Task is being worked on
- `done` - Task is completed

## 🔒 Security Features

- **Password Hashing**: Passwords are hashed using bcrypt
- **JWT Tokens**: Secure token-based authentication
- **CORS**: Cross-Origin Resource Sharing configured
- **User Isolation**: Users can only access their own tasks
- **Input Validation**: Request validation on all endpoints

## 💡 Key Features Explained

### Partial Updates

The `UpdateTask` function supports partial updates. You can update only the fields you provide:

```go
// Only update status
PUT /tasks/:id
{
  "status": "done"
}

// Only update title and description
PUT /tasks/:id
{
  "title": "New title",
  "description": "New description"
}
```

Empty strings are ignored, so you can safely omit fields you don't want to change.

## 🗄️ Database

The application uses MongoDB with the following collections:

- **users**: User accounts and authentication data
- **tasks**: Task data with user association

Database name: `trio_assistant` (configurable via MongoDB URI)

## 🧪 Testing

You can test the API using tools like:
- Postman
- cURL
- HTTPie
- Any REST client

Example cURL request:
```bash
# Register
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123","full_name":"Test User"}'

# Login
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# Create Task (with token)
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{"title":"My Task","description":"Task description"}'
```

## 📝 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `MONGO_URI` | MongoDB connection string | `mongodb://localhost:27017` |
| `JWT_SECRET` | Secret key for JWT access tokens | Required |
| `JWT_REFRESH_SECRET` | Secret key for JWT refresh tokens | Required |

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License.

## 👤 Author

Anil Rajput

## 🙏 Acknowledgments

- Gin Web Framework
- MongoDB Go Driver
- JWT Go Library

