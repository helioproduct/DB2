FROM golang:1.19

ENV GOPATH=/

WORKDIR .
COPY . .

RUN go mod download
RUN go build -o httteell-app main.go

CMD ["./httteell-app"]