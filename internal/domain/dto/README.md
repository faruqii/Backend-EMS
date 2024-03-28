# DTO (Data Transfer Object)

## What is a DTO?

A Data Transfer Object (DTO) is a design pattern used in software development to transfer data between different layers or components of an application. It is a simple container object that holds data and provides getter and setter methods to access and modify the data.

## Why use DTOs?

DTOs are commonly used in scenarios where data needs to be transferred between different parts of an application, such as between the frontend and backend, or between different layers of a backend system. Some reasons to use DTOs include:

- **Data encapsulation**: DTOs encapsulate data and provide a clear contract for transferring data, helping to maintain data integrity and consistency.

- **Reduced network traffic**: By transferring only the required data, DTOs can help reduce network traffic and improve performance.

- **Versioning and compatibility**: DTOs can be used to handle versioning and compatibility issues between different parts of an application, allowing for easier maintenance and evolution of the system.

## How to use DTOs?

To use DTOs, follow these steps:

1. Define the DTO struct: Create a struct that represents the data you want to transfer. Add private fields for the data and public getter and setter methods to access and modify the data.

2. Transfer data: Use the DTO struct to transfer data between different layers or components of your application. Instantiate the DTO, set the data using the setter methods, and pass the DTO object to the target component.

3. Access data: In the target component, use the getter methods of the DTO to access the transferred data.

Here is an example of a simple DTO struct in Go:
```go
package entities

type User struct {
    id   int
    email string
    password string
}
```

```go
package dto

type UserLoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
```

In this example, we have defined a `User` struct to represent a user entity and a `UserLoginRequest` struct to represent a user login request. The `UserLoginRequest` struct is used to transfer login data between the frontend and backend of an application.

## Conclusion

DTOs are a useful design pattern for transferring data between different parts of an application. By encapsulating data and providing a clear contract for transferring data, DTOs can help improve data integrity, reduce network traffic, and handle versioning and compatibility issues. When using DTOs, make sure to define clear contracts for data transfer and follow best practices for data encapsulation and access.
```

