FROM golang:latest

WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o api-backend-test

EXPOSE 8000

CMD ["./api-backend-test"]