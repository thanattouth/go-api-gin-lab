![Go CI](https://github.com/thanattouth/go-api-gin-lab/actions/workflows/go-ci.yml/badge.svg)

# Student API with Gin (CS362 Lab)

## Tech Stack
- **Language:** Go 1.21+
- **Framework:** [Gin Gonic](https://gin-gonic.com/)
- **Database:** SQLite3 (Lightweight & Reliable)

## API Endpoint
- GET       /students       : ดึงรายชื่อนักศึกษาทั้งหมด
- GET       /students/:id   : ดึงข้อมูลนักศึกษาตาม ID
- POST	    /students       : เพิ่มนักศึกษาใหม่ (มี Validation)
- PUT	    /students/:id   : แก้ไขข้อมูลนักศึกษา
- DELETE    /students/:id   : ลบข้อมูลนักศึกษา (HTTP 204)

## Getting Started

### 1. Prerequisites
- ติดตั้ง Go (Version 1.21 ขึ้นไป)
- GCC Compiler (สำหรับ SQLite3 driver)

### 2. Installation & Running
    ```bash
    # Clone the repository
    git clone [https://github.com/thanattouth/go-api-gin-lab.git](https://github.com/thanattouth/go-api-gin-lab.git)

    # Install dependencies
    go mod tidy

    # Start the server
    go run main.go