FROM golang:1.16-alpine
WORKDIR /app

# prevent the re-installation of vendors at every change in the source code
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /server

EXPOSE 8080

CMD ["/server"]