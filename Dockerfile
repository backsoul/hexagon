FROM golang:1.21-alpine as hexagon
WORKDIR /hexagon

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o hexagon ./cmd/main.go

EXPOSE 8000
CMD ["./hexagon"]