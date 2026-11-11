
# Student Management System

A REST API built with Golang and Echo Framework.

## Technologies
- Go 1.26.5
- Echo Framework
- REST API
- JSON
- Git and GitHub

## Features
- Get all students
- Get student by ID
- Add a new student
- Update student information
- Delete a student

## Installation

Install dependencies:

```bash
go mod tidy
```

Run the application:

```bash
go run main.go
```

## Server

http://localhost:8080

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| GET | /students | Get all students |
| GET | /students/:id | Get student |
| POST | /students | Add student |
| PUT | /students/:id | Update student |
| DELETE | /students/:id | Delete student |
