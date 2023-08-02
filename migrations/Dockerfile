# Use Golang version 1.20 based on Alpine Linux as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /queries

# Install goose, a database migration tool in Go
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy the contents of the local 'queries' directory into the working directory of the container
COPY queries .

# Set the command that will be executed when the container starts
ENTRYPOINT [ "goose" ]