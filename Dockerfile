# Use the official Golang image to create a build artifact.
# This is based on Debian and is a larger image that includes a full build toolset.
FROM golang:1.21.3 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.* ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/
# Use alpine:latest as the base image for the final stage of the Docker multi-stage build
FROM alpine:latest

# Install CA certificates for TLS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/api .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./api"]
