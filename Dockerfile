FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY ./server/go.mod ./server/go.sum ./

COPY ./server/ ./server/

WORKDIR /app/server
	
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

# Stage 2: Create a minimal image for the final executable
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /server .

# Expose the port your Go server listens on
EXPOSE 8080

# Command to run the application
CMD ["./server"]
