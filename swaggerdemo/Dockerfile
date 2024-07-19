# Start with the official Golang image as the base image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files from the demo folder
COPY demo/go.mod demo/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the demo folder into the container
COPY demo ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /swaggerdemo main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/swaggerdemo"]
