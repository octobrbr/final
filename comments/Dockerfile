FROM golang:1.19-alpine

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o comments ./cmd/main.go

CMD ["./comments"]