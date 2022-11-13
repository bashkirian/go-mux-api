# Start from golang base image
FROM golang:1.19-alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Almaz Shagiev <elmuzon@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /balance

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /balance/main .
COPY --from=builder /balance/.env .       

# Expose port 8010 to the outside world
EXPOSE 8010

#Command to run the executable
CMD ["./main"]