# Use the official Golang image as the base image
FROM golang:alpine as builder

ENV GO111MODULE=on
# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the container
# Copy the Go module files to the container
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .




# Build the microservice binary
RUN go build -o bin . 
# Copy the "app" configuration file
COPY app.env /app/app
# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app


# Copy the binary from the builder stage to the final container
COPY --from=builder  /app/bin .
COPY --from=builder /app/app .
# Expose the gRPC and RESTful ports used by the microservice
EXPOSE 9097 8087

# Run the microservice binary
RUN chmod +x ./bin
CMD ["./bin"]
