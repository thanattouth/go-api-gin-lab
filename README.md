![Go CI](https://github.com/thanattouth/go-api-gin-lab/actions/workflows/go-ci.yml/badge.svg)

# Student API with Gin (CS362 Lab)

## How to Run
1. Install dependencies:
   ```bash
   go mod tidy

2. Run the server:
    ```bash
    go run main.go

API Endpoints
    - GET /students - Get all students
    - GET /students/:id - Get student by ID
    - POST /students - Create new student
    - PUT /students/:id - Update student
    - DELETE /students/:id - Delete student

Tech Stack
    - Go (Gin Framework)
    - SQLite3