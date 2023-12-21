FROM golang:1.19

ENV GOPATH=/

WORKDIR .
COPY . .

RUN go mod download
RUN go build -o hhhttteel-app main.go

CMD ["./hhhttteel-app" >> output]