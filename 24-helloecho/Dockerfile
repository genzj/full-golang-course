# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.12.7-alpine3.10 as builder

# Add Maintainer Info
LABEL maintainer="ZHU Jie <zj0512@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# China goproxy accelerating
ENV GOPROXY=https://goproxy.io

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .


######## Start a new stage from scratch #######
FROM alpine:3.10

# use aliyun mirror
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 1323

# Command to run the executable
CMD ["./main"] 
