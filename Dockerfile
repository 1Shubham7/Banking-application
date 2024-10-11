FROM golang:1.23.0-alpine3.20

# this will be our working dir
WORKDIR /app

# now lets copy all the ifles to the app folder
# copy source destination, and the our destination is
# /app in . dir (base dir)
COPY . .

# now let's build the binary (I am making a binary called
# shubham from main.go)
RUN go build -o shubham main.go

# this is just for documentation on which port will be exposed
EXPOSE 8080

# this is the default command that will run when the container starts
# it is an array for CLI arguments
# this will run the executable file
CMD [ "/app/shubham" ]