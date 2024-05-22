# Todo API

This is a simple Todo API built with Go and MongoDB. It provides a RESTful interface to manage a todo list. The API supports creating, reading, updating, and deleting todo items.

## Prerequisites

- Go (version 1.16 or later)
- MongoDB (version 4.2 or later)

## Installation

1. Clone the repository:

```
git clone https://github.com/your-username/todo-api.git
```

2. Change to the project directory:

```
cd todo-api
```

3. Install the dependencies:

```
go get ./...
```

## Configuration

The API expects a MongoDB server running on `localhost:27017`. If your MongoDB server is running on a different host or port, you'll need to update the connection string in `main.go`:

```go
client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://your-mongodb-host:port"))
```

## Running the Application

1. Start the MongoDB server (if it's not already running).

2. Run the Go application:

```
go run main.go
```

The API will be available at `http://localhost:8080`.

## API Endpoints

### Get All Todo Items

```
GET /todo
```

Query parameters:

- `list` (optional): Filter todo items by list name.
- `sort` (optional, default: `timestamp`): Sort todo items by the specified field (`timestamp` or `text`).

### Create a Todo Item

```
POST /todo
```

Request body:

```json
{
  "text": "Buy groceries",
  "timestamp": "2023-05-23T14:30:00Z",
  "list": "Shopping"
}
```

### Update a Todo Item

```
PUT /todo/:id
```

Request body:

```json
{
  "text": "Buy bread and milk",
  "timestamp": "2023-05-23T15:00:00Z",
  "list": "Shopping"
}
```

### Delete a Todo Item

```
DELETE /todo/:id
```

## Testing

The project includes a basic test case for the `GET /todo` endpoint. To run the tests, execute the following command:

```
go test
```

## Docker

The project includes a `Dockerfile` and a `docker-compose.yml` file for easy deployment with Docker. To run the application using Docker Compose, execute the following command:

```
docker-compose up
```

This will build the Go application and the MongoDB container, and start both services.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.
