# Weather API Go

## Overview
This project implements a Weather API in Go that provides weather updates and subscription management. It adheres to the Swagger API specification for clear documentation and ease of use.

## Project Structure
```
weather-api-go
├── cmd
│   └── main.go            # Entry point of the application
├── internal
│   ├── api
│   │   ├── handlers.go    # HTTP handler functions for API endpoints
│   │   └── routes.go      # API route definitions
│   ├── models
│   │   └── weather.go     # Data structures for weather and subscriptions
│   └── service
│       └── weather_service.go # Business logic for weather data and subscriptions
├── docs
│   └── swagger.yaml       # Swagger API specification
├── go.mod                 # Go module definition
├── go.sum                 # Module dependency checksums
└── README.md              # Project documentation
```

## Setup Instructions
1. Clone the repository:
   ```
   git clone <repository-url>
   cd weather-api-go
   ```

2. Install the dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run cmd/main.go
   ```

4. Access the API at `http://localhost:8080/weather`.

## Usage Examples
- Get weather data:
  ```
  GET /weather?city=London
  ```

- Subscribe to weather updates:
  ```
  POST /subscribe
  ```

- Confirm subscription:
  ```
  POST /confirm
  ```

- Unsubscribe from updates:
  ```
  DELETE /unsubscribe
  ```

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.