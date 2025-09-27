# build stage
FROM golang:1.23.0-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o shubham main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.20
WORKDIR /app

# means copy from the builder stage and put it in /app in . dir
COPY --from=builder /app/shubham .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
# COPY app.env .
COPY start.sh .
#  waits for the database to be up and accepting connections before the app starts.
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8080

CMD [ "/app/shubham" ]
ENTRYPOINT [ "/app/start.sh" ]
# start.sh script will run when the container starts,
# then /app/shubham will be passed as an argument to it.