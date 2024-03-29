# docker build -t localhost:5001/go-grpc:1.0.0 .
# Build stage
FROM golang:1.21.0-alpine3.17 AS build
WORKDIR /app
COPY . .
RUN go build -o main .

# Run stage  
FROM alpine:3.18.3
WORKDIR /app
COPY --from=build /app/main .

EXPOSE 3002 3002

CMD ["/app/main"]