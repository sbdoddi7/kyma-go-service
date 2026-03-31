# Using mulistage to reduce size of the image
FROM golang:1.26-alpine AS builder 

# Setting the working directory
WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o app ./cmd/server

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]