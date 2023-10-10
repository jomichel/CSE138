FROM golang:1.21
LABEL authors="jonny"
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /server
CMD ["/server"]
