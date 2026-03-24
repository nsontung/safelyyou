# SafelyYou API Server

## Introduction
This project implements a server for the SafelyYou API, which provides endpoints for managing users, devices, and alerts. The server is built using the Gin-Gonic router.

## Setup
1. To run the server, ensure you have Go installed and set up on your machine.
2. Clone the repository and navigate to the project directory.
3. Install the required dependencies using the following command:
```bash
go mod tidy
```
4. Run the following command to start the server:
```bash
go run main.go
```
5. The server will start on port 6733, load all the devices from the `devices.csv` file. You can access the API endpoints at `http://localhost:6733/api/v1`.

## Device Simulation
To simulate device heartbeats, you can use the following command in a separate terminal:
```bash
./device-simulator-linux-amd64
```
This will send heartbeat signals to the server, allowing you to test the API's functionality.

The output of the device simulator will be stored in the `results.txt` file.

## Unit Testing
To run the unit tests, use the following command:
```bash
go test ./...
```

## Project Structure
- `main.go`: The entry point of the application, where the server is initialized and routes are defined.
- `api/handlers.go`: Contains the handler functions for the API endpoints.
- `api/router.go`: Defines the API routes and associates them with the corresponding handler functions.
- `api/dto.go`: Contains the Data Transfer Object (DTO) definitions for the API requests and responses.
- `api/utils.go`: Contains utility functions for the API, such as parsing csv files.
- `service/device_service.go`: Contains the business logic for managing devices, including loading devices from a CSV file and handling device heartbeats.
- `service/device_service_test.go`: Contains unit tests for the device service.
- `syerror/syerror.go`: Defines custom error types for the application.
- `store/data_store.go`: Contains the `DataStore` struct and methods for managing device data.
- `model/device.go`: Defines the `Device` struct and related methods.
- `devices.csv`: A CSV file containing the initial device data to be loaded into the server.
- `device-simulator-linux-amd64`: A binary file for simulating device heartbeats.
- `results.txt`: A file where the output of the device simulator is stored.

# How do I use AI assistant for this project?
I use a couple of AI tools to help me with this project:

- I use a code editor with AI capabilities, such as Visual Studio Code with the GitHub Copilot extension, which provides code suggestions and helps me write code faster. 
- I also use a ChatGPT to ask questions about programming concepts, get help with debugging, and receive explanations for complex code snippets.
