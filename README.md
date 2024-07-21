
# Introduction

This is a microservice made with GO that provides an endpoint to transform an Excel file into DB records.

## How to run

To run this microservice, you need to have GO installed on your machine. You can download it [here](https://golang.org/dl/).

After installing GO, you can run the following command to start the microservice:

```bash
go run main.go config.go models.go handlers.go
```

This command will start the microservice on port 8081.

## Project Structure

```
- middleware/
  - logging.go
- config.go
- go.mod
- go.sum
- handlers.go
- insertion.log
- main.go
- models.go
- README.md
```

### Description

- **middleware/logging.go**: Contains middleware for logging requests.
- **config.go**: Contains configuration settings for the microservice.
- **go.mod**: The Go module file.
- **go.sum**: The Go checksum file.
- **handlers.go**: Contains HTTP handlers for the endpoints.
- **insertion.log**: Log file for recording data insertions.
- **main.go**: Entry point for the microservice.
- **models.go**: Contains data models used in the microservice.
