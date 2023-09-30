FROM golang:alpine as builder


ENV GO111MODULE=on
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o bin . 



FROM alpine:latest

WORKDIR /app

COPY --from=builder  /app/bin .
COPY --from=builder /app/app .


EXPOSE 9097 8087

RUN chmod +x ./bin
CMD ["./bin"]
