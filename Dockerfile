# Use Golang version 1.20 based on Alpine Linux as the base for the build stage
FROM golang:1.20-alpine AS builder

# Set the working directory in the container
WORKDIR /build
# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./
# Download project dependencies
RUN go mod download

# Copy all other files to the working directory
COPY . .
# Build our application into an executable named "server". We disable CGO and set the target OS to Linux.
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Use Alpine Linux version 3.14 as the base for the final stage
FROM alpine:3.14 AS final

# Copy the "server" executable from the build stage to "/bin/server" of our final image
COPY --from=builder /build/server /bin/server

# Set the entry point for our container. When the container is run, it will execute the command "/bin/server"
ENTRYPOINT ["/bin/server"]