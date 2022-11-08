FROM golang:1.17-buster

WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN go build -v ./...

RUN go build -o /creditAssignment

EXPOSE 8089

CMD ["/creditAssignment"]