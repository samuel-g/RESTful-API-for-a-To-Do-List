# To-Do List RESTful API

This project is a simple RESTful API for managing a to-do list. The API allows users to perform CRUD operations (Create, Read, Update, Delete) on to-do items. It's built using Go (Golang) with the Gorilla Mux router.

## Features

- **Create a new to-do item:** Add tasks to your to-do list.
- **Get all to-do items:** Retrieve a list of all tasks.
- **Update a to-do item:** Modify details of an existing task.
- **Delete a to-do item:** Remove a task from the list.

## Getting Started

### Prerequisites

- **Go (Golang):** Ensure that Go is installed on your machine. You can download it from [https://golang.org/dl/](https://golang.org/dl/).
- **Git:** You should have Git installed to clone the repository. Download it from [https://git-scm.com/](https://git-scm.com/).

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/go-todo-app.git
   cd go-todo-app
   ```

2. **Install dependencies:**

   This project uses the Gorilla Mux package for routing. Install it using:

   ```bash
   go get -u github.com/gorilla/mux
   ```

3. **Run the application:**

   Start the server by running:

   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8080`.

### API Endpoints

- **GET /todos**
  - Description: Retrieve all to-do items.
  - Response: A list of all to-do items in JSON format.

- **POST /todos**
  - Description: Create a new to-do item.
  - Body: JSON object containing `id`, `title`, and `status`.
  - Example Request Body:
    ```json
    {
      "id": "1",
      "title": "Learn Go",
      "status": "Pending"
    }
    ```
  - Response: The created to-do item in JSON format.

- **PUT /todos/{id}**
  - Description: Update an existing to-do item.
  - Parameters: `id` (path parameter).
  - Body: JSON object containing updated `title` and `status`.
  - Example Request Body:
    ```json
    {
      "title": "Learn Go in depth",
      "status": "In Progress"
    }
    ```
  - Response: The updated to-do item in JSON format.

- **DELETE /todos/{id}**
  - Description: Delete a to-do item.
  - Parameters: `id` (path parameter).
  - Response: A list of remaining to-do items in JSON format.

### Example Usage

Use a tool like [Postman](https://www.postman.com/) or `curl` to interact with the API:

- **Create a new to-do item:**

   ```bash
   curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d '{"id":"1", "title":"Learn Go", "status":"Pending"}'
   ```

- **Get all to-do items:**

   ```bash
   curl http://localhost:8080/todos
   ```

- **Update a to-do item:**

   ```bash
   curl -X PUT http://localhost:8080/todos/1 -H "Content-Type: application/json" -d '{"title":"Learn Go in depth", "status":"In Progress"}'
   ```

- **Delete a to-do item:**

   ```bash
   curl -X DELETE http://localhost:8080/todos/1
   ```