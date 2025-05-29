**Problem:**  
Build a task management system that allows users to:
- Create, read, update, and delete tasks
- Filter tasks by status
- Support pagination
- Be designed as a microservice with clear modularity

**Design Decisions:**
- Used the Gin framework for clean and fast HTTP routing.
- Followed Single Responsibility Principle by separating concerns into packages: `controllers`, `models`, `routes`, and `storage`.
- Used in-memory storage for simplicity and stateless design.
- Prepared the architecture for scaling and future database integration.


##  Instructions to Run the Service

### Prerequisites:
- Go 1.21+
- Git (to clone the repo)

### Steps:
```bash
git clone https://github.com/sudhanshubhushan4535/task-service.git
cd task-service
go mod tidy
go run main.go
The server will run on http://localhost:8080

API Documentation

POST /tasks
Create a new task
Request Body:
{
  "title": "Read book",
  "status": "Pending"
}

GET /tasks
Get all tasks
Optional Query Params:
status, page, limit

GET /tasks?status=Pending&page=1&limit=5

GET /tasks/:id
Fetch a task by ID

PUT /tasks/:id
Update a task
Request Body:
{
  "title": "Buy groceries",
  "status": "Completed"
}

DELETE /tasks/:id
Delete a task by ID