# Build Image
FROM golang:1.16-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY *.go ./
RUN go build -o /imagehost

VOLUME ["/images"]

EXPOSE 10110

CMD ["/imagehost"]
