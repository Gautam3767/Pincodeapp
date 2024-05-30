# Use the official Golang image as the base image
FROM golang:1.21-alpine

# Enable CGO and install necessary dependencies
RUN apk update && apk add --no-cache gcc musl-dev sqlite-dev

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application with CGO enabled
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
