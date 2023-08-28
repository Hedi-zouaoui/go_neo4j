# Stage 1: Build the Golang binary
FROM golang:alpine as builder

ENV GO111MODULE=on

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the container
COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy the application source code into the container
COPY . .

# Build the microservice binary
RUN go build -o bin .

FROM nginx:alpine
RUN rm -rf /usr/share/nginx/html/*
# Copy and move your own website file to the nginx html folder
COPY index.html /usr/share/nginx/html
# Copy the "app" configuration file
COPY app.env /app/app

# Set the working directory inside the container
WORKDIR /app



# Expose the gRPC and RESTful ports used by the microservice
EXPOSE 8080

# Start Nginx to serve the static files
CMD ["nginx", "-g", "daemon off;"]
