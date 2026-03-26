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
go test -v ./...
```
We have some testing files that cover the functionality of the device service, ensuring that the business logic is working as expected.
- `service/device_service_test.go` contains unit tests for the device service, covering scenarios for some edge cases...
- `api_test.go` contains unit tests for the API endpoints, ensuring that the API responds correctly to various requests and handles edge cases appropriately.

## Project Structure
- `main.go`: The entry point of the application, where the server is initialized and routes are defined.
- `api/handlers.go`: Contains the handler functions for the API endpoints.
- `api/router.go`: Defines the API routes and associates them with the corresponding handler functions.
- `api/dto.go`: Contains the Data Transfer Object (DTO) definitions for the API requests and responses.
- `api/utils.go`: Contains utility functions for the API, such as parsing csv files.
- `service/device_service.go`: Contains the business logic for managing devices
- `syerror/syerror.go`: Defines custom error types for the application.
- `store/data_store.go`: Contains the `DataStore` struct and methods for managing device data.
- `model/device.go`: Defines the `Device` struct and related methods.
- `devices.csv`: A CSV file containing the initial device data to be loaded into the server.
- `device-simulator-linux-amd64`: A binary file for simulating device heartbeats.
- `results.txt`: A file where the output of the device simulator is stored.

## How do I use AI assistant for this project?
I use a couple of AI tools to help me with this project:

- I use a code editor with AI capabilities, such as Visual Studio Code with the `GitHub Copilot` extension, which provides code suggestions and helps me write code faster. 
- I also use a `ChatGPT` to ask questions about programming concepts, get help with debugging, and receive explanations for complex code snippets.

## Answer to some questions about the project:

1. How long did you spend working on the problem? What did you find to be the most difficult part?
I spent approximately 10 hours working on this project. The most difficult part was designing the data model and ensuring that it could efficiently handle the various operations required by the API, such as loading devices from a CSV file, managing device heartbeats, and handling alerts. Additionally, implementing the unit tests to cover all edge cases was also challenging.
One thing I found particularly difficult was calculating the time difference between the current time and the last heartbeat timestamp to determine if a device is offline. I had to ensure that the logic was accurate and efficient, especially when dealing with the simulation of multiple devices sending heartbeats simultaneously. 

2. How would you modify your data model or code to account for more kinds of metrics?
To account for more kinds of metrics, I would modify the `Device` struct to include a map or a slice of metric types and their corresponding values. This way, we can easily add new metrics without changing the overall structure of the device. For example, we could have a `Metrics` field that is a map where the key is the metric name (e.g., "heart_rate", "temperature") and the value is the metric value. This would allow us to easily extend the functionality to support additional metrics in the future without needing to modify existing code significantly.

3. Discuss your solution’s runtime complexity.
The runtime complexity of the solution depends on the specific operations being performed. For example, loading devices from a CSV file has a time complexity of O(n), where n is the number of devices in the file. Managing device heartbeats involves updating the device's last heartbeat timestamp, which has a time complexity of O(1) for each heartbeat. Retrieving device information or alerts also has a time complexity of O(1) for each request, as we are using a map to store and access device data efficiently. Overall, the solution is designed to be efficient and scalable, with most operations having constant or linear time complexity.

