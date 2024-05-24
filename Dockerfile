# Use the official Go image with version 1.22
FROM golang:1.22-alpine

# Set the working directory
WORKDIR /app

# Copy the application source code
COPY . .

# Install dependencies
RUN go mod download

# Build the application
RUN go build -o main .

# Expose the port 8080
EXPOSE 8080

# Define the default command
CMD ["/app/main"]
