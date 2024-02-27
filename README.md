
# JWT Authentication in Golang

This project demonstrates how to implement JWT (JSON Web Token) authentication in a Golang web application using Gin framework.

## Features

- User signup: Users can register by providing a username, email, and password.
- User login: Registered users can log in using their email and password.
- JWT Authentication: JWT is used to generate and authenticate user tokens for secure access to protected routes.

## Prerequisites

- Go (v1.16 or later)
- PostgreSQL (v9.5 or later)

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/CABON-TECH/jwt-auth-in-golang.git
    ```

2. Navigate to the project directory:

    ```bash
    cd jwt-auth-in-golang
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Set up the PostgreSQL database:

    - Create a PostgreSQL database named `jwt_db`.
    - Update the database connection details in `initializers/connectDB.go` file.

5. Set environment variables:

    Create a `.env` file in the root directory and define the following environment variables:

    ```plaintext
    SECRET=your-secret-key
    ```

6. Run the application:

    ```bash
    go run main.go
    ```

7. Access the application in your web browser at `http://localhost:8080`.

## API Endpoints

### User Signup

- **URL:** `/signup`
- **Method:** POST
- **Body:**

    ```json
    {
        "username": "Cabon TEch",
        "email": "cabontech@gmail.com",
        "password": "cabon123"
    }
    ```

### User Login

- **URL:** `/login`
- **Method:** POST
- **Body:**

    ```json
    {
        "email": "cabontech@gmail.com",
        "password": "cabon123"
    }
    ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
