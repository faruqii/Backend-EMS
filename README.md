# Backend Application for Education Management System (EMS) 

## Overview
Welcome to the Education Management System Backend App! This application is built using Go, PostgreSQL, and adheres to Domain-Driven Design (DDD) principles to ensure a robust, scalable, and maintainable codebase.

## Introduction
The Education Management System Backend App is designed to manage educational institutions' data efficiently. It supports functionalities such as managing students, courses, instructors, and grades. The backend is implemented in Go using Fiber as framework, utilizing PostgreSQL as the database, and follows the principles of Domain-Driven Design (DDD) for a well-structured codebase.

## Tech Stack
- **Language**: Go
- **Framework**: Fiber
- **Database**: PostgreSQL
- **Containerization**: Docker
- **Architecture**: Domain-Driven Design (DDD)

## Architecture
The application is structured using Domain-Driven Design (DDD) principles. The codebase is divided into multiple layers, each with its own responsibilities. The main layers are:
- **Domain Layer**: Contains the core business logic and domain models.
- **Repository Layer**: Handles the database operations and queries.
- **Service Layer**: Implements the use cases and business logic.
- **Handler Layer**: Handles the HTTP requests and responses.
- **Middleware Layer**: Contains the middleware functions for the application.

## Dependency Injection
Dependency Injection (DI) is a technique used to achieve Inversion of Control (IoC) between classes and their dependencies. In this application, DI helps in decoupling the creation of dependencies from the business logic, making the code more modular and easier to test.

### How Dependency Injection is Implemented
- **Repository Injection**: Repositories are injected into services to provide data access.
```go
type StudentRepository interface {
    GetByID(id uint) (*models.Student, error)
    Create(student *models.Student) error
    Update(student *models.Student) error
    Delete(id uint) error
}

type StudentService struct {
    repo repositories.StudentRepository
}

func NewStudentService(repo repositories.StudentRepository) *StudentService {
    return &StudentService{repo: repo}
}
```

- **Service Injection**: Services are injected into handlers to implement business logic.
```go
type StudentService interface {
    GetStudentByID(id uint) (*models.Student, error)
    CreateStudent(student *models.Student) error
    UpdateStudent(id uint, student *models.Student) error
    DeleteStudent(id uint) error
}

type StudentHandler struct {
    service services.StudentService
}

func NewStudentHandler(service services.StudentService) *StudentHandler {
    return &StudentHandler{service: service}
}
```

- **Handler Injection**: Handlers are injected into the main application to handle HTTP requests.
```go
type StudentHandler struct {
    service services.StudentService
}

func NewStudentHandler(service services.StudentService) *StudentHandler {
    return &StudentHandler{service: service}
}
```

- **Router Injection**: Routers are injected with handlers to define the API routes.
```go
func NewStudentRouter(handler *handlers.StudentHandler) *fiber.Router {
    router := fiber.New()

    router.Get("/:id", handler.GetStudentByID)
    router.Post("/", handler.CreateStudent)
    router.Put("/:id", handler.UpdateStudent)
    router.Delete("/:id", handler.DeleteStudent)

    return router
}
```

### Example of Dependency Injection
```go
func main() {
    // Initialize the database connection
    db, err := database.NewDatabase()
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    // Initialize the repositories
    studentRepo := repositories.NewStudentRepository(db)

    // Initialize the services
    studentService := services.NewStudentService(studentRepo)

    // Initialize the handlers
    studentHandler := handlers.NewStudentHandler(studentService)

    // Initialize the routers
    studentRouter := routers.NewStudentRouter(studentHandler)

    // Create a new Fiber app
    app := fiber.New()

    // Register the student router
    app.Use("/students", studentRouter)

    // Start the server
    log.Fatal(app.Listen(":3000"))
}
```

## Folder Structure
The folder structure of the application is designed to be modular and scalable. Each layer of the application has its own folder, and the code is organized based on its functionality.

```
├───cmd
└───internal
    ├───app
    ├───config
    │   ├───database
    │   └───seeder
    ├───domain
    │   ├───dto
    │   ├───entities
    │   └───repositories
    ├───handlers
    ├───helper
    ├───middleware
    ├───mocks
    ├───routes
    └───services
```


