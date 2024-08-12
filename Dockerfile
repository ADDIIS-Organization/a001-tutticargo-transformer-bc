# Use the official Golang image as the base image
FROM golang:1.16-alpine as build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Install tzdata for timezone handling
RUN apk add --no-cache tzdata

# Set the timezone to America/Bogota
ENV TZ=America/Bogota
RUN cp /usr/share/zoneinfo/America/Bogota /etc/localtime && echo "America/Bogota" > /etc/timezone

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/main .

# Copy the .env file
COPY .env .

# Expose port 8080 to the outside world
EXPOSE 8080

# Set environment variables from .env file
ENV $(cat .env | xargs)

# Command to run the executable
CMD ["./main"]
