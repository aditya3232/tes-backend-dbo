FROM golang:1.21

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN go get
RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]

# docker build . -t go-containerized:latest -> run this command to build the image
# docker image ls | grep go-containerized -> run this command to check the image

# docker run -p 8080:8080 go-containerized:latest -> run this command to run the image
# docker run -e PORT=3636 -p 3636:3636 tes-backend-dbo:latest -> run this command to run the image with environment variable