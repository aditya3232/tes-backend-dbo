FROM golang:1.21

WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker caching
COPY go.mod .
COPY go.sum .

# # Copy .env file specifically
COPY tesbackenddbo.env .

# Copy the rest of the application code, including .env file
COPY . .

# Install dependencies
RUN go get

# Build the application
RUN go build -o bin .

# Specify the entry point
ENTRYPOINT ["/app/bin"]

# docker build . -t tes-backend-dbo:latest -> run this command to build the image
# docker image ls | grep tes-backend-dbo -> run this command to check the image

# docker run -e PORT=3636 -p 3636:3636 tes-backend-dbo:latest -> run this command to run the image with environment variable