<h1 align="center">student-api</h1>

<div align="center">

This microservice was developed in Go, implementing the best practices of the RESTful standard. It offers a robust and efficient API, with endpoints that support full CRUD (Create, Read, Update and Delete) operations.

[![Go Version](https://img.shields.io/badge/Go-1.23.2-blue)](https://go.dev/doc/devel/release#go1.23.0)

</div>


## ✨ Features

- Healthcheck
- Create student
- Read student
- Update student
- Delete student
- Search student

## ⚙️ Main technologies

- Go
- Echo Web Framework
- Open Telemetry
- Otel Collector + Zipkin
- Zap Logger
- FX Dependency Injection System
- Go's database package + Postgres
- golang-migrate
- Docker
- Docker Compose
- Go's testing package
- Testcontainers (integration tests)
- Httpexpect  (end-to-end tests)

## 🤓 Main Code Patterns

This project also incorporates the following key design and code patterns to ensure robustness and maintainability:

- **Graceful Shutdown:** Ensures the application handles termination signals properly, cleaning up resources like connections, goroutines, and processes.  
- **Decorator:** Adds behavior to existing structures or functions dynamically without modifying their source code.  
- **Builder:** Provides a clear and fluent way to construct complex objects step by step.  
- **Repository:** Abstracts data access logic, separating it from business logic and enabling easier testing and extensibility.

## 🏢 Architecture

This project follows the principles of **Clean Architecture**, promoting a clear separation of responsibilities and ensuring modular, testable and easy-to-maintain code. The layers used are:

- **Domain:**  
  Contains the core and independent elements of the system, including:  
  - **Entities:** Fundamental business rules.  
  - **Use Case Interfaces:** Define contracts that use cases must implement.  
  - **Repositories:** Interfaces for data persistence, abstracting implementation details.  
  - **DTOs:** Data Transfer Objects for clear communication between layers.  
  - **Mocks:** Simulated implementations to facilitate isolated testing and development.  

- **Adapters:**  
  Responsible for communication between the domain and the external world, such as inputs and outputs:  
  - **Handlers:** Entry layer exposing endpoints and translating requests into calls to use cases.  
  - **Repositories:** Concrete implementations of the repositories defined in the domain, connecting to the database.  

- **Infrastructure:**  
  Deals with technical details such as configurations, external integrations, and database connectivity.  

- **Use Cases:**  
  Contains the application logic, orchestrating operations between entities, repositories, and other layers.  

This organization ensures the independence of the domain from frameworks and technical details, providing greater flexibility and ease of maintenance.

## 🚀 How to run the project in local mode

```
1. Clone the project repository:
$ git clone https://github.com/dmarins/student-api.git

2. Access the project folder on your terminal:
$ cd student-api

3. Install dependencies:
$ go mod tidy

4. Run tests (make sure you are running docker on your machine):
$ make tests

6. Run the application in local mode (make sure you are running docker on your machine):
$ make local-restart
$ make run

7. Go to http://localhost:8080/swagger/index.html in your browser or see all http files in /.requests directory and tests all endpoints.

8. To stop all containers
$ make down
```

## 📦 How to run the project in container mode

```
1. Clone the project repository:
$ git clone https://github.com/dmarins/student-api.git

2. Access the project folder on your terminal:
$ cd student-api

3. Install dependencies:
$ go mod tidy

4. Run tests (make sure you are running docker on your machine):
$ make tests

6. Run the application in container mode (make sure you are running docker on your machine):
$ make docker-restart

7. Go to http://localhost:8080/swagger/index.html in your browser or see all http files in /.requests directory and tests all endpoints.

8. To stop all containers
$ make down
```