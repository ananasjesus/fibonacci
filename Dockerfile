FROM golang:1.16-buster

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o fibonacci ./cmd/main.go

CMD ["./fibonacci"]