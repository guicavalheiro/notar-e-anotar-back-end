FROM golang:1.16

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
# COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o ./bin/main ./cmd/server

# Move to /build/bin and set permissions to the executable
WORKDIR /build/bin
RUN chmod +x main

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["/build/bin/main"]