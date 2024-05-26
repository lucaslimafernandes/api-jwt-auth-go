# api-jwt-auth-go

## Description

The api-jwt-auth-go project is an API developed in Go (Golang) that implements authentication and authorization using JSON Web Tokens (JWT). This API offers a secure and scalable framework to protect endpoints, allowing only authenticated users to access specific resources.

## Key Features

- User Registration: Allows new users to register in the application.

- Authentication: Verifies user credentials to issue a JWT.

- Authorization: Protects endpoints by verifying the JWT provided in subsequent requests.

- Token Renewal: Provides functionality to renew JWTs that are nearing expiration.

- Logout: Invalidates JWTs, ensuring that users cannot continue using revoked tokens.


## Technologies Used

- Golang: The programming language used to develop the API.

- JWT: A token standard for secure authentication.

- Gin: A web framework for Go, used to facilitate the development of RESTful APIs.

- GORM: An ORM for Go, used for database interactions.

- PostgreSQL/MySQL: Relational databases for storing user information and tokens.

## Project Structure

- main.go: The entry point of the application.

- models/: Defines data models and database interactions.

- controllers/: Handles the logic for API requests and responses.

- routes/: Defines the routes and their respective handlers.

- middlewares/: Middleware for JWT verification and other security features.

- migrate/: Manages database migrations or automigrate, included in init into main file.


## How to Use

1. Clone the Repository:

    git clone https://github.com/your-username/api-jwt-auth-go.git


2. Install Dependencies:

    go mod download


3. Set Up Environment Variables: Configure the necessary environment variables in the .env file.

4. Run docker postgres

    sudo docker compose -f stack.yaml -d     

5. Run the Application:

    go run main.go

## Final Considerations

The api-jwt-auth-go project is ideal for developers looking for a ready-made and secure solution for authentication and authorization in their Go applications. With a modular and easy-to-understand architecture, this project can be expanded and adapted to different needs and application scales.