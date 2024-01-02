FROM golang:latest

WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

WORKDIR /app/cmd
RUN go build -o bill-manager

EXPOSE 8001

CMD ["./bill-manager"]