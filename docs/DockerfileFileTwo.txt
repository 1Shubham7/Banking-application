# build stage
FROM golang:1.23.0-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o shubham main.go

# Run stage
FROM alpine:3.20
WORKDIR /app

# means copy from the builder stage and put it in /app in . dir
COPY --from=builder /app/shubham .

EXPOSE 8080

CMD [ "/app/shubham" ]